package tl2compatibility

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vkcom/tl/internal/utils"
	"github.com/vkcom/tl/pkg/basictl"
	"math/rand"

	factory1 "github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/factory"
	meta1 "github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/meta"

	factory2 "github.com/vkcom/tl/internal/tlcodegen/test/gen/casesTL2/factory"
	meta2 "github.com/vkcom/tl/internal/tlcodegen/test/gen/casesTL2/meta"

	"sort"
	"strings"
	"testing"
)

func TestJsonCompatibility(t *testing.T) {
	tl1Items := meta1.GetAllTLItems()
	tl1ItemsNames := utils.MapSlice(tl1Items, func(a meta1.TLItem) string {
		return a.TLName()
	})

	sort.Slice(tl1Items, func(i, j int) bool {
		return strings.Compare(tl1Items[i].TLName(), tl1Items[j].TLName()) > 0
	})

	tl2Items := meta2.GetAllTLItems()
	tl2ItemsNames := utils.MapSlice(tl2Items, func(a meta2.TLItem) string {
		return a.TLName()
	})

	commonNames := utils.SetIntersection(utils.SliceToSet(tl1ItemsNames), utils.SliceToSet(tl2ItemsNames))
	if len(commonNames) == 0 {
		t.Fatalf("no interscation between generated")
		return
	}

	seed := int64(rand.Uint64())
	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(seed)))

	for _, item := range tl1Items {
		t.Run(item.TLName(), func(t *testing.T) {
			tl1Obj := factory1.CreateObjectFromName(item.TLName())
			if tl1Obj == nil {
				t.Fatalf("can't create in TL1: %s", item.TLName())
				return
			}
			tl2Obj := factory2.CreateObjectFromName(item.TLName())
			if tl2Obj == nil {
				t.Skipf("for some reason can't find in tl2: %s", item.TLName())
				return
			}
			const NumberOfRuns = 10
			for i := 0; i < NumberOfRuns; i++ {
				t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
					tl1Obj.FillRandom(rg)
					jsonValue, err := tl1Obj.WriteJSONGeneral(&basictl.JSONWriteContext{}, nil)
					if err != nil {
						t.Fatalf("can't tl1 write json: %s", err)
						return
					}

					err = tl2Obj.ReadJSON(false, &basictl.JsonLexer{Data: jsonValue})
					//err = tl2Obj.ReadJSON(false, &basictl.JsonLexer{Data: []byte(`{smth}`)})

					if err != nil {
						t.Logf("[info] json value: %s", jsonValue)
						t.Fatalf("can't read tl2 from json: %s", err)
						return
					}

					newJsonValue, err := tl1Obj.WriteJSONGeneral(&basictl.JSONWriteContext{}, nil)
					if err != nil {
						t.Fatalf("can't tl2 write json: %s", err)
						return
					}
					assert.Equal(t, jsonValue, newJsonValue)
				})
			}
		})
	}

	if t.Failed() {
		t.Logf("seed: %d", seed)
	}
}

func TestTL2Compatibility(t *testing.T) {
	tl1Items := meta1.GetAllTLItems()
	tl1ItemsNames := utils.MapSlice(tl1Items, func(a meta1.TLItem) string {
		return a.TLName()
	})

	sort.Slice(tl1Items, func(i, j int) bool {
		return strings.Compare(tl1Items[i].TLName(), tl1Items[j].TLName()) > 0
	})

	tl2Items := meta2.GetAllTLItems()
	tl2ItemsNames := utils.MapSlice(tl2Items, func(a meta2.TLItem) string {
		return a.TLName()
	})

	commonNames := utils.SetIntersection(utils.SliceToSet(tl1ItemsNames), utils.SliceToSet(tl2ItemsNames))
	if len(commonNames) == 0 {
		t.Fatalf("no interscation between generated")
		return
	}

	seed := int64(rand.Uint64())
	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(seed)))

	for _, item := range tl1Items {
		t.Run(item.TLName(), func(t *testing.T) {
			tl1Obj := factory1.CreateObjectFromName(item.TLName())
			if tl1Obj == nil {
				t.Fatalf("can't create in TL1: %s", item.TLName())
				return
			}
			tl2Obj := factory2.CreateObjectFromName(item.TLName())
			if tl2Obj == nil {
				t.Skipf("for some reason can't find in tl2: %s", item.TLName())
				return
			}
			buffer := make([]int, 1000)
			ctx := basictl.TL2WriteContext{SizeBuffer: buffer}

			const NumberOfRuns = 10
			for i := 0; i < NumberOfRuns; i++ {
				t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
					tl1Obj.FillRandom(rg)

					tl2Value := tl1Obj.WriteTL2(nil, &ctx)
					//tl2Value := utils.ParseHexToBytesTL2(`{smth}`)

					_, err := tl2Obj.ReadTL2(tl2Value, &basictl.TL2ReadContext{})
					if err != nil {
						t.Logf("[info] tl2 value: %s", utils.SprintHexDumpTL2(tl2Value))
						t.Fatalf("can't read tl2 from tl2: %s", err)
						return
					}

					newTl2Value := tl1Obj.WriteTL2(nil, &ctx)
					assert.Equal(t, utils.SprintHexDumpTL2(tl2Value), utils.SprintHexDumpTL2(newTl2Value))
				})
			}
		})
	}

	if t.Failed() {
		t.Logf("seed: %d", seed)
	}
}

