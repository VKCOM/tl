package tl2pure

import "github.com/vkcom/tl/internal/tlast"

type TypeInstance interface {
	CanonicalName() string
	Declaration() tlast.TL2TypeDeclaration

	GoodForMapKey() bool
	FindCycle(c *cycleFinder)

	CreateValue() KernelValue
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

type TypeInstancePrimitive struct {
	TypeInstanceCommon
	goodForMapKey bool
	clone         KernelValue
}

func (ins *TypeInstancePrimitive) GoodForMapKey() bool {
	return ins.goodForMapKey
}

func (ins *TypeInstancePrimitive) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstancePrimitive) CreateValue() KernelValue {
	return ins.clone.Clone()
}
