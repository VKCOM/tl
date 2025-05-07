package goldmastertests

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	testformat "github.com/vkcom/tl/internal/tlcodegen/test/codegen_test/golang/common"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/factory"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/meta"
	"github.com/vkcom/tl/internal/utils"
	"math/rand"
	"os"
	"sort"
	"testing"

	"github.com/vkcom/tl/pkg/basictl"
)

func TestGoldmasterStressTest(t *testing.T) {
	tests, err := readTestData()
	if err != nil {
		t.Fatalf(err.Error())
	}
	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(123432)))

	testNames := utils.Keys(tests.Tests)
	sort.Strings(testNames)

	for _, testName := range testNames {
		t.Run(testName, func(t *testing.T) {
			testingInfo := tests.Tests[testName]
			t.Run(testingInfo.TestingType, func(t *testing.T) {
				dst := factory.CreateObjectFromName(testingInfo.TestingType)
				if dst == nil {
					t.Fatalf("can't create object of type\"%s\"", testingInfo.TestingType)
				}
				dst.FillRandom(rg)
				for _, success := range testingInfo.Successes {
					t.Run(fmt.Sprintf("TL[%s]", success.Bytes), func(t *testing.T) {
						var writeFunc func([]byte) ([]byte, error) = dst.WriteGeneral
						var readFunc func([]byte) ([]byte, error) = dst.Read

						if success.IsTLBytesBoxed {
							writeFunc = dst.WriteBoxedGeneral
							readFunc = dst.ReadBoxed
						}

						_, err := readFunc(utils.ParseHexToBytes(success.Bytes))
						if err != nil {
							t.Fatalf("read error: %s", err.Error())
						}

						writeReturn, err := writeFunc(nil)
						if err != nil {
							t.Fatalf("write error: %s", err.Error())
						}
						if !assert.Equal(t, success.Bytes, utils.SprintHexDump(writeReturn)) {
							return
						}
					})
				}
			})
		})
	}
}

// fix after
//func TestGoldmasterStressTestTL2(t *testing.T) {
//	tests, err := readTestData()
//	if err != nil {
//		t.Fatalf(err.Error())
//	}
//	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(123432)))
//
//	testNames := utils.Keys(tests.Tests)
//	sort.Strings(testNames)
//
//	for _, testName := range testNames {
//		t.Run(testName, func(t *testing.T) {
//			testingInfo := tests.Tests[testName]
//			t.Run(testingInfo.TestingType, func(t *testing.T) {
//				dst := factory.CreateObjectFromName(testingInfo.TestingType)
//				if dst == nil {
//					t.Fatalf("can't create object of type\"%s\"", testingInfo.TestingType)
//				}
//				dst.FillRandom(rg)
//				for _, success := range testingInfo.Successes {
//					t.Run(fmt.Sprintf("TL[%s]", success.Bytes), func(t *testing.T) {
//						_, err := dst.ReadTL2(utils.ParseHexToBytesTL2(success.BytesTL2))
//						if err != nil {
//							t.Fatalf("read error: %s", err.Error())
//						}
//
//						newDst := factory.CreateObjectFromName(testingInfo.TestingType)
//						if success.IsTLBytesBoxed {
//							_, err := newDst.ReadBoxed(utils.ParseHexToBytes(success.Bytes))
//							if err != nil {
//								t.Fatalf("read tl1 error: %s", err.Error())
//							}
//						} else {
//							_, err := newDst.Read(utils.ParseHexToBytes(success.Bytes))
//							if err != nil {
//								t.Fatalf("read tl1 error: %s", err.Error())
//							}
//						}
//
//						if !cmp.Equal(newDst, dst, cmpopts.EquateEmpty()) {
//							t.Fatalf("no same objects: %s", cmp.Diff(newDst, dst, cmpopts.EquateEmpty()))
//						}
//
//						writeReturn, _ := dst.WriteTL2(nil, nil)
//						if !assert.Equal(t, success.BytesTL2, utils.SprintHexDumpTL2(writeReturn)) {
//							return
//						}
//					})
//				}
//			})
//		})
//	}
//}

const PathToBytesData = "../../data/test-stress-data-goldmaster.json"
const UpdateTests = true
const NumberOfSamples = 10

func TestGoldmasterGenerateStressTestData(t *testing.T) {
	if UpdateTests {
		rg := basictl.NewRandGenerator(rand.New(rand.NewSource(123432)))
		tests := testformat.AllTestsBytes{Tests: map[string]testformat.MappingTestSamplesBytes{}}

		bannedSet := utils.SliceToSet(bannedTypes)

		items := meta.GetAllTLItems()
		for _, item := range items {
			if bannedSet[item.TLName()] {
				continue
			}

			testingData := testformat.MappingTestSamplesBytes{}
			testingData.TestingType = item.TLName()

			dst := factory.CreateObject(item.TLTag())
			if dst == nil {
				t.Fatalf("can't init object")
			}
			if _, ok := dst.(*meta.TLItem); ok {
				continue
			}

			for i := 0; i < NumberOfSamples; i++ {
				dst.FillRandom(rg)
				tl1write, err := dst.WriteBoxedGeneral(nil)
				if err != nil {
					t.Fatalf(err.Error())
				}
				tl2write, _ := dst.WriteTL2(nil, nil)

				exactTest := testformat.MappingSuccessBytes{}
				exactTest.Bytes = utils.SprintHexDump(tl1write)
				exactTest.BytesTL2 = utils.SprintHexDumpTL2(tl2write)
				exactTest.IsTLBytesBoxed = true

				testingData.Successes = append(testingData.Successes, exactTest)
			}
			tests.Tests[fmt.Sprintf("Test[%s]", testingData.TestingType)] = testingData
		}

		err := writeTestData(tests)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
}

func readTestData() (testformat.AllTestsBytes, error) {
	data, readErr := os.ReadFile(PathToBytesData)

	if readErr != nil {
		return testformat.AllTestsBytes{}, fmt.Errorf("testing data is not provided")
	}

	tests := testformat.AllTestsBytes{Tests: map[string]testformat.MappingTestSamplesBytes{}}
	err := json.Unmarshal(data, &tests)

	if err != nil {
		return testformat.AllTestsBytes{}, fmt.Errorf("can't unmarshall test data")
	}

	return tests, nil
}

func writeTestData(tests testformat.AllTestsBytes) error {
	bytes, err := json.MarshalIndent(tests, "", "\t")
	file, err := os.Create(PathToBytesData)

	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
