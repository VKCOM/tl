package goldmastertests

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"testing"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/pure/vkext"
	testformat "github.com/VKCOM/tl/internal/tlcodegen/test/codegen_test/golang/common"
	"github.com/VKCOM/tl/internal/tlcodegen/test/gen/goldmaster/factory"
	"github.com/VKCOM/tl/internal/tlcodegen/test/gen/goldmaster/meta"
	"github.com/VKCOM/tl/internal/utils"
	"github.com/stretchr/testify/assert"

	"github.com/VKCOM/tl/pkg/basictl"
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

	kernel := pure.NewKernel(&pure.OptionsKernel{})
	if err := kernel.AddFileTL1("../../../tls/goldmaster.tl"); err != nil {
		t.Fatal(err.Error())
	}
	if err := kernel.AddFileTL1("../../../tls/goldmaster2.tl"); err != nil {
		t.Fatal(err.Error())
	}
	if err := kernel.AddFileTL1("../../../tls/goldmaster3.tl"); err != nil {
		t.Fatal(err.Error())
	}
	if err := kernel.Compile(); err != nil {
		t.Fatal(err.Error())
	}

	for _, testName := range testNames {
		t.Run(testName, func(t *testing.T) {
			testingInfo := tests.Tests[testName]
			t.Run(testingInfo.TestingType, func(t *testing.T) {
				dst := factory.CreateObjectFromName(testingInfo.TestingType)
				if dst == nil {
					t.Fatalf("can't create object of type %q", testingInfo.TestingType)
				}
				dst.FillRandom(rg)
				ins := kernel.GetObjectInstance(testingInfo.TestingType)
				var dst2 vkext.KernelValue
				if ins == nil {
					fmt.Printf("probably union variant %q, skipping\n", testingInfo.TestingType)
					//t.Fatalf("can't create vkext object of type %q", testingInfo.TestingType)
				} else {
					if _, ok := ins.(*pure.TypeInstancePrimitive); !ok {
						// factory and canonical types are different for string, canonical is String for wrapper
						dst2 = vkext.CreateValue(ins)
					}
				}
				for _, success := range testingInfo.Successes {
					t.Run(fmt.Sprintf("TL[%s]", success.Bytes), func(t *testing.T) {
						writeFunc := dst.WriteTL1General
						readFunc := dst.ReadTL1

						if success.IsTLBytesBoxed {
							writeFunc = dst.WriteTL1BoxedGeneral
							readFunc = dst.ReadTL1Boxed
						}

						origBytes := utils.ParseHexToBytes(success.Bytes)
						_, err := readFunc(origBytes)
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
						if dst2 == nil {
							return // skipping
						}
						_, _, err = dst2.ReadTL1(origBytes, nil, !success.IsTLBytesBoxed, nil)
						if err != nil {
							t.Fatalf("read vkext error: %s", err.Error())
						}

						var writeReturn2 vkext.ByteBuilder
						_ = dst2.WriteTL1(&writeReturn2, !success.IsTLBytesBoxed, nil, false, 0, nil)
						if !assert.Equal(t, success.Bytes, utils.SprintHexDump(writeReturn2.Buf())) {
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
				switch testingInfo.TestingType {
				case "useDictUgly":
					// StrangeDictElem has tl2mask0, which is never set, because it is stored as a map,
					// not a struct
					return
				}
				dst := factory.CreateObjectFromName(testingInfo.TestingType)
				if dst == nil {
					t.Fatalf("can't create object of type %q", testingInfo.TestingType)
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
							_, err := newDst.ReadTL1Boxed(utils.ParseHexToBytes(success.Bytes))
							if err != nil {
								t.Fatalf("read tl1 error: %s", err.Error())
							}
						} else {
							_, err := newDst.ReadTL1(utils.ParseHexToBytes(success.Bytes))
							if err != nil {
								t.Fatalf("read tl1 error: %s", err.Error())
							}
						}

						writeReturn := dst.WriteTL2(nil, &basictl.TL2WriteContext{})
						if !assert.Equal(t, success.BytesTL2, utils.SprintHexDumpTL2(writeReturn)) {
							writeReturn2 := dst.WriteTL2(nil, &basictl.TL2WriteContext{})
							//writeReturn3 := newDst.WriteTL2(nil, &basictl.TL2WriteContext{})
							fmt.Printf("%s %x\n", testingInfo.TestingType, writeReturn2)
							//fmt.Printf("%s %x\n", testingInfo.TestingType, writeReturn3)
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

		for testName := range restoredValues.Tests {
			testingType := restoredValues.Tests[testName].TestingType
			for i, testCase := range restoredValues.Tests[testName].Successes {
				obj := factory.CreateObjectFromName(testingType)
				if obj == nil {
					t.Fatalf("can't create %q", testingType)
				}
				tl1Data := utils.ParseHexToBytes(testCase.Bytes)
				if testCase.IsTLBytesBoxed {
					_, _ = obj.ReadTL1Boxed(tl1Data)
				} else {
					_, _ = obj.ReadTL1(tl1Data)
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

		dst := factory.CreateObjectFromName(item.TLName())
		if dst == nil {
			t.Fatalf("can't init object")
		}
		if _, ok := dst.(meta.TLItem); ok {
			continue
		}

		for i := 0; i < NumberOfSamples; i++ {
			dst.FillRandom(rg)
			tl1write, err := dst.WriteTL1BoxedGeneral(nil)
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
	//jsonBytes, _ := json.MarshalIndent(tests, "", "\t")
	//return os.WriteFile(PathToBytesData, jsonBytes, 0666)
	return nil
}
