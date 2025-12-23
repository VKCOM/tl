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

	"github.com/vkcom/tl/internal/utils"

	common "github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/tl"
	cases "github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/tlcases"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/factory"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/meta"
	"github.com/vkcom/tl/pkg/basictl"

	"github.com/stretchr/testify/assert"

	testformat "github.com/vkcom/tl/internal/tlcodegen/test/codegen_test/golang/common"
)

type mappingTestBytes struct {
	// testing type object
	object meta.Object
	// testing samples with results
	samples testformat.MappingTestSamplesBytes
}

func runMappingTestBytes(t *testing.T, mt mappingTestBytes) {
	seed := rand.Uint64()
	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(int64(seed))))

	fmt.Println("Seed: ", seed)

	for sId, success := range mt.samples.Successes {
		t.Run(fmt.Sprintf("Object %d - bytes [%s]", sId, success.Bytes), func(t *testing.T) {
			mt.object.FillRandom(rg)

			trueBytes := utils.ParseHexToBytes(success.Bytes)
			_, readErr := mt.object.Read(trueBytes)

			assert.Nil(t, readErr)
			writeData, writeErr := mt.object.WriteGeneral(nil)

			assert.Nil(t, writeErr)
			assert.Equal(t, utils.SprintHexDump(trueBytes), utils.SprintHexDump(writeData))

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

	seed := rand.Uint64()

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

const PathToBytesData = "../../data/test-objects-bytes.json"

func TestAppendNewCasesForTesting(t *testing.T) {
	tests, success := initTestData(t)
	if !success {
		return
	}

	type Sample struct {
		sample   meta.Object
		testName string
		// if "" then check is skipped
		expectingHexValue string
	}

	// write your samples here
	newSamples := []Sample{
		{
			testName: "TestInplace",
			sample: &cases.TestInplaceStructArgs{
				A1: 1,
				A2: 2,
				A3: 3,
				Arg: cases.Inplace1Int{Value: cases.Inplace2Int{Value: cases.Inplace3TupleInt2{
					Value: common.PairTupleTupleInt2TupleTupleInt2{
						X: common.TupleTupleInt2{[2]int32{3, 4}},
						Y: common.TupleTupleInt2{[2]int32{5, 6}, [2]int32{7, 8}},
					},
				}}},
			},
			expectingHexValue: "00000001 00000002 00000003 00000003 00000004 00000005 00000006 00000007 00000008",
		},
		{
			testName: "TestInplace2",
			sample: &cases.TestInplaceStructArgs2{
				A1: 1,
				A2: 2,
				A3: 3,
				Arg: cases.Inplace1PairTupleIntTupleInt{Value: cases.Inplace2PairTupleIntTupleInt{Value: cases.Inplace3TuplePairTupleIntTupleInt2{
					Value: common.PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2{
						X: common.TupleTuplePairTupleIntTupleInt2{
							common.TuplePairTupleIntTupleInt2{
								common.PairTupleIntTupleInt{
									X: common.TupleInt{
										4,
										5,
										6,
									},
									Y: common.TupleInt{
										7,
										8,
									},
								},
								common.PairTupleIntTupleInt{
									X: common.TupleInt{
										9,
										10,
										11,
									},
									Y: common.TupleInt{
										12,
										13,
									},
								},
							},
						},
						Y: common.TupleTuplePairTupleIntTupleInt2{
							common.TuplePairTupleIntTupleInt2{
								common.PairTupleIntTupleInt{
									X: common.TupleInt{
										14,
										15,
										16,
									},
									Y: common.TupleInt{
										17,
										18,
									},
								},
								common.PairTupleIntTupleInt{
									X: common.TupleInt{
										19,
										20,
										21,
									},
									Y: common.TupleInt{
										22,
										23,
									},
								},
							},
							common.TuplePairTupleIntTupleInt2{
								common.PairTupleIntTupleInt{
									X: common.TupleInt{
										24,
										25,
										26,
									},
									Y: common.TupleInt{
										27,
										28,
									},
								},
								common.PairTupleIntTupleInt{
									X: common.TupleInt{
										29,
										30,
										31,
									},
									Y: common.TupleInt{
										32,
										33,
									},
								},
							},
						},
					},
				}}},
			},
			expectingHexValue: "00000001 00000002 00000003 00000004 00000005 00000006 00000007 00000008 00000009 0000000a 0000000b 0000000c 0000000d 0000000e 0000000f 00000010 00000011 00000012 00000013 00000014 00000015 00000016 00000017 00000018 00000019 0000001a 0000001b 0000001c 0000001d 0000001e 0000001f 00000020 00000021",
		},
	}

	newSamplesCount := 0

	for i, sample := range newSamples {
		success, err := addSample(&tests, sample.testName, sample.sample, sample.expectingHexValue)
		if success {
			newSamplesCount += 1
		} else if err != nil {
			t.Fatalf("Incorrent test sample #%[1]d for \"%[2]s\" was provided, error message: %s", i, sample.sample.TLName(), err.Error())
			return
		}
	}

	fmt.Println("")

	if newSamplesCount != 0 {
		bytes, err := json.MarshalIndent(&tests, "", "\t")

		if err != nil {
			t.Fatalf("can't marshal new values")
			return
		}

		file, err := os.Create(PathToBytesData)

		if err != nil {
			t.Fatalf("can't open file with test cases")
			return
		}
		defer func() { _ = file.Close() }()

		_, err = file.Write(bytes)
		if err != nil {
			t.Fatalf("can't overwrite file with test cases")
			return
		}

		fmt.Printf("To \"%s\" were added %d test cases\n", PathToBytesData, newSamplesCount)
	} else {
		fmt.Printf("Nothing to add to \"%s\"\n", PathToBytesData)
	}

	fmt.Println("")
}

func TestGeneralCasesBytes(t *testing.T) {
	tests, success := initTestData(t)
	if !success {
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

func TestUpdateForTL2(t *testing.T) {
	updateTestData(t)
}

func checkExistenceOfTest(tests *testformat.AllTestsBytes, testName, typeName, bytesHex string) (testCanBeAdded bool, testExists bool, testNameToAdd string) {
	for testNameValue, test := range tests.Tests {
		// testName exists, but it is not for this type then return false
		if testNameValue == testName {
			if test.TestingType != typeName {
				return false, false, ""
			}
		} else {
			// testName different, but maybe it is for same type
			if test.TestingType != typeName {
				continue
			}
		}

		for _, success := range test.Successes {
			if success.Bytes == bytesHex {
				return true, true, testNameValue
			}
		}
	}
	return true, false, testName
}

func addSample(tests *testformat.AllTestsBytes, testName string, sample meta.Object, expectingHex string) (bool, error) {
	bytes, err := sample.WriteGeneral(nil)
	if err != nil {
		return false, err
	}
	hexBytes := utils.SprintHexDump(bytes)
	if expectingHex != "" && hexBytes != expectingHex {
		return false, fmt.Errorf("bytes are different from expecting:\n\tactual:   %s\n\texpected: %s", hexBytes, expectingHex)
	}
	canAdd, exists, placeToAdd := checkExistenceOfTest(tests, testName, sample.TLName(), hexBytes)
	if canAdd && !exists {
		successes := tests.Tests[placeToAdd]
		if successes.Successes == nil {
			successes.TestingType = sample.TLName()
		}
		successes.Successes = append(successes.Successes, testformat.MappingSuccessBytes{Bytes: hexBytes})
		tests.Tests[placeToAdd] = successes
		return true, nil
	}
	return false, nil
}

func initTestData(t *testing.T) (_ testformat.AllTestsBytes, success bool) {
	data, readErr := os.ReadFile(PathToBytesData)

	if readErr != nil {
		t.Fatalf("testing data is not provided")
		return testformat.AllTestsBytes{}, false
	}

	tests := testformat.AllTestsBytes{Tests: map[string]testformat.MappingTestSamplesBytes{}}
	err := json.Unmarshal(data, &tests)

	if err != nil {
		t.Fatalf("can't unmarshall test data")
		return testformat.AllTestsBytes{}, false
	}
	return tests, true
}

func updateTestData(t *testing.T) {
	data, readErr := os.ReadFile(PathToBytesData)
	if readErr != nil {
		t.Fatalf("testing data is not provided")
		return
	}

	tests := testformat.AllTestsBytes{Tests: map[string]testformat.MappingTestSamplesBytes{}}
	_ = json.Unmarshal(data, &tests)

	updated := false

	for testName, testValues := range tests.Tests {
		for i, success := range testValues.Successes {
			if !success.IsTL2DataFixed {
				testObject := factory.CreateObjectFromName(testValues.TestingType)
				if testObject == nil {
					t.Fatalf("No testing object for test \"%s\"", testName)
					return
				}
				bytes := utils.ParseHexToBytes(success.Bytes)
				_, readErr := testObject.Read(bytes)
				if readErr != nil {
					t.Fatalf("can't read %s, reason: %s", testValues.TestingType, readErr)
					return
				}
				data := testObject.WriteTL2(nil, nil)
				tl2StringData := utils.SprintHexDumpTL2(data)

				if success.BytesTL2 != tl2StringData {
					updated = true
					newTestValues := tests.Tests[testName]
					newTestValues.Successes[i].BytesTL2 = tl2StringData
					tests.Tests[testName] = newTestValues
				}
			}
		}
	}

	if updated {
		file, err := os.Create(PathToBytesData)

		if err != nil {
			t.Fatalf("can't open file with test cases")
			return
		}
		defer func() { _ = file.Close() }()

		bytes, _ := json.MarshalIndent(tests, "", "\t")
		_, err = file.Write(bytes)
		if err != nil {
			t.Fatalf("can't overwrite file with test cases")
			return
		}
	}
}
