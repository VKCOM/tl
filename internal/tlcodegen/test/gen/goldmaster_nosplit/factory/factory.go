// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package factory

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster_nosplit/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster_nosplit/meta"
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
	meta.SetGlobalFactoryCreateForEnumElement(0x623360f3)
	meta.SetGlobalFactoryCreateForEnumElement(0xf35d7a69)
	meta.SetGlobalFactoryCreateForEnumElement(0x6127e7b8)
	meta.SetGlobalFactoryCreateForEnumElement(0xb83a723d)
	meta.SetGlobalFactoryCreateForObject(0x7082d18f, func() meta.Object { var ret internal.ATop2; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xa7662843, func() meta.Object { var ret internal.AUNionA; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x944aaa97, func() meta.Object { var ret internal.AbAlias; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x20c5fb2d, func() meta.Object { var ret internal.AbCall1; return &ret }, func() meta.Function { var ret internal.AbCall1; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x8db2a4f8, func() meta.Object { var ret internal.AbCall10; return &ret }, func() meta.Function { var ret internal.AbCall10; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xecb2a36c, func() meta.Object { var ret internal.AbCall11; return &ret }, func() meta.Function { var ret internal.AbCall11; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x77d5f057, func() meta.Object { var ret internal.AbCall2; return &ret }, func() meta.Function { var ret internal.AbCall2; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x0a083445, func() meta.Object { var ret internal.AbCall3; return &ret }, func() meta.Function { var ret internal.AbCall3; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xc1220a1e, func() meta.Object { var ret internal.AbCall4; return &ret }, func() meta.Function { var ret internal.AbCall4; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x7ba4d28d, func() meta.Object { var ret internal.AbCall5; return &ret }, func() meta.Function { var ret internal.AbCall5; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x84d815cb, func() meta.Object { var ret internal.AbCall6; return &ret }, func() meta.Function { var ret internal.AbCall6; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x46ec10bf, func() meta.Object { var ret internal.AbCall7; return &ret }, func() meta.Function { var ret internal.AbCall7; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x1b8652d9, func() meta.Object { var ret internal.AbCall8; return &ret }, func() meta.Function { var ret internal.AbCall8; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x75de906c, func() meta.Object { var ret internal.AbCall9; return &ret }, func() meta.Function { var ret internal.AbCall9; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x7651b1ac, func() meta.Object { var ret internal.AbCode; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x14a35d80, func() meta.Object { var ret internal.AbCounterChangeRequestPeriodsMany; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xd9c36de5, func() meta.Object { var ret internal.AbCounterChangeRequestPeriodsOne; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x1ec6a63e, func() meta.Object { var ret internal.AbEmpty; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xe0e96c86, func() meta.Object { var ret internal.AbMyType; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x4dac492a, func() meta.Object { var ret internal.AbTestMaybe; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xe67bce28, func() meta.Object { var ret internal.AbTopLevel1; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xcef933fb, func() meta.Object { var ret internal.AbTopLevel2; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xa99fef6a, func() meta.Object { var ret internal.AbTypeA; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xff2e6d58, func() meta.Object { var ret internal.AbTypeB; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x69920d6e, func() meta.Object { var ret internal.AbTypeC; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x76615bf1, func() meta.Object { var ret internal.AbTypeD; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x71687381, func() meta.Object { var ret internal.AbUseCycle; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x3325d884, func() meta.Object { var ret internal.AbUseDictString; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xdf61f632, func() meta.Object { var ret internal.AuNionA; return &ret })
	meta.SetGlobalFactoryCreateForEnumElement(0xa9471844)
	meta.SetGlobalFactoryCreateForFunction(0xa7302fbc, func() meta.Object { var ret internal.Call1; return &ret }, func() meta.Function { var ret internal.Call1; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xf02024c6, func() meta.Object { var ret internal.Call2; return &ret }, func() meta.Function { var ret internal.Call2; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x6ace6718, func() meta.Object { var ret internal.Call3; return &ret }, func() meta.Function { var ret internal.Call3; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x46d7de8f, func() meta.Object { var ret internal.Call4; return &ret }, func() meta.Function { var ret internal.Call4; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xfc51061c, func() meta.Object { var ret internal.Call5; return &ret }, func() meta.Function { var ret internal.Call5; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xe41e4696, func() meta.Object { var ret internal.Call6; return &ret }, func() meta.Function { var ret internal.Call6; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x262a43e2, func() meta.Object { var ret internal.Call7; return &ret }, func() meta.Function { var ret internal.Call7; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x7b400184, func() meta.Object { var ret internal.Call8; return &ret }, func() meta.Function { var ret internal.Call8; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x67a0d62d, func() meta.Object { var ret internal.Call9; return &ret }, func() meta.Function { var ret internal.Call9; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0xeab6a6b4, func() meta.Object { var ret internal.CdMyType; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x8c202f64, func() meta.Object { var ret internal.CdResponse; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x5cd1ca89, func() meta.Object { var ret internal.CdTopLevel3; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xa831a920, func() meta.Object { var ret internal.CdTypeA; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x377b4996, func() meta.Object { var ret internal.CdTypeB; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xdb0f93d4, func() meta.Object { var ret internal.CdTypeC; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xb5528285, func() meta.Object { var ret internal.CdTypeD; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x6ed67ca0, func() meta.Object { var ret internal.CdUseCycle; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x136ecc9e, func() meta.Object { var ret internal.Cyc1MyCycle; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xfba5eecb, func() meta.Object { var ret internal.Cyc2MyCycle; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x47866860, func() meta.Object { var ret internal.Cyc3MyCycle; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xc867fae3, func() meta.Object { var ret internal.CycleTuple; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x647ddaf5, func() meta.Object { var ret internal.HalfStr; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x12ab5219, func() meta.Object { var ret internal.Hren; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xa8509bda, func() meta.Object { var ret internal.Int; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x7934e71f, func() meta.Object { var ret internal.Int32; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xf5609de0, func() meta.Object { var ret internal.Int64; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x22076cba, func() meta.Object { var ret internal.Long; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xc457763c, func() meta.Object { var ret internal.MaybeTest1; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x0e1ae81e, func() meta.Object { var ret internal.MultiPoint; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xba59e151, func() meta.Object { var ret internal.MyInt32; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x1d95db9d, func() meta.Object { var ret internal.MyInt64; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xc60c1b41, func() meta.Object { var ret internal.MyNat; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x79e0c6df, func() meta.Object { var ret internal.MyPlus; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x692c291b, func() meta.Object { var ret internal.MyPlus3; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x8d868379, func() meta.Object { var ret internal.MyZero; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x103a40cf, func() meta.Object { var ret internal.MyZero3; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x344ddf50, func() meta.Object { var ret internal.NativeWrappers; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x3a728324, func() meta.Object { var ret internal.NoStr; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xebb26b29, func() meta.Object { var ret internal.Replace; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xf46f9b9b, func() meta.Object { var ret internal.Replace17; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xe2d4ebee, func() meta.Object { var ret internal.Replace2; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x51e324e4, func() meta.Object { var ret internal.Replace3; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x8b5bc78a, func() meta.Object { var ret internal.Replace5; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xabd49d06, func() meta.Object { var ret internal.Replace6; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xff8f7db8, func() meta.Object { var ret internal.Service5EmptyOutput; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x7cf362ba, func() meta.Object { var ret internal.Service5Insert; return &ret }, func() meta.Function { var ret internal.Service5Insert; return &ret }, func() meta.Function { var ret internal.Service5LongInsert; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x7cf362bc, func() meta.Object { var ret internal.Service5InsertList; return &ret }, func() meta.Function { var ret internal.Service5InsertList; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0xff8f7db9, func() meta.Object { var ret internal.Service5LongEmptyOutput; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x7cf362bb, func() meta.Object { var ret internal.Service5LongInsert; return &ret }, func() meta.Function { var ret internal.Service5LongInsert; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0xdc170ff5, func() meta.Object { var ret internal.Service5LongStringOutput; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xdc170ff4, func() meta.Object { var ret internal.Service5StringOutput; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xb5286e24, func() meta.Object { var ret internal.String; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x88920e90, func() meta.Object { var ret internal.TestMaybe; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x0aa03cf2, func() meta.Object { var ret internal.TestMaybe2; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x3fedd339, func() meta.Object { var ret internal.True; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x157673c1, func() meta.Object { var ret internal.TypeA; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x9d024802, func() meta.Object { var ret internal.TypeB; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x6b8ef43f, func() meta.Object { var ret internal.TypeC; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xb1f4369e, func() meta.Object { var ret internal.TypeD; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x742161d2, func() meta.Object { var ret internal.UnionArgsUse; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xfb9ce817, func() meta.Object { var ret internal.UseDictUgly; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x0a63ec5f, func() meta.Object { var ret internal.UseResponse; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x9aa3dee5, func() meta.Object { var ret internal.UseStr; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xdfdd4180, func() meta.Object { var ret internal.UseTrue; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x3c857e52, func() meta.Object { var ret internal.UsefulServiceGetUserEntity; return &ret }, func() meta.Function { var ret internal.UsefulServiceGetUserEntity; return &ret }, nil)
}