func TestJsonTL2Compatibility(t *testing.T) {
	tl1Items := meta1.GetAllTLItems()
	tl1ItemsNames := utils.MapSlice(tl1Items, func(a meta1.TLItem) string {
		return a.TLName()
	})

	sort.Slice(tl1Items, func(i, j int) bool {
		return strings.Compare(tl1Items[i].TLName(), tl1Items[j].TLName()) > 0
	})

	tl2Items := meta2.GetAllTLItems()
	tl2ItemsNames := utils.MapSlice(tl2Items, func(a meta2.TLItem) string {
		return a.TLName()
	})

	commonNames := utils.SetIntersection(utils.SliceToSet(tl1ItemsNames), utils.SliceToSet(tl2ItemsNames))
	if len(commonNames) == 0 {
		t.Fatalf("no interscation between generated")
		return
	}

	seed := int64(rand.Uint64())
	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(seed)))

	for _, item := range tl1Items {
		t.Run(item.TLName(), func(t *testing.T) {
			tl1Obj := factory1.CreateObjectFromName(item.TLName())
			if tl1Obj == nil {
				t.Fatalf("can't create in TL1: %s", item.TLName())
				return
			}
			tl2Obj := factory2.CreateObjectFromName(item.TLName())
			if tl2Obj == nil {
				t.Skipf("for some reason can't find in tl2: %s", item.TLName())
				return
			}
			buffer := make([]int, 1000)
			ctx := basictl.TL2WriteContext{SizeBuffer: buffer}

			const NumberOfRuns = 10
			for i := 0; i < NumberOfRuns; i++ {
				t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
					tl1Obj.FillRandom(rg)
					// exact value
					//_ = tl1Obj.ReadJSON(false, &basictl.JsonLexer{Data: []byte(`{"f1":1,"f2":2,"f3":true,"f4":true}`)})

					// o1 -> TL1_old
					tl1Value, _ := tl1Obj.WriteGeneral(nil)
					// o1 -> JSON
					jsonValue, err := tl1Obj.WriteJSONGeneral(&basictl.JSONWriteContext{}, nil)
					if err != nil {
						t.Fatalf("can't tl1 write json: %s", err)
						return
					}

					// JSON -> o2
					err = tl2Obj.ReadJSON(false, &basictl.JsonLexer{Data: jsonValue})
					//err = tl2Obj.ReadJSON(false, &basictl.JsonLexer{Data: []byte(`{smth}`)})

					if err != nil {
						t.Logf("[info] json value: %s", jsonValue)
						t.Fatalf("can't read tl2 from json: %s", err)
						return
					}

					// o2 -> TL2
					tl2Value := tl2Obj.WriteTL2(nil, &ctx)
					//tl2Value := utils.ParseHexToBytesTL2(`{smth}`)

					// TL2 -> o1
					_, err = tl1Obj.ReadTL2(tl2Value, &basictl.TL2ReadContext{})
					if err != nil {
						t.Logf("[info] json value: %s", jsonValue)
						t.Logf("[info] tl2 value: %s", utils.SprintHexDumpTL2(tl2Value))
						t.Fatalf("can't read tl2 from tl2: %s", err)
						return
					}

					// o1 -> TL1_new
					newTl1Value, _ := tl1Obj.WriteGeneral(nil)
					// TL1_old =?= TL2_new
					if !assert.Equal(t, tl1Value, newTl1Value) {
						newJsonValue, _ := tl1Obj.WriteJSONGeneral(&basictl.JSONWriteContext{}, nil)
						t.Logf("[info] old/new json represenation:\n\t%s\n\t%s", jsonValue, newJsonValue)
					}
				})
			}
		})
	}

	if t.Failed() {
		t.Logf("seed: %d", seed)
	}
}
