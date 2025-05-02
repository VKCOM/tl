package casetests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/factory"
	"github.com/vkcom/tl/internal/utils"
	"github.com/vkcom/tl/pkg/basictl"
	"math/rand"
	"testing"
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
			resultTL2, sizes := mt.object.WriteTL2(nil, nil)

			assert.Empty(t, sizes)
			assert.Equal(t, success.BytesTL2, utils.SprintHexDumpTL2(resultTL2))

			_, readTL2Err := mt.object.ReadTL2(resultTL2)
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
