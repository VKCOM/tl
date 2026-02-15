// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package casetests

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/factory"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/factory_bytes"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/meta"
	"github.com/vkcom/tl/pkg/basictl"

	"github.com/stretchr/testify/assert"
)

type MappingSuccess struct {
	// absolute correct json form of same object
	GoldenInput string
	// json Alternatives which map to GoldenInput
	Alternatives []string
	// wrong json Alternatives
	IncorrectAlternatives []string
}

type mappingTestSamples struct {
	// TL name for type to test
	TestingType string
	// use bytes version of code
	UseBytes bool
	// json values which must success
	Successes []MappingSuccess
	// json values which must fail on read
	Failures []string
}

type mappingTest struct {
	// testing type object
	object meta.Object
	// testing samples with results
	samples mappingTestSamples
}

type allTests struct {
	Tests map[string]mappingTestSamples
}

func runMappingTest(t *testing.T, mt mappingTest) {
	seed := rand.Uint64()
	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(int64(seed))))

	fmt.Println("Seed: ", seed)

	for sId, success := range mt.samples.Successes {
		alternatives := success.Alternatives
		if len(alternatives) == 0 {
			alternatives = append(alternatives, success.GoldenInput)
		}
		for aId, alternative := range alternatives {
			t.Run(fmt.Sprintf("Object %d - Alternative %d", sId, aId), func(t *testing.T) {
				mt.object.FillRandom(rg)
				readErr := mt.object.ReadJSON(false, &basictl.JsonLexer{Data: []byte(alternative)})

				assert.Nil(t, readErr)
				writeData, writeErr := mt.object.MarshalJSON()

				assert.Nil(t, writeErr)
				assert.Equal(t, success.GoldenInput, string(writeData))

				readAgainErr := mt.object.ReadJSON(false, &basictl.JsonLexer{Data: []byte(success.GoldenInput)})
				assert.Nil(t, readAgainErr)

				writeAgainData, writeAgainErr := mt.object.MarshalJSON()

				assert.Nil(t, writeAgainErr)
				assert.Equal(t, success.GoldenInput, string(writeAgainData))
			})

			if t.Failed() {
				return
			}
		}

		for aId, alternative := range success.IncorrectAlternatives {
			t.Run(fmt.Sprintf("Object %d - Wrong alternative %d", sId, aId), func(t *testing.T) {
				mt.object.FillRandom(rg)
				readErr := mt.object.ReadJSON(false, &basictl.JsonLexer{Data: []byte(alternative)})

				assert.Nil(t, readErr)
				writeData, writeErr := mt.object.MarshalJSON()

				assert.Nil(t, writeErr)
				assert.NotEqual(t, success.GoldenInput, string(writeData))
			})

			if t.Failed() {
				return
			}
		}
	}

	for fId, failure := range mt.samples.Failures {
		t.Run(fmt.Sprintf("Failure %d", fId), func(t *testing.T) {
			mt.object.FillRandom(rg)
			readErr := mt.object.ReadJSON(false, &basictl.JsonLexer{Data: []byte(failure)})

			assert.NotNil(t, readErr)
		})

		if t.Failed() {
			return
		}
	}
}

func TestAllTLObjectsReadJsonByRandom(t *testing.T) {
	const RepeatNumber = 100

	seed := rand.Uint64()
	rnd := rand.New(rand.NewSource(int64(seed)))

	t.Logf("Seed: %d\n", seed)

	buf1 := make([]byte, 0)
	buf2 := make([]byte, 0)

	var err error

	for _, tlItem := range meta.GetAllTLItems() {
		obj := factory.CreateObject(tlItem.TLTag())
		for i := 0; i < RepeatNumber; i++ {
			buf1 = buf1[:0]
			buf2 = buf2[:0]
			t.Run(tlItem.TLName(), func(t *testing.T) {
				obj.FillRandom(basictl.NewRandGenerator(rnd))
				buf1, err = obj.WriteJSONGeneral(&basictl.JSONWriteContext{}, buf1)
				if err != nil {
					t.Logf("Seed: %d\n", seed)
					t.Fatal("first serialization wasn't succeeded", err.Error())
					return
				}
				err = obj.ReadJSON(false, &basictl.JsonLexer{Data: buf1})
				if err != nil {
					t.Logf("Seed: %d\n", seed)
					t.Fatal("first deserialization wasn't succeeded", err.Error())
					return
				}
				obj1 := obj
				buf2, err = obj.WriteJSONGeneral(&basictl.JSONWriteContext{}, buf2)
				if err != nil {
					t.Logf("Seed: %d\n", seed)
					t.Fatal("second serialization wasn't succeeded", err.Error())
					return
				}
				err = obj.ReadJSON(false, &basictl.JsonLexer{Data: buf2})
				if err != nil {
					t.Logf("Seed: %d\n", seed)
					t.Fatal("second deserialization wasn't succeeded", err.Error())
					return
				}
				obj2 := obj
				assert.Equal(t, buf1, buf2, "serializations must be same")
				assert.Equal(t, obj1, obj2, "translations to object must be equal")
				if t.Failed() {
					t.Logf("Seed: %d\n", seed)
					t.Logf("Test failed on %d iteration to %s", i, tlItem.TLName())
					return
				}
			})
		}
	}
}

func TestGeneralCases(t *testing.T) {
	const PathToJsonData = "../../data/test-objects-json.json"
	data, readErr := os.ReadFile(PathToJsonData)

	if readErr != nil {
		t.Fatalf("testing data is not provided")
		return
	}

	tests := allTests{map[string]mappingTestSamples{}}
	err := json.Unmarshal(data, &tests)

	if err != nil {
		t.Fatalf("can't unmarshall test data")
		return
	}

	for testName, testValues := range tests.Tests {
		t.Run(testName, func(t *testing.T) {
			testObject := factory.CreateObjectFromName(testValues.TestingType)
			if testValues.UseBytes {
				testObject = factory_bytes.CreateObjectFromName(testValues.TestingType)
			}
			if testObject == nil {
				t.Fatalf("No testing object for test \"%s\"", testName)
				return
			}
			runMappingTest(t, mappingTest{
				object:  testObject,
				samples: testValues,
			})
		})
	}
}
