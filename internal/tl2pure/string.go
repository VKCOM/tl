package tl2pure

import (
	"cmp"
	"fmt"
	"math/rand"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceString struct {
	TypeInstanceCommon
}

type KernelValueString struct {
	value string
}

func (ins *TypeInstanceString) GoodForMapKey() bool {
	return true
}

func (ins *TypeInstanceString) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstanceString) CreateValue() KernelValue {
	return &KernelValueString{}
}

func (ins *TypeInstanceString) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (v *KernelValueString) Reset() {
	v.value = ""
}

func (v *KernelValueString) Random(rg *rand.Rand) {
	count := 0
	if (rg.Uint32() & 3) != 0 { // many strings empty
		count = 1 + rg.Intn(8)
	}
	res := make([]byte, count)
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	for i := range res {
		res[i] = letters[rg.Int()%len(letters)]
	}
	v.value = string(res)
}

func (v *KernelValueString) WriteTL2(w []byte, optimizeEmpty bool, ctx *TL2Context) []byte {
	if optimizeEmpty && len(v.value) == 0 {
		return w
	}
	return basictl.StringWriteTL2(w, v.value)
}

func (v *KernelValueString) ReadTL2(r []byte, ctx *TL2Context) ([]byte, error) {
	return basictl.StringReadTL2(r, &v.value)
}

func (v *KernelValueString) WriteJSON(w []byte, ctx *TL2Context) []byte {
	return basictl.JSONWriteString(w, v.value)
}

func (v *KernelValueString) UIWrite(sb *strings.Builder, onPath bool, level int, path []int, model *UIModel) {
	w := basictl.JSONWriteString(nil, v.value)
	sb.Write(w)
}

func (v *KernelValueString) UIFixPath(side int, level int, model *UIModel) int {
	model.Path = model.Path[:level]
	return 0
}

func (v *KernelValueString) UIStartEdit(level int, model *UIModel) {
}

func (v *KernelValueString) Clone() KernelValue {
	return &KernelValueString{v.value}
}

func (v *KernelValueString) CompareForMapKey(other KernelValue) int {
	if v2, ok := other.(*KernelValueString); ok {
		return cmp.Compare(v.value, v2.value)
	}
	return 0
}

func (k *Kernel) addString() {
	name := "string"
	comb := tlast.TL2Combinator{
		TypeDecl: tlast.TL2TypeDeclaration{
			Name: tlast.TL2TypeName{Name: name},
			Type: tlast.TL2TypeDefinition{IsConstructorFields: true}, // for the purpose of type check, this is object with no fields
		},
	}
	ins := TypeInstanceString{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: name,
			comb:          comb,
		},
	}
	ref := &TypeInstanceRef{
		ins: &ins,
	}
	kt := &KernelType{
		comb:      comb,
		instances: map[string]*TypeInstanceRef{name: ref},
	}
	if _, ok := k.instances[name]; ok {
		panic(fmt.Sprintf("error adding primitive type %s: exist in global list", name))
	}
	if err := k.addTip(kt); err != nil {
		panic(fmt.Sprintf("error adding primitive type %s: %v", name, err))
	}
	k.instances[name] = ref
	// k.instancesOrdered = append(k.instancesOrdered, ref) - we do not yet know if we need them here
}
