// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package factory

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/cycle_42a1a8597f818829cd168dce9785322f"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/cycle_44515dca4b2e76ca676b13645e716786"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/cycle_4a1568ff5f665a65be83c5d14a33c0d0"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlInt32"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlInt64"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlLong"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlString"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlTrue"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlbenchmarks/tlBenchmarksVruHash"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlbenchmarks/tlBenchmarksVruPosition"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlbenchmarks/tlBenchmarksVrutoyTopLevelContainer"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlbenchmarks/tlBenchmarksVrutoyTopLevelContainerWithDependency"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestAllPossibleFieldConfigsContainer"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestArray"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestBeforeReadBitValidation"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestDictAny"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestDictInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestDictString"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestEnumContainer"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestInplaceStructArgs"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestInplaceStructArgs2"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestLocalFieldmask"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestOutFieldMaskContainer"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestRecursiveFieldmask"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestTuple"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestUnionContainer"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestVector"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestArrayFixedBool"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestArrayFlexibleBool"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestFixedParam"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestFunctionNoDep1"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestFunctionNoDep2"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestFunctionNoDep3"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestFunctionNoDep4"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestFunctionNoDep5"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestFunctionWithDep1"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestFunctionWithDep2"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestObject"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestParamsGeneration"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcasesTL2/tlCasesTL2TestVectorBool"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases_bytes/tlCasesBytesTestArray"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases_bytes/tlCasesBytesTestDictAny"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases_bytes/tlCasesBytesTestDictInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases_bytes/tlCasesBytesTestDictString"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases_bytes/tlCasesBytesTestDictStringString"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases_bytes/tlCasesBytesTestEnumContainer"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases_bytes/tlCasesBytesTestTuple"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases_bytes/tlCasesBytesTestVector"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/meta"
)

func CreateFunction(tag uint32) meta.Function {
	return meta.CreateFunction(tag)
}

func CreateObject(tag uint32) meta.Object {
	return meta.CreateObject(tag)
}

// name can be in any of 3 forms "ch_proxy.insert#7cf362ba", "ch_proxy.insert" or "#7cf362ba"
func CreateFunctionFromName(name string) meta.Function {
	return meta.CreateFunctionFromName(name)
}

// name can be in any of 3 forms "ch_proxy.insert#7cf362ba", "ch_proxy.insert" or "#7cf362ba"
func CreateObjectFromName(name string) meta.Object {
	return meta.CreateObjectFromName(name)
}

