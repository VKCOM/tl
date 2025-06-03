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

const randomSeed = 123432
const PathToBytesData = "../../data/test-stress-data-goldmaster.json"
const NumberOfSamples = 10

func TestGoldmasterStressTest(t *testing.T) {
	tests, err := readTestData()
	if err != nil {
		t.Fatal(err.Error())
	}
	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(randomSeed)))

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

func TestGoldmasterStressTestTL2(t *testing.T) {
	tests, err := readTestData()
	if err != nil {
		t.Fatal(err.Error())
	}
	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(randomSeed)))

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
						readData := utils.ParseHexToBytesTL2(success.BytesTL2)
						readOffset, err := dst.ReadTL2(readData, nil)
						if err != nil {
							t.Fatalf("read error: %s", err.Error())
						}
						if len(readOffset) > 0 {
							t.Fatalf("read tl2 offset not zero")
						}

						newDst := factory.CreateObjectFromName(testingInfo.TestingType)
						if success.IsTLBytesBoxed {
							_, err := newDst.ReadBoxed(utils.ParseHexToBytes(success.Bytes))
							if err != nil {
								t.Fatalf("read tl1 error: %s", err.Error())
							}
						} else {
							_, err := newDst.Read(utils.ParseHexToBytes(success.Bytes))
							if err != nil {
								t.Fatalf("read tl1 error: %s", err.Error())
							}
						}

						writeReturn := dst.WriteTL2(nil, &basictl.TL2WriteContext{})
						if !assert.Equal(t, success.BytesTL2, utils.SprintHexDumpTL2(writeReturn)) {
							t.Fatalf("write tl2 unexpected result")
						}
					})
				}
			})
		})
	}
}

func TestGoldmasterUpdateTL2StressTestData(t *testing.T) {
	restoredValues, restoreErr := readTestData()
	if restoreErr != nil {
		createTestSamples(t)
	} else {
		changedSomething := false
		context := basictl.TL2WriteContext{SizeBuffer: make([]int, 100)}
		writeBuffer := make([]byte, 100)

		for testName, _ := range restoredValues.Tests {
			testingType := restoredValues.Tests[testName].TestingType
			for i, testCase := range restoredValues.Tests[testName].Successes {
				obj := factory.CreateObjectFromName(testingType)
				if obj == nil {
					t.Fatalf("can't create \"%s\"", testingType)
				}
				tl1Data := utils.ParseHexToBytes(testCase.Bytes)
				if testCase.IsTLBytesBoxed {
					_, _ = obj.ReadBoxed(tl1Data)
				} else {
					_, _ = obj.Read(tl1Data)
				}
				writeBuffer = obj.WriteTL2(writeBuffer[0:0], &context)
				if testCase.BytesTL2 != utils.SprintHexDumpTL2(writeBuffer) {
					changedSomething = true
					restoredValues.Tests[testName].Successes[i].BytesTL2 = utils.SprintHexDumpTL2(writeBuffer)
				}
			}
		}

		if changedSomething {
			_ = writeTestData(restoredValues)
		}
	}
}

func createTestSamples(t *testing.T) {
	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(randomSeed)))
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
				t.Fatal(err.Error())
			}
			tl2write := dst.WriteTL2(nil, &basictl.TL2WriteContext{})

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
		t.Fatal(err.Error())
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
	jsonBytes, _ := json.MarshalIndent(tests, "", "\t")
	file, err := os.Create(PathToBytesData)

	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(jsonBytes)
	if err != nil {
		return err
	}
	return nil
}
