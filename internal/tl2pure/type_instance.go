package tl2pure

import "github.com/vkcom/tl/internal/tlast"

type TypeInstance interface {
	CanonicalName() string
	Declaration() tlast.TL2TypeDeclaration

	GoodForMapKey() bool
	IsBit() bool // for vector/tuple special case
	FindCycle(c *cycleFinder)

	CreateValue() KernelValue
	SkipTL2(r []byte) ([]byte, error)
}

// during recursive type resolution, we store pointer to this type,
// later type instance is instantiated and ins is set
type TypeInstanceRef struct {
	ins TypeInstance
}

type TypeInstanceCommon struct {
	canonicalName string
	declaration   tlast.TL2TypeDeclaration
}

func (ins *TypeInstanceCommon) CanonicalName() string {
	return ins.canonicalName
}

func (ins *TypeInstanceCommon) Declaration() tlast.TL2TypeDeclaration {
	return ins.declaration
}

func (ins *TypeInstanceCommon) GoodForMapKey() bool {
	return false
}

func (ins *TypeInstanceCommon) IsBit() bool {
	return false
}
