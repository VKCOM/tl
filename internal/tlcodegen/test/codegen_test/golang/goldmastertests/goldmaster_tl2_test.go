package goldmastertests

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/factory"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/meta"
	"github.com/vkcom/tl/internal/utils"
	"github.com/vkcom/tl/pkg/basictl"
)

var bannedTypes = []string{"cycleTuple"}

func TestExactCases(t *testing.T) {
	// take them from results of testing below
	const TestingName = "maybeTest1"
	const TestingBytes = "00000001_3f9c8ef8_6c7aebef_27930a7b_3f9c8ef8_5be1ea97_69b661d9_49821d42_27930a7b_0a7d3b9e_3f9c8ef8_3b68a044_0a7d3b9e_3f9c8ef8_27930a7b_27930a7b_27930a7b_27930a7b"

	context := basictl.TL2WriteContext{SizeBuffer: make([]int, 100)}
	writeBuffer := make([]byte, 100)

	testingBytesChanged := strings.ReplaceAll(TestingBytes, "_", " ")

	dst := factory.CreateObjectFromName(TestingName)
	_, _ = dst.Read(utils.ParseHexToBytes(testingBytesChanged))

	json, _ := dst.WriteJSONGeneral(&basictl.JSONWriteContext{}, nil)
	fmt.Println(string(json))

	writeBuffer = dst.WriteTL2(writeBuffer[:0], &context)

	newDst := factory.CreateObjectFromName(TestingName)
	_, err := newDst.ReadTL2(writeBuffer, nil)
	if err != nil {
		t.Fatalf("can't readTL2: %s", err)
	}
	newData, err := newDst.WriteGeneral(nil)
	if err != nil {
		t.Fatalf("can't write")
	}
	assert.Equal(t, testingBytesChanged, utils.SprintHexDump(newData))
}

func TestGoldmasterTL2Random(t *testing.T) {
	//str := utils.ParseHexToBytes("00000000 bc799737 00000000 00000002 ce27c770 ef556bee 00000000")
	//ddst := tlcasesTL2.TestObject{}
	//str, err := ddst.Read(str)
	//if err != nil {
	//	t.Fatalf("can't read: %s", err)
	//}
	//_ = ddst.WriteTL2(nil, nil)

	const NumberOfSamples = 20
	seed := rand.Int63()
	fmt.Printf("Seed: %d\n", seed)

	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(seed)))
	allItems := meta.GetAllTLItems()

	context := basictl.TL2WriteContext{SizeBuffer: make([]int, 100)}
	writeBuffer := make([]byte, 100)

	bannedSet := utils.SliceToSet(bannedTypes)

	for _, item := range allItems {
		if _, ok := bannedSet[item.TLName()]; ok {
			continue
		}
		t.Run(item.TLName(), func(t *testing.T) {
			switch item.TLName() {
			case "useDictUgly":
				return
			}
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
				t.Run(fmt.Sprintf("TL[%s]", utils.SprintHexDump(data)), func(t *testing.T) {
					writeBuffer = dst.WriteTL2(writeBuffer[:0], &context)
					newDst := factory.CreateObjectFromName(item.TLName())
					_, err = newDst.ReadTL2(writeBuffer, nil)
					if err != nil {
						fmt.Printf("%d %s %x\n", seed, item.TLName(), writeBuffer)
						writeBuffer2 := dst.WriteTL2(nil, &context)
						fmt.Printf("%d %s %x\n", seed, item.TLName(), writeBuffer2)

						//newDst2 := factory.CreateObject(item.TLTag())
						//_, _ = newDst2.ReadTL2(writeBuffer, nil)
						t.Fatalf("can't readTL2 %d-th object", i)
					}
					newData, err := newDst.WriteGeneral(nil)
					if err != nil {
						t.Fatalf("can't write %d-th object", i)
					}
					if !assert.Equal(t, utils.SprintHexDump(data), utils.SprintHexDump(newData)) {
						newDst2 := factory.CreateObjectFromName(item.TLName())
						_, _ = newDst2.ReadTL2(writeBuffer, nil)
						t.Fatalf("write tl2 unexpected result")
					}
				})
			}
		})
	}
}
