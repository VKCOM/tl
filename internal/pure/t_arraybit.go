package pure

import (
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceArrayBit struct {
	TypeInstanceCommon
	isTuple bool
	count   int
}

func (ins *TypeInstanceArrayBit) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstanceArrayBit) CreateValue() KernelValue {
	value := &KernelValueArrayBit{
		instance: ins,
	}
	if ins.isTuple {
		value.resize(ins.count)
	}
	return value
}

func (ins *TypeInstanceArrayBit) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}
