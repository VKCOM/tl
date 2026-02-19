package casetests

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/factory"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/meta"
	"github.com/vkcom/tl/internal/utils"
	"github.com/vkcom/tl/pkg/basictl"
)

func TestGeneralCasesTL2(t *testing.T) {
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
			runMappingTestBytesTL2(t, mappingTestBytes{
				object:  testObject,
				samples: testValues,
			})
		})
	}
}

func runMappingTestBytesTL2(t *testing.T, mt mappingTestBytes) {
	seed := rand.Uint64()
	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(int64(seed))))

	fmt.Println("Seed: ", seed)

	for sId, success := range mt.samples.Successes {
		t.Run(fmt.Sprintf("Object %d - bytes [%s]", sId, success.Bytes), func(t *testing.T) {
			mt.object.FillRandom(rg)

			trueBytes := utils.ParseHexToBytes(success.Bytes)
			_, readErr := mt.object.Read(trueBytes)

			assert.Nil(t, readErr)
			resultTL2 := mt.object.WriteTL2(nil, nil)

			assert.Equal(t, success.BytesTL2, utils.SprintHexDumpTL2(resultTL2))

			_, readTL2Err := mt.object.ReadTL2(resultTL2, nil)
			assert.Nil(t, readTL2Err)

			newBytes, writeErr := mt.object.WriteGeneral(nil)
			assert.Nil(t, writeErr)
			assert.Equal(t, success.Bytes, utils.SprintHexDump(newBytes))
		})

		if t.Failed() {
			return
		}
	}
}

var bannedTypes []string

func TestGeneralCasesTL2Random(t *testing.T) {
	const NumberOfSamples = 20
	seed := rand.Int63()
	fmt.Printf("Seed: %d\n", seed)

	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(seed)))
	allItems := meta.GetAllTLItems()
	sort.Slice(allItems, func(i, j int) bool {
		return strings.Compare(allItems[i].TLName(), allItems[j].TLName()) > 0
	})

	context := basictl.TL2WriteContext{SizeBuffer: make([]int, 100)}
	writeBuffer := make([]byte, 100)

	bannedSet := utils.SliceToSet(bannedTypes)

	for _, item := range allItems {
		if _, ok := bannedSet[item.TLName()]; ok {
			continue
		}
		t.Run(item.TLName(), func(t *testing.T) {
			dstFun := factory.CreateFunctionFromName(item.TLName())
			dst := factory.CreateObjectFromName(item.TLName())
			if dst == nil {
				t.Fatalf("can't init %s", item.TLName())
			}

			for i := 0; i < NumberOfSamples; i++ {
				dst.FillRandom(rg)
				data, err := dst.WriteGeneral(nil)
				if err != nil {
					t.Fatalf("can't seriliaze %d-th object", i)
				}
				var resultData []byte
				if dstFun != nil {
					resultData, err = dstFun.FillRandomResult(rg, resultData[:0])
					if err != nil {
						t.Fatalf("can't serialize %d-th function result", i)
					}
				}

				t.Run(fmt.Sprintf("TL[%s]", utils.SprintHexDump(data)), func(t *testing.T) {
					writeBuffer = dst.WriteTL2(writeBuffer[:0], &context)
					newDst := factory.CreateObjectFromName(item.TLName())
					_, err = newDst.ReadTL2(writeBuffer, nil)
					if err != nil {
						writeBuffer = dst.WriteTL2(writeBuffer[:0], &context)
						_, err = newDst.ReadTL2(writeBuffer, nil)
						t.Fatalf("can't readTL2 %d-th object", i)
					}
					newData, err := newDst.WriteGeneral(nil)
					if err != nil {
						t.Fatalf("can't write %d-th object", i)
					}
					if !assert.Equal(t, utils.SprintHexDump(data), utils.SprintHexDump(newData), fmt.Sprintf("Seed %d", seed)) {
						fmt.Printf("place for a breakpoint\n")
					}

					if dstFun == nil {
						return
					}
					_, writeBuffer, err = dstFun.ReadResultWriteResultTL2(nil, resultData, writeBuffer[:0])
					if err != nil {
						t.Fatalf("can't readTL2 %d-th result", i)
					}
					_, newData, err = dstFun.ReadResultTL2WriteResult(nil, writeBuffer, newData[:0])
					if err != nil {
						t.Fatalf("can't write %d-th result", i)
					}
					if !assert.Equal(t, utils.SprintHexDump(resultData), utils.SprintHexDump(newData), fmt.Sprintf("Seed %d", seed)) {
						fmt.Printf("place for a breakpoint\n")
					}
				})
			}
		})
	}
}

