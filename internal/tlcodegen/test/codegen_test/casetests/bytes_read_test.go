// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package casetests

import (
	"encoding/json"
	"fmt"
	"github.com/vkcom/tl/internal/utils"
	"math/rand"
	"os"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/factory"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/meta"
	"github.com/vkcom/tl/pkg/basictl"

	"github.com/stretchr/testify/assert"
	"testing"
)

type MappingSuccessBytes struct {
	// expected bytes input and output
	Bytes string
}

type mappingTestSamplesBytes struct {
	// TL name for type to test
	TestingType string
	// json values which must success
	Successes []MappingSuccessBytes
}

type mappingTestBytes struct {
	// testing type object
	object meta.Object
	// testing samples with results
	samples mappingTestSamplesBytes
}

type allTestsBytes struct {
	Tests map[string]mappingTestSamplesBytes
}

func runMappingTestBytes(t *testing.T, mt mappingTestBytes) {
	seed := rand.Uint64()
	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(int64(seed))))

	fmt.Println("Seed: ", seed)

	for sId, success := range mt.samples.Successes {
		t.Run(fmt.Sprintf("Object %d", sId), func(t *testing.T) {
			mt.object.FillRandom(rg)

			trueBytes := utils.ParseHexToBytes(success.Bytes)
			_, readErr := mt.object.Read(trueBytes)

			assert.Nil(t, readErr)
			writeData, writeErr := mt.object.WriteGeneral(nil)

			assert.Nil(t, writeErr)
			assert.Equal(t, trueBytes, writeData)

			_, readAgainErr := mt.object.Read(trueBytes)
			assert.Nil(t, readAgainErr)

			writeAgainData, writeAgainErr := mt.object.WriteGeneral(nil)

			assert.Nil(t, writeAgainErr)
			assert.Equal(t, trueBytes, writeAgainData)
		})

		if t.Failed() {
			return
		}
	}
}

func TestAllTLObjectsReadJsonByRandomBytes(t *testing.T) {
	const RepeatNumber = 100

	seed := uint64(12501105899753422230) // rand.Uint64()

	t.Logf("Seed: %d\n", seed)

	buf1 := make([]byte, 0)
	buf2 := make([]byte, 0)

	var err error

	for _, tlItem := range meta.GetAllTLItems() {
		t.Run(tlItem.TLName(), func(t *testing.T) {
			// TODO need to check with Grisha
			if tlItem.TLName() == "cases.replace7" || tlItem.TLName() == "cases.replace7plus" || tlItem.TLName() == "cases.replace7plusplus" {
				t.Skip("Skip until checkSanity will be fixed")
				return
			}

			rnd := rand.New(rand.NewSource(int64(seed)))

			var objects []meta.Object
			for i := 0; i < RepeatNumber; i++ {
				obj := factory.CreateObject(tlItem.TLTag())
				obj.FillRandom(basictl.NewRandGenerator(rnd))
				objects = append(objects, obj)
			}

			for i := 0; i < RepeatNumber; i++ {
				buf1 = buf1[:0]
				buf2 = buf2[:0]
				t.Run(fmt.Sprintf("Object %d", i), func(t *testing.T) {
					obj := objects[i]
					buf1, err = obj.WriteGeneral(buf1)
					if err != nil {
						t.Logf("Seed: %d\n", seed)
						t.Fatal("first serialization wasn't succeeded", err.Error())
						return
					}
					_, err = obj.Read(buf1)
					if err != nil {
						t.Logf("Seed: %d\n", seed)
						t.Fatal("first deserialization wasn't succeeded", err.Error())
						return
					}
					obj1 := obj
					buf2, err = obj.WriteGeneral(buf2)
					if err != nil {
						t.Logf("Seed: %d\n", seed)
						t.Fatal("second serialization wasn't succeeded", err.Error())
						return
					}
					_, err = obj.Read(buf2)
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
		})
	}
}

func TestGeneralCasesBytes(t *testing.T) {
	const PathToJsonData = "../data/test-objects-bytes.json"
	data, readErr := os.ReadFile(PathToJsonData)

	if readErr != nil {
		t.Fatalf("testing data is not provided")
		return
	}

	tests := allTestsBytes{map[string]mappingTestSamplesBytes{}}
	err := json.Unmarshal(data, &tests)

	if err != nil {
		t.Fatalf("can't unmarshall test data")
		return
	}

	for testName, testValues := range tests.Tests {
		t.Run(testName, func(t *testing.T) {
			testObject := factory.CreateObjectFromName(testValues.TestingType)
			if testObject == nil {
				t.Fatalf("No testing object for test \"%s\"", testName)
				return
			}
			runMappingTestBytes(t, mappingTestBytes{
				object:  testObject,
				samples: testValues,
			})
		})
	}
}
