package goldmastertests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/factory"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/meta"
	"github.com/vkcom/tl/internal/utils"
	"github.com/vkcom/tl/pkg/basictl"
	"math/rand"
	"strings"
	"testing"
)

var bannedTypes = []string{"cycleTuple"}

func TestExactCases(t *testing.T) {
	// take them from results of testing below
	const TestingName = "useResponse"
	const TestingBytes = "00000002_1ec6a63e_944aaa97_312e676c"

	sizeBuffer := make([]int, 100)
	writeBuffer := make([]byte, 100)

	testingBytesChanged := strings.ReplaceAll(TestingBytes, "_", " ")

	dst := factory.CreateObjectFromName(TestingName)
	_, _ = dst.Read(utils.ParseHexToBytes(testingBytesChanged))

	json, _ := dst.WriteJSONGeneral(nil)
	fmt.Println(string(json))

	writeBuffer, sizeBuffer = dst.WriteTL2(writeBuffer[0:0], sizeBuffer[0:0])

	newDst := factory.CreateObjectFromName(TestingName)
	_, err := newDst.ReadTL2(writeBuffer)
	if err != nil {
		t.Fatalf("can't readTL2")
	}
	newData, err := newDst.WriteGeneral(nil)
	if err != nil {
		t.Fatalf("can't write")
	}
	assert.Equal(t, testingBytesChanged, utils.SprintHexDump(newData))
}

func TestGoldmasterTL2Random(t *testing.T) {
	const NumberOfSamples = 20
	seed := rand.Int63()
	fmt.Printf("Seed: %d\n", seed)

	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(seed)))
	allItems := meta.GetAllTLItems()

	sizeBuffer := make([]int, 100)
	writeBuffer := make([]byte, 100)

	bannedSet := utils.SliceToSet(bannedTypes)

	for _, item := range allItems {
		if _, ok := bannedSet[item.TLName()]; ok {
			continue
		}
		t.Run(item.TLName(), func(t *testing.T) {
			dst := factory.CreateObject(item.TLTag())
			if dst == nil {
				t.Fatalf("can't init %s", item.TLName())
			}

			for i := 0; i < NumberOfSamples; i++ {
				dst.FillRandom(rg)
				data, err := dst.WriteGeneral(nil)
				if err != nil {
					t.Fatalf("can't seriliaze %d-th object", i)
				}
				t.Run(fmt.Sprintf("TL[%s]", utils.SprintHexDump(data)), func(t *testing.T) {
					writeBuffer, sizeBuffer = dst.WriteTL2(writeBuffer[0:0], sizeBuffer[0:0])
					newDst := factory.CreateObject(item.TLTag())
					_, err = newDst.ReadTL2(writeBuffer)
					if err != nil {
						t.Fatalf("can't readTL2 %d-th object", i)
					}
					newData, err := newDst.WriteGeneral(nil)
					if err != nil {
						t.Fatalf("can't write %d-th object", i)
					}
					assert.Equal(t, utils.SprintHexDump(data), utils.SprintHexDump(newData))
				})
			}
		})
	}
}