func TestExactTestCasesTL2Random(t *testing.T) {
	objectName := "casesTL2.testObject"                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           // place here
	objectData := "00000000_bc799737_0000000b_bc799737_997275b5_997275b5_997275b5_bc799737_997275b5_bc799737_bc799737_bc799737_bc799737_bc799737_00000007_ce27c770_ef556bee_00000000_ef556bee_00000000_ef556bee_00000000_ef556bee_0000000b_00000000_1da8fbff_1e22d321_068c7baf_31d1eb5f_386f845d_5fc6eb6d_6ac9e76e_3eadd2e3_bd170f35_401dde0b_00000000_916ef505_46f55329_dcbc7a01_06a89d32_d17b0afc_1eec065c_7a09df46_45dbe0bb_6f345428_605632ae_00000000_0329de4e_1deb80e9_bf518153_281ce2a6_fc15794c_0ec42392_4e58352b_6796f592_7c1e00f6_2fb066ab_00000000_317fddee_2451a08a_754deccc_03ed165a_516d2ef8_625a2175_223f9fcf_58200c44_c8c49a95_3c35a7bc_00000000_d5e11a0a_7967305d_8505a8c3_765db5c8_f2878bff_616c2105_fc7e2b35_1ff3f64f_657475b8_6b69d865_00000003_18e804c0_6509a6d3_f6a76b21_333dda77_c85cd794_76ed7c5e_1473cbc4_1a329ccd_fd003b71_2fd9f34b_00000008_b691fa09_05618924_d7e4e585_59e124e6_159d183b_72e33920_b4d25a3c_3b3f001a_1a403604_093ced22_0000000b_4d058513_6289a89d_bcc6b3e1_2a36e635_a7abe3e1_3936e81f_9b664a14_2381c407_0e7cea6d_252e2f3c_00000003_64a53cb0_07489f5c_b4a80541_5c42bfa0_d33b696c_756f3735_2ed6cda6_5e232edd_2941cf33_56944bbb_00000000_79bbbe85_338fc962_62aa4b8a_5c020889_38473fdb_191ab414_7d6757fa_4b42e1e4_537fd41f_418682aa_00000000_007c3da7_5e360822_5c2c74ca_74f01136_cea7ef4c_1ccd6a39_8b4d5d8b_58249cab_ab786157_39433a09_ef556bee_0000000b_00000003_434d0840_155e14d9_4a3df50d_205936d6_e4023b1f_385714f1_f461b64f_491dc74b_c4acfeec_45480105_00000000_5d3c9d12_5e5597c5_6b9938bd_71e17244_e20bfb7d_6da67f8f_5f0b01e9_727c1bca_e44b1060_37bf84d1_00000001_60cb6258_2387ca41_fef3bded_380b39af_e1378780_20577d86_19736d39_7e78cb6d_f83fe615_64a76b6c_00000000_4b35eef2_4e79f291_5b5205b3_3f3e68b5_1167f2a8_013ea4d3_28e0c698_43e4bf74_d9dd3aba_65762e70_00000000_2f3533c8_2fb08027_1dded17c_384d78f1_3bfa21fe_52f79a63_2bd7a7bd_281a6c61_047748de_2216096f_00000009_3e8d93a8_469f601f_e6107171_15c5a37e_bb9651da_43e32784_fc4e5d64_4597f4b4_c668166a_25ed960f_00000001_bc12e2c7_21f5e45c_ac8acdad_584289ca_c921ea2d_51dee886_37d5b227_0d248cde_23759319_0cae3347_00000000_e852c2a0_0d5f593d_c03ec643_3907937d_98330c32_74e383e8_409a5f48_70f0a017_6b6e522c_3af61423_00000002_5f9066f7_5d51fcd9_47147aa1_070fcb45_ab70f1db_67a2a822_039ebded_0a667461_0a8ad024_5a6e8d49_00000000_20f74ba0_4878f422_b015d18a_5415073e_9b88611f_1658bb77_663c676f_570b4e48_901fa35c_5f1efdab_00008009_f49ac4c8_1877ef49_0659ce48_73037c0e_cd4f232c_343f2119_6071a925_36dc9e79_c9f77969_385c5ef8_ef556bee_00000000" // place here

	context := basictl.TL2WriteContext{SizeBuffer: make([]int, 100)}
	writeBuffer := make([]byte, 100)

	dst := factory.CreateObjectFromName(objectName)
	if dst == nil {
		t.Fatalf("can't init %s", objectName)
	}

	cleanBytes := utils.ParseHexToBytes(strings.ReplaceAll(objectData, "_", " "))
	_, err := dst.Read(cleanBytes)
	if err != nil {
		t.Fatalf("can't read %s", objectData)
	}

	writeBuffer = dst.WriteTL2(writeBuffer[:0], &context)
	newDst := factory.CreateObjectFromName(objectName)
	_, err = newDst.ReadTL2(writeBuffer, nil)
	if err != nil {
		t.Fatalf("can't readTL2 object")
	}
	newData, err := newDst.WriteGeneral(nil)
	if err != nil {
		t.Fatalf("can't write object")
	}
	assert.Equal(t, utils.SprintHexDump(cleanBytes), utils.SprintHexDump(newData))
}
