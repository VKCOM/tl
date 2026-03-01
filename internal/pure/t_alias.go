// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

// commented for now. decided to use TypeStruct with isAlias true for now
//type TypeInstanceAlias struct {
//	TypeInstanceCommon
//	fieldType *TypeInstanceRef
//	fieldBare bool // TODO - actually use it
//}
//
//func (ins *TypeInstanceAlias) GoodForMapKey() bool {
//	return ins.fieldType.ins.GoodForMapKey()
//}
//
//func (ins *TypeInstanceAlias) IsBit() bool {
//	return ins.fieldType.ins.IsBit()
//}
//
//func (ins *TypeInstanceAlias) FindCycle(c *cycleFinder) {
//	if !c.push(ins) {
//		return
//	}
//	defer c.pop(ins)
//	ins.fieldType.ins.FindCycle(c)
//}
//
//func (ins *TypeInstanceAlias) GetChildren(children []TypeInstance, withReturnType bool) []TypeInstance {
//	return append(children, ins.fieldType.ins)
//}
//
//func (ins *TypeInstanceAlias) CreateValue() KernelValue {
//	value := &KernelValueAlias{
//		instance: ins,
//		value:    ins.fieldType.ins.CreateValue(),
//	}
//	return value
//}
//
//func (ins *TypeInstanceAlias) SkipTL2(r []byte) ([]byte, error) {
//	return ins.fieldType.ins.SkipTL2(r)
//}