func init() {
	meta.SetGlobalFactoryCreateForObject(0xd31bd0fd, func() meta.Object { var ret tlBenchmarksVruHash.BenchmarksVruHash; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x32792c04, func() meta.Object { var ret tlBenchmarksVruPosition.BenchmarksVruPosition; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xfb442ca5, func() meta.Object {
		var ret tlBenchmarksVrutoyTopLevelContainer.BenchmarksVrutoyTopLevelContainer
		return &ret
	})
	meta.SetGlobalFactoryCreateForObject(0xc176008e, func() meta.Object {
		var ret tlBenchmarksVrutoyTopLevelContainerWithDependency.BenchmarksVrutoyTopLevelContainerWithDependency
		return &ret
	})
	meta.SetGlobalFactoryCreateForObject(0xef556bee, func() meta.Object {
		var ret cycle_4a1568ff5f665a65be83c5d14a33c0d0.BenchmarksVrutoytopLevelUnionBig
		return &ret
	})
	meta.SetGlobalFactoryCreateForObject(0xce27c770, func() meta.Object {
		var ret cycle_4a1568ff5f665a65be83c5d14a33c0d0.BenchmarksVrutoytopLevelUnionEmpty
		return &ret
	})
	meta.SetGlobalFactoryCreateForObject(0x3762fb81, func() meta.Object { var ret tlCasesBytesTestArray.CasesBytesTestArray; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x5a5fce57, func() meta.Object { var ret tlCasesBytesTestDictAny.CasesBytesTestDictAny; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x453ace07, func() meta.Object { var ret tlCasesBytesTestDictInt.CasesBytesTestDictInt; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x6c04d6ce, func() meta.Object { var ret tlCasesBytesTestDictString.CasesBytesTestDictString; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xad69c772, func() meta.Object {
		var ret tlCasesBytesTestDictStringString.CasesBytesTestDictStringString
		return &ret
	})
	meta.SetGlobalFactoryCreateForEnumElement(0x58aad3f5)
	meta.SetGlobalFactoryCreateForEnumElement(0x00b47add)
	meta.SetGlobalFactoryCreateForEnumElement(0x81911ffa)
	meta.SetGlobalFactoryCreateForObject(0x32b92037, func() meta.Object { var ret tlCasesBytesTestEnumContainer.CasesBytesTestEnumContainer; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x2dd3bacf, func() meta.Object { var ret tlCasesBytesTestTuple.CasesBytesTestTuple; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x3647c8ae, func() meta.Object { var ret tlCasesBytesTestVector.CasesBytesTestVector; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xd3ca919d, func() meta.Object { var ret cycle_44515dca4b2e76ca676b13645e716786.CasesMyCycle1; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x5444c9a2, func() meta.Object { var ret cycle_44515dca4b2e76ca676b13645e716786.CasesMyCycle2; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x7624f86b, func() meta.Object { var ret cycle_44515dca4b2e76ca676b13645e716786.CasesMyCycle3; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xf704cf4e, func() meta.Object { var ret tlCasesTL2TestArrayFixedBool.CasesTL2TestArrayFixedBool; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x974a9b29, func() meta.Object { var ret tlCasesTL2TestArrayFlexibleBool.CasesTL2TestArrayFlexibleBool; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x22c48297, func() meta.Object { var ret tlCasesTL2TestFixedParam.CasesTL2TestFixedParam; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x1b8b9feb, func() meta.Object { var ret tlCasesTL2TestFunctionNoDep1.CasesTL2TestFunctionNoDep1; return &ret }, func() meta.Function { var ret tlCasesTL2TestFunctionNoDep1.CasesTL2TestFunctionNoDep1; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x0a2c0bf9, func() meta.Object { var ret tlCasesTL2TestFunctionNoDep2.CasesTL2TestFunctionNoDep2; return &ret }, func() meta.Function { var ret tlCasesTL2TestFunctionNoDep2.CasesTL2TestFunctionNoDep2; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xf020849b, func() meta.Object { var ret tlCasesTL2TestFunctionNoDep3.CasesTL2TestFunctionNoDep3; return &ret }, func() meta.Function { var ret tlCasesTL2TestFunctionNoDep3.CasesTL2TestFunctionNoDep3; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x5a933a50, func() meta.Object { var ret tlCasesTL2TestFunctionNoDep4.CasesTL2TestFunctionNoDep4; return &ret }, func() meta.Function { var ret tlCasesTL2TestFunctionNoDep4.CasesTL2TestFunctionNoDep4; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x2b47b925, func() meta.Object { var ret tlCasesTL2TestFunctionNoDep5.CasesTL2TestFunctionNoDep5; return &ret }, func() meta.Function { var ret tlCasesTL2TestFunctionNoDep5.CasesTL2TestFunctionNoDep5; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xb6c63b07, func() meta.Object { var ret tlCasesTL2TestFunctionWithDep1.CasesTL2TestFunctionWithDep1; return &ret }, func() meta.Function { var ret tlCasesTL2TestFunctionWithDep1.CasesTL2TestFunctionWithDep1; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x9d44a2fd, func() meta.Object { var ret tlCasesTL2TestFunctionWithDep2.CasesTL2TestFunctionWithDep2; return &ret }, func() meta.Function { var ret tlCasesTL2TestFunctionWithDep2.CasesTL2TestFunctionWithDep2; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x4f96dd95, func() meta.Object { var ret tlCasesTL2TestObject.CasesTL2TestObject; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xaac2f033, func() meta.Object { var ret tlCasesTL2TestParamsGeneration.CasesTL2TestParamsGeneration; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x644bb447, func() meta.Object { var ret tlCasesTL2TestVectorBool.CasesTL2TestVectorBool; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xe3fae936, func() meta.Object {
		var ret tlCasesTestAllPossibleFieldConfigsContainer.CasesTestAllPossibleFieldConfigsContainer
		return &ret
	})
	meta.SetGlobalFactoryCreateForObject(0xa888030d, func() meta.Object { var ret tlCasesTestArray.CasesTestArray; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x9b2396db, func() meta.Object {
		var ret tlCasesTestBeforeReadBitValidation.CasesTestBeforeReadBitValidation
		return &ret
	})
	meta.SetGlobalFactoryCreateForObject(0xe29b8ae6, func() meta.Object { var ret tlCasesTestDictAny.CasesTestDictAny; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xd3877643, func() meta.Object { var ret tlCasesTestDictInt.CasesTestDictInt; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xc463c79b, func() meta.Object { var ret tlCasesTestDictString.CasesTestDictString; return &ret })
	meta.SetGlobalFactoryCreateForEnumElement(0x6c6c55ac)
	meta.SetGlobalFactoryCreateForEnumElement(0x86ea88ce)
	meta.SetGlobalFactoryCreateForEnumElement(0x69b83e2f)
	meta.SetGlobalFactoryCreateForObject(0xcb684231, func() meta.Object { var ret tlCasesTestEnumContainer.CasesTestEnumContainer; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xa9e4441e, func() meta.Object { var ret tlCasesTestInplaceStructArgs.CasesTestInplaceStructArgs; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xaa9f2480, func() meta.Object { var ret tlCasesTestInplaceStructArgs2.CasesTestInplaceStructArgs2; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xf68fd3f9, func() meta.Object { var ret tlCasesTestLocalFieldmask.CasesTestLocalFieldmask; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xd6602613, func() meta.Object { var ret tlCasesTestMaybe.CasesTestMaybe; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x1850ffe4, func() meta.Object {
		var ret tlCasesTestOutFieldMaskContainer.CasesTestOutFieldMaskContainer
		return &ret
	})
	meta.SetGlobalFactoryCreateForObject(0xc58cf85e, func() meta.Object { var ret tlCasesTestRecursiveFieldmask.CasesTestRecursiveFieldmask; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x4b9caf8f, func() meta.Object { var ret tlCasesTestTuple.CasesTestTuple; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x4b4f09b1, func() meta.Object { var ret cycle_42a1a8597f818829cd168dce9785322f.CasesTestUnion1; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x464f96c4, func() meta.Object { var ret cycle_42a1a8597f818829cd168dce9785322f.CasesTestUnion2; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x4497a381, func() meta.Object { var ret tlCasesTestUnionContainer.CasesTestUnionContainer; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x4975695c, func() meta.Object { var ret tlCasesTestVector.CasesTestVector; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xa8509bda, func() meta.Object { var ret tlInt.Int; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x7934e71f, func() meta.Object { var ret tlInt32.Int32; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xf5609de0, func() meta.Object { var ret tlInt64.Int64; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x22076cba, func() meta.Object { var ret tlLong.Long; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xb5286e24, func() meta.Object { var ret tlString.String; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x3fedd339, func() meta.Object { var ret tlTrue.True; return &ret })
}
