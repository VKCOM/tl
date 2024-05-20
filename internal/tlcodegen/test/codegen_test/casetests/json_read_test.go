// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package casetests

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/tlcases_bytes"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/factory"

	"github.com/stretchr/testify/assert"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/meta"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/tlcases"
	"github.com/vkcom/tl/pkg/basictl"
)

type mappingSuccess struct {
	// absolute correct json form of same object
	goldenInput string
	// json alternatives which map to goldenInput
	alternatives []string
}

type mappingTest struct {
	// testing type object
	object meta.Object
	// json values which must success
	successes []mappingSuccess
	// json values which must fail on read
	failures []string
}

func runMappingTest(t *testing.T, mt mappingTest) {
	seed := rand.Uint64()
	rg := basictl.NewRandGenerator(rand.New(rand.NewSource(int64(seed))))

	fmt.Println("Seed: ", seed)

	for sId, success := range mt.successes {
		alternatives := success.alternatives
		if len(alternatives) == 0 {
			alternatives = append(alternatives, success.goldenInput)
		}
		for aId, alternative := range alternatives {
			t.Run(fmt.Sprintf("Object %d - Alternative %d", sId, aId), func(t *testing.T) {
				mt.object.FillRandom(rg)
				readErr := mt.object.ReadJSON(false, &basictl.JsonLexer{Data: []byte(alternative)})

				assert.Nil(t, readErr)
				writeData, writeErr := mt.object.MarshalJSON()

				assert.Nil(t, writeErr)
				assert.Equal(t, success.goldenInput, string(writeData))

				readAgainErr := mt.object.ReadJSON(false, &basictl.JsonLexer{Data: []byte(success.goldenInput)})
				assert.Nil(t, readAgainErr)

				writeAgainData, writeAgainErr := mt.object.MarshalJSON()

				assert.Nil(t, writeAgainErr)
				assert.Equal(t, success.goldenInput, string(writeAgainData))
			})

			if t.Failed() {
				return
			}
		}
	}

	for fId, failure := range mt.failures {
		t.Run(fmt.Sprintf("Failure %d", fId), func(t *testing.T) {
			mt.object.FillRandom(rg)
			readErr := mt.object.ReadJSON(false, &basictl.JsonLexer{Data: []byte(failure)})

			assert.NotNil(t, readErr)
		})

		if t.Failed() {
			return
		}
	}
}

func TestAllTLObjectsReadJsonByRandom(t *testing.T) {
	const RepeatNumber = 100

	seed := rand.Uint64()
	rnd := rand.New(rand.NewSource(int64(seed)))

	t.Logf("Seed: %d\n", seed)

	buf1 := make([]byte, 0)
	buf2 := make([]byte, 0)

	var err error

	for _, tlItem := range meta.GetAllTLItems() {
		obj := factory.CreateObject(tlItem.TLTag())
		for i := 0; i < RepeatNumber; i++ {
			buf1 = buf1[:0]
			buf2 = buf2[:0]
			t.Run(tlItem.TLName(), func(t *testing.T) {
				obj.FillRandom(basictl.NewRandGenerator(rnd))
				buf1, err = obj.WriteJSONGeneral(buf1)
				if err != nil {
					t.Logf("Seed: %d\n", seed)
					t.Fatal("first serialization wasn't succeeded", err.Error())
					return
				}
				err = obj.ReadJSON(false, &basictl.JsonLexer{Data: buf1})
				if err != nil {
					t.Logf("Seed: %d\n", seed)
					t.Fatal("first deserialization wasn't succeeded", err.Error())
					return
				}
				obj1 := obj
				buf2, err = obj.WriteJSONGeneral(buf2)
				if err != nil {
					t.Logf("Seed: %d\n", seed)
					t.Fatal("second serialization wasn't succeeded", err.Error())
					return
				}
				err = obj.ReadJSON(false, &basictl.JsonLexer{Data: buf2})
				if err != nil {
					t.Logf("Seed: %d\n", seed)
					t.Fatal("second deserialization wasn't succeeded", err.Error())
					return
				}
				obj2 := obj
				assert.Equal(t, buf1, buf2, "serializations must be same")
				assert.Equal(t, obj1, obj2, "translations to object must be equal")
				if t.Failed() {
					t.Logf("Seed: %d\n", seed)
					t.Logf("Test failed on %d iteration to %s", i, tlItem.TLName())
					return
				}
			})
		}
	}
}

func TestReadOrder(t *testing.T) {
	test := mappingTest{
		object: &tlcases.Replace7{},
		successes: []mappingSuccess{
			{
				goldenInput: "{\"n\":1,\"m\":2,\"a\":[[1,2]]}",
				alternatives: []string{
					"{\"n\": 1,\"m\": 2,\"a\": [[1,2]]}",
					"{\"n\": 1,\"m\": 2,\"a\": [[1,2]]}",
					"{\"m\": 2,\"n\": 1,\"a\": [[1,2]]}",
					"{\"a\": [[1,2]],\"n\": 1,\"m\": 2}",
					"{\"m\": 2,\"a\": [[1,2]],\"n\": 1}",
					"{\"a\": [[1,2]],\"m\": 2,\"n\": 1}",
				},
			},
		},
	}
	runMappingTest(t, test)
}

func TestArray(t *testing.T) {
	test := mappingTest{
		object: &tlcases.TestArray{},
		successes: []mappingSuccess{
			{
				goldenInput: `{}`, // empty array
				alternatives: []string{
					`{"arr":[],"n":0}`,
					`{"arr": []}`,
					`{"n": 0}`,
				},
			},
			{
				goldenInput: `{"n":3,"arr":[1,2,3]}`, // non-empty array
				alternatives: []string{
					`{"arr":[1,2,3],"n":3}`,
					`{"n":3, "arr":[1,2,3]}`,
				},
			},
		},
		failures: []string{
			`{"arr":[1,2],"n":3}`,      // non-empty array with incorrect length (less)
			`{"arr":[1,2,3,4],"n":3}`,  // non-empty array with incorrect length (more)
			`{"arr":[1,2,"3a"],"n":3}`, // non-empty array with incorrect data
		},
	}
	runMappingTest(t, test)
}

func TestArrayBytes(t *testing.T) {
	test := mappingTest{
		object: &tlcases_bytes.TestArray{},
		successes: []mappingSuccess{
			{
				goldenInput: `{}`, // empty array
				alternatives: []string{
					`{"arr":[],"n":0}`,
					`{"arr": []}`,
					`{"n": 0}`,
				},
			},
			{
				goldenInput: `{"n":3,"arr":["1","2","3"]}`, // non-empty array
				alternatives: []string{
					`{"arr":["1","2","3"],"n":3}`,
					`{"n":3, "arr":["1","2","3"]}`,
				},
			},
		},
		failures: []string{
			`{"arr":["1","2"],"n":3}`,         // non-empty array with incorrect length (less)
			`{"arr":["1","2","3","4"],"n":3}`, // non-empty array with incorrect length (more)
			`{"arr":[1,2,"3a"],"n":3}`,        // non-empty array with incorrect data
		},
	}
	runMappingTest(t, test)
}

func TestVector(t *testing.T) {
	test := mappingTest{
		object: &tlcases.TestVector{},
		successes: []mappingSuccess{
			{
				goldenInput:  `{}`, // empty vector
				alternatives: []string{`{"arr":[]}`},
			},
			{
				goldenInput: `{"arr":[1,2,3]}`, // non-empty vector
			},
		},
		failures: []string{
			`{"arr":[1,2,"3a"]}`, // non-empty vector with incorrect data
		},
	}
	runMappingTest(t, test)
}

func TestVectorBytes(t *testing.T) {
	test := mappingTest{
		object: &tlcases_bytes.TestVector{},
		successes: []mappingSuccess{
			{
				goldenInput:  `{}`, // empty vector
				alternatives: []string{`{"arr":[]}`},
			},
			{
				goldenInput: `{"arr":["1","2","3"]}`, // non-empty vector
			},
		},
		failures: []string{
			`{"arr":[1,2,"3a"]}`, // non-empty vector with incorrect data
		},
	}
	runMappingTest(t, test)
}

func TestTuple(t *testing.T) {
	test := mappingTest{
		object: &tlcases.TestTuple{},
		successes: []mappingSuccess{
			{
				goldenInput: `{"tpl":[1,2,3,4]}`,
			},
		},
		failures: []string{
			`{"tpl":[1,2]}`,        // tuple with incorrect length (less)
			`{"tpl":[1,2,3,4,5]}`,  // tuple with incorrect length (more)
			`{"tpl":[1,2,"3a",4]}`, // tuple with incorrect data
		},
	}
	runMappingTest(t, test)
}

func TestTupleBytes(t *testing.T) {
	test := mappingTest{
		object: &tlcases_bytes.TestTuple{},
		successes: []mappingSuccess{
			{
				goldenInput: `{"tpl":["1","2","3","4"]}`,
			},
		},
		failures: []string{
			`{"tpl":["1","2","3"]}`,         // tuple with incorrect length (less)
			`{"tpl":["1","2","3","4","5"]}`, // tuple with incorrect length (more)
			`{"tpl":[1,2,"3a",4]}`,          // tuple with incorrect data
		},
	}
	runMappingTest(t, test)
}

func TestDictionaryString(t *testing.T) {
	test := mappingTest{
		object: &tlcases.TestDictString{},
		successes: []mappingSuccess{
			{
				goldenInput: `{}`,
				alternatives: []string{
					`{"dict":{}}`,
				},
			},
			{
				goldenInput: `{"dict":{"k1":1,"k2":2}}`,
				alternatives: []string{
					`{"dict":{"k2":2,"k1":1}}`,
				},
			},
		},
		failures: []string{
			`{"dict":{"k1":1,"k2":"2a"}}`,
			`{"dict":{2:1,"k2":2}}`,
		},
	}
	runMappingTest(t, test)
}

func TestDictionaryStringStringBytes(t *testing.T) {
	test := mappingTest{
		object: &tlcases_bytes.TestDictStringString{},
		successes: []mappingSuccess{
			{
				goldenInput: `{}`,
				alternatives: []string{
					`{"dict":{}}`,
				},
			},
			{
				goldenInput: `{"dict":{"k1":"1","k2":"2"}}`,
				alternatives: []string{
					`{"dict":{"k2":"2","k1":"1"}}`,
				},
			},
		},
		failures: []string{
			`{"dict":{"k1":"1","k2":2}}`,
			`{"dict":{2:"1","k2":"2"}}`,
		},
	}
	runMappingTest(t, test)
}

func TestDictionaryStringBytes(t *testing.T) {
	test := mappingTest{
		object: &tlcases_bytes.TestDictString{},
		successes: []mappingSuccess{
			{
				goldenInput: `{}`,
				alternatives: []string{
					`{"dict":{}}`,
				},
			},
			{
				goldenInput: `{"dict":{"k1":1,"k2":2}}`,
				alternatives: []string{
					`{"dict":{"k2":2,"k1":1}}`,
				},
			},
		},
		failures: []string{
			`{"dict":{"k1":1,"k2":"2a"}}`,
			`{"dict":{2:1,"k2":2}}`,
		},
	}
	runMappingTest(t, test)
}

func TestDictionaryInt(t *testing.T) {
	test := mappingTest{
		object: &tlcases.TestDictInt{},
		successes: []mappingSuccess{
			{
				goldenInput: `{}`,
				alternatives: []string{
					`{"dict":{}}`,
				},
			},
			{
				goldenInput: `{"dict":{"1":1,"2":2}}`,
			},
		},
		failures: []string{
			`{"dict":{1:1,2:"2a"}}`,
			`{"dict":{"1":1,2:2}}`,
		},
	}
	runMappingTest(t, test)
}

func TestDictionaryIntBytes(t *testing.T) {
	test := mappingTest{
		object: &tlcases_bytes.TestDictInt{},
		successes: []mappingSuccess{
			{
				goldenInput: `{}`,
				alternatives: []string{
					`{"dict":{}}`,
				},
			},
			{
				goldenInput: `{"dict":{"1":1,"2":2}}`,
			},
		},
		failures: []string{
			`{"dict":{1:1,2:"2a"}}`,
			`{"dict":{"1":1,2:2}}`,
		},
	}
	runMappingTest(t, test)
}

func TestDictionaryAny(t *testing.T) {
	test := mappingTest{
		object: &tlcases.TestDictAny{},
		successes: []mappingSuccess{
			{
				goldenInput: `{}`,
				alternatives: []string{
					`{"dict":[]}`,
				},
			},
			{
				goldenInput: `{"dict":[{"key":1,"value":1},{"key":2.1,"value":2}]}`,
				alternatives: []string{
					`{"dict":[{"key":"1.0","value":1},{"key":"2.1","value":2}]}`,
				},
			},
		},
		failures: []string{
			`{"dict":[{"key":"1.0","value":1},{"key":"2.1","value":"2a"}]}`,
			`{"dict":[{"key":"1.0","value":1},{"key":"2.a","value":2}]}`,
		},
	}
	runMappingTest(t, test)
}

func TestDictionaryAnyBytes(t *testing.T) {
	test := mappingTest{
		object: &tlcases_bytes.TestDictAny{},
		successes: []mappingSuccess{
			{
				goldenInput: `{}`,
				alternatives: []string{
					`{"dict":[]}`,
				},
			},
			{
				goldenInput: `{"dict":[{"key":1,"value":1},{"key":2.1,"value":2}]}`,
				alternatives: []string{
					`{"dict":[{"key":"1.0","value":1},{"key":"2.1","value":2}]}`,
				},
			},
		},
		failures: []string{
			`{"dict":[{"key":"1.0","value":1},{"key":"2.1","value":"2a"}]}`,
			`{"dict":[{"key":"1.0","value":1},{"key":"2.a","value":2}]}`,
		},
	}
	runMappingTest(t, test)
}

func TestMaybe(t *testing.T) {
	test := mappingTest{
		object: &tlcases.TestMaybe{},
		successes: []mappingSuccess{
			{
				goldenInput: `{}`,
				alternatives: []string{
					`{"value":{"ok":false}}`,
					`{"value":{}}`,
				},
			},
			{
				goldenInput: `{"value":{"ok":true}}`,
				alternatives: []string{
					`{"value":{"ok":true, "value":0}}`, // if value is default, it will be omitted
				},
			},
			{
				goldenInput: `{"value":{"ok":true,"value":1}}`,
				alternatives: []string{
					`{"value":{"value":1}}`,
				},
			},
		},
		failures: []string{
			`{"value":{"ok":false,"value":1}}`, // ok is false but value is presented
		},
	}
	runMappingTest(t, test)
}

func TestUnion(t *testing.T) {
	test := mappingTest{
		object: &tlcases.TestUnionContainer{},
		successes: []mappingSuccess{
			{
				goldenInput: `{"value":{"type":"cases.testUnion1","value":{"value":1}}}`,
			},
			{
				goldenInput: `{"value":{"type":"cases.testUnion2","value":{"value":"2"}}}`,
			},
			{
				goldenInput: `{"value":{"type":"cases.testUnion1","value":{}}}`,
				alternatives: []string{
					`{"value":{"type":"cases.testUnion1"}}`,
				},
			},
		},
		failures: []string{
			`{"value":{"value":1}}`,
			`{"value":{}}`,
		},
	}
	runMappingTest(t, test)
}

func TestEnum(t *testing.T) {
	test := mappingTest{
		object: &tlcases.TestEnumContainer{},
		successes: []mappingSuccess{
			{
				goldenInput: `{"value":"cases.testEnum1"}`,
			},
			{
				goldenInput: `{"value":"cases.testEnum2"}`,
			},
		},
		failures: []string{
			`{"value":"cases.testEnum-1"}`, // non-existing case
			`{"value":1}`,                  // value is in another format
		},
	}
	runMappingTest(t, test)
}

func TestEnumBytes(t *testing.T) {
	test := mappingTest{
		object: &tlcases_bytes.TestEnumContainer{},
		successes: []mappingSuccess{
			{
				goldenInput: `{"value":"cases.testEnum1"}`,
			},
			{
				goldenInput: `{"value":"cases.testEnum2"}`,
			},
		},
		failures: []string{
			`{"value":"cases.testEnum-1"}`, // non-existing case
			`{"value":1}`,                  // value is in another format
		},
	}
	runMappingTest(t, test)
}

func TestLocalFieldMask(t *testing.T) {
	// dependency graph
	// f1 <--0-- f1 <--1-- f3
	//             \
	//              \--1-- f4
	test := mappingTest{
		object: &tlcases.TestLocalFieldmask{},
		successes: []mappingSuccess{
			{
				goldenInput: `{"f1":3,"f2":2,"f3":true,"f4":true}`,
				alternatives: []string{
					`{"f1":2,"f3":true,"f4":true}`,
					`{"f1":2,"f4":true}`,
					`{"f1":2,"f2":2}`,
				},
			},
			{
				goldenInput: `{"f1":2}`,
			},
			{
				goldenInput: `{"f1":1,"f2":2,"f3":true,"f4":true}`,
				alternatives: []string{
					`{"f3": true}`,
					`{"f4": true}`,
					`{"f2": 2}`,
				},
			},
			{
				goldenInput: `{"f1":1,"f2":0}`,
				alternatives: []string{
					`{"f2":0}`,
					`{"f3":false}`,
					`{"f3":false,"f4":false}`,
					`{"f1":1,"f2":0,"f3":false}`,
				},
			},
		},
		failures: []string{
			`{"f2":0,"f3":true,"f4":false}`, // dependant fields overrides fieldmask bit differently
			`{"f2":2,"f3":false}`,           // true type overwrites bit to 0, when it is 1
		},
	}
	runMappingTest(t, test)
}

func TestLocalFieldMaskRecursive(t *testing.T) {
	// dependency graph
	// f0 <--0-- f1 <--1-- f2 <--2-- t3
	//  \         \
	//   \---0--t1 \---1--t2
	test := mappingTest{
		object: &tlcases.TestRecursiveFieldmask{},
		successes: []mappingSuccess{
			{
				goldenInput: `{"f0":1,"f1":2,"f2":4,"t1":true,"t2":true,"t3":true}`,
				alternatives: []string{
					`{"t1":true,"t2":true,"t3":true}`, // only true types are presented, all fieldmasks are affected
					`{"t3":true}`,
					`{"f2":4}`, // chain of fieldmasks, presented: only third
				},
			},
			{
				goldenInput: `{"f0":1,"f1":2,"f2":0,"t1":true,"t2":true}`,
				alternatives: []string{
					`{"t3":false}`, // field recursively affects on all field masks
					`{"t3":false}`,
				},
			},
			{
				goldenInput: `{}`,
			},
			{
				goldenInput: `{"f0":15,"f1":0,"t1":true}`,
				alternatives: []string{
					`{"f0":15}`, // chain of fieldmasks, presented: only first
					`{"f0":14, "t1": true}`,
				},
			},
			{
				goldenInput: `{"f0":1,"f1":15,"f2":0,"t1":true,"t2":true}`,
				alternatives: []string{
					`{"f1":15}`,
					`{"f1":13, "t2":true}`,
				},
			},
			{
				goldenInput: `{"f0":31,"f1":2,"f2":15,"t1":true,"t2":true,"t3":true}`,
				alternatives: []string{
					`{"f0":30,"f2":15}`,
				},
			},
			{
				goldenInput: `{"f0":1,"f1":31,"f2":15,"t1":true,"t2":true,"t3":true}`,
				alternatives: []string{
					`{"f1":29,"f2":15}`,
				},
			},
			{
				goldenInput: `{"f0":25,"f1":31,"f2":15,"t1":true,"t2":true,"t3":true}`,
				alternatives: []string{
					`{"f0":24,"f1":29,"f2":15}`,
				},
			},
		},
		failures: []string{
			`{"f2":1,"t2":false}`, // different fields affect on same bit differently
			`{"t1":false, "t3":true}`,
		},
	}
	runMappingTest(t, test)
}

func TestOuterFieldMask(t *testing.T) {
	mt := mappingTest{
		object: &tlcases.TestOutFieldMaskContainer{},
		successes: []mappingSuccess{
			{
				goldenInput: `{"f":1,"inner":{"f1":1,"f3":[1]}}`, // all dependant non-true fields are presented and fieldmask is set
			},
			{
				goldenInput: `{"inner":{}}`,
				alternatives: []string{
					`{"f":0,"inner":{}}`, // fieldmask bit is zero but field is presented and same as default
					`{"f":0}`,            // fieldmask bit is zero and field is fully absent
				},
			},
			{
				goldenInput: `{"f":1,"inner":{"f1":0,"f3":[1]}}`,
				alternatives: []string{
					`{"f":1,"inner":{"f3":[1]}}`, // fieldmask is set but dependant from it field is absent
				},
			},
		},
		failures: []string{
			`{"f":1,"inner":{"f3":[1,2]}}`,       // fieldmask affects array size but size is different
			`{"f":1,"inner":{"f1":1,"f2":true}}`, // true type dependant from outer fieldmask is declared and try to affect
			`{"f":1,"inner":{"f1":1}}`,           // field with type arguments is absent and default value of it doesn't meet requirements
			`{"f":0,"inner":{"f1":1}}`,           // fieldmask bit is zero but field is presented and not same as default
		},
	}
	runMappingTest(t, mt)
}

func TestBeforeReadBitValidation(t *testing.T) {
	mt := mappingTest{
		object: &tlcases.TestBeforeReadBitValidation{},
		successes: []mappingSuccess{
			{
				goldenInput: `{"n":3,"a":[1,1,1],"b":[1,1,1]}`,
				alternatives: []string{
					`{"a":[1,1,1],"b":[1,1,1],"n":3}`,
					`{"a":[1,1,1],"b":[1,1,1]}`, // fieldmask is absent and its value must be reconstructed from other fields
				},
			},
		},
		failures: []string{
			`{"a":[1],"b":[1,1],"n":1}`,
		},
	}
	runMappingTest(t, mt)
}

func TestRecursiveTypes(t *testing.T) {
	mt := mappingTest{
		object: &tlcases.MyCycle2{},
		successes: []mappingSuccess{
			{
				goldenInput: `{}`,
				alternatives: []string{
					`{"fields_mask":0}`,
				},
			},
			{
				goldenInput: `{"fields_mask":1,"a":{}}`,
				alternatives: []string{
					`{"fields_mask":1}`,
				},
			},
			{
				goldenInput: `{"fields_mask":1,"a":{"fields_mask":1,"a":{}}}`,
				alternatives: []string{
					`{"a":{"a":{"fields_mask":0},"fields_mask":1},"fields_mask":1}`,
				},
			},
			{
				goldenInput: `{"fields_mask":1,"a":{"fields_mask":1,"a":{"fields_mask":1,"a":{"fields_mask":2}}}}`,
				alternatives: []string{
					`{"a":{"a":{"a":{"fields_mask":2},"fields_mask":0},"fields_mask":1},"fields_mask":1}`,
				},
			},
			{
				goldenInput: `{"fields_mask":1,"a":{"fields_mask":1,"a":{"fields_mask":1,"a":{}}}}`,
				alternatives: []string{
					`{"a":{"a":{"fields_mask":1},"fields_mask":1},"fields_mask":1}`,
				},
			},
		},
	}
	runMappingTest(t, mt)
}

func TestReadWithDifferentNatDependencies(t *testing.T) {
	mt := mappingTest{
		object: &tlcases.TestAllPossibleFieldConfigsContainer{},
		successes: []mappingSuccess{
			{
				goldenInput: `{"value":{"f00":239}}`,
				alternatives: []string{
					`{"outer":0,"value":{"f00":239}}`,
				},
			},
			{
				goldenInput: `{"value":{}}`,
				alternatives: []string{
					`{"outer":0,"value":{}}`,
					`{"outer":0}`,
					`{}`,
				},
			},
			{
				goldenInput: `{"value":{}}`,
				alternatives: []string{
					`{"outer":0,"value":{"f01":{}}}`,
				},
			},
			{
				goldenInput:  `{"value":{}}`,
				alternatives: []string{`{"outer":0,"value":{}}`},
			},
			{
				goldenInput: `{"value":{"local":1,"f02":[1],"f10":0}}`,
				alternatives: []string{
					`{"outer":0,"value":{"f02":[1],"local":1}}`,
				},
			},
			{
				goldenInput: `{"value":{}}`,
				alternatives: []string{
					`{"outer":0,"value":{"local":0}}`,
				},
			},
			{
				goldenInput:  `{"outer":1,"value":{"f03":[1],"f20":0}}`,
				alternatives: []string{`{"outer":1,"value":{"f03":[1]}}`},
			},
			{
				goldenInput:  `{"value":{}}`,
				alternatives: []string{`{"outer":0,"value":{}}`},
			},
			{
				goldenInput:  `{"value":{"local":1,"f02":[1],"f10":239}}`,
				alternatives: []string{`{"outer":0,"value":{"f02":[1],"f10":239,"local":1}}`},
			},
			{
				goldenInput:  `{"value":{"local":1,"f02":[1],"f10":239}}`,
				alternatives: []string{`{"outer":0,"value":{"f02":[1],"f10":239}}`},
			},
			{
				goldenInput:  `{"value":{"local":1,"f02":[1],"f10":0}}`,
				alternatives: []string{`{"outer":0,"value":{"f02":[1],"local":1}}`},
			},
			{
				goldenInput:  `{"value":{}}`,
				alternatives: []string{`{"outer":0,"value":{}}`},
			},
			{
				goldenInput:  `{"value":{"local":2,"f02":[2,2],"f11":true}}`,
				alternatives: []string{`{"outer":0,"value":{"f02":[2,2],"f11":true,"local":2}}`},
			},
			{
				goldenInput:  `{"value":{"local":3,"f02":[3,3,3],"f10":0,"f11":true}}`,
				alternatives: []string{`{"outer":0,"value":{"f02":[3,3,3],"f11":true,"local":1}}`},
			},
			{
				goldenInput:  `{"value":{"local":1,"f02":[1],"f10":0}}`,
				alternatives: []string{`{"outer":0,"value":{"f02":[1],"local":1}}`},
			},
			{
				goldenInput:  `{"value":{"local":2,"f02":[2,2],"f11":true}}`,
				alternatives: []string{`{"outer":0,"value":{"f02":[2,2],"f11":true}}`},
			},
			{
				goldenInput:  `{"value":{}}`,
				alternatives: []string{`{"outer":0,"value":{"f11":false}}`},
			},
			{
				goldenInput:  `{"value":{}}`,
				alternatives: []string{`{"outer":0,"value":{}}`},
			},
			{
				goldenInput:  `{"value":{"local":4,"f02":[4,4,4,4],"f12":[4,4,4,4]}}`,
				alternatives: []string{`{"outer":0,"value":{"f02":[4,4,4,4],"f12":[4,4,4,4],"local":4}}`},
			},
			{
				goldenInput:  `{"value":{"local":2,"f02":[2,2],"f11":true}}`,
				alternatives: []string{`{"outer":0,"value":{"f02":[2,2],"local":2}}`},
			},
			{
				goldenInput:  `{"outer":1,"value":{"local":8,"f02":[8,8,8,8,8,8,8,8],"f03":[1],"f13":[1],"f20":0}}`,
				alternatives: []string{`{"outer":1,"value":{"f02":[8,8,8,8,8,8,8,8],"f03":[1],"f13":[1],"local":8}}`},
			},
			{
				goldenInput:  `{"value":{"local":2,"f02":[2,2],"f11":true}}`,
				alternatives: []string{`{"outer":0,"value":{"f02":[2,2],"local":2}}`},
			},
			{
				goldenInput:  `{"outer":1,"value":{"f03":[1],"f20":239}}`,
				alternatives: []string{`{"outer":1,"value":{"f03":[1],"f20":239}}`},
			},
			{
				goldenInput:  `{"outer":1,"value":{"f03":[1],"f20":0}}`,
				alternatives: []string{`{"outer":1,"value":{"f03":[1]}}`},
			},
			{
				goldenInput:  `{"value":{}}`,
				alternatives: []string{`{"outer":0,"value":{}}`},
			},
			{
				goldenInput:  `{"value":{}}`,
				alternatives: []string{`{"outer":0,"value":{}}`},
			},
			{
				goldenInput:  `{"outer":4,"value":{"local":2,"f02":[2,2],"f03":[4,4,4,4],"f11":true,"f22":[2,2]}}`,
				alternatives: []string{`{"outer":4,"value":{"f02":[2,2],"f03":[4,4,4,4],"f22":[2,2],"local":2}}`},
			},
			{
				goldenInput:  `{"outer":4,"value":{"f03":[4,4,4,4],"f22":[]}}`,
				alternatives: []string{`{"outer":4,"value":{"f02":[],"f03":[4,4,4,4],"local":0}}`},
			},
			{
				goldenInput:  `{"value":{"local":2,"f02":[2,2],"f11":true}}`,
				alternatives: []string{`{"outer":0,"value":{"f02":[2,2],"local":2}}`},
			},
			{
				goldenInput:  `{"outer":8,"value":{"f03":[8,8,8,8,8,8,8,8],"f23":[8,8,8,8,8,8,8,8]}}`,
				alternatives: []string{`{"outer":8,"value":{"f03":[8,8,8,8,8,8,8,8],"f23":[8,8,8,8,8,8,8,8]}}`},
			},
			{
				goldenInput:  `{"value":{}}`,
				alternatives: []string{`{"outer":0,"value":{"f03":[]}}`},
			},
		},
		failures: []string{
			`{"outer":0,"value":{"f01":{"key":"value"}}}`,                                     // field is presented but it's not empty object
			`{"outer":0,"value":{"local":1}}`,                                                 // field is absent and default value doesn't fit nat requirements
			`{"outer":1,"value":{"f23":[1]}}`,                                                 // field is absent and default value doesn't fit nat requirements
			`{"outer":0,"value":{"f02":[2,2],"f11":false,"local":2}}`,                         // field is presented but is false while bit is 1, local fieldmask is presented
			`{"outer":0,"value":{"f02":[4,4,4,4],"local":4}}`,                                 // field is absent and default value of it doesn't fit to requirements, local fieldmask is presented
			`{"outer":0,"value":{"f02":[4,4,4,4],"f12":[3,3,3],"local":4}}`,                   // field is presented but doesn't fit to requirements, local fieldmask is presented
			`{"outer":1,"value":{"f02":[8,8,8,8,8,8,8,8],"f03":[1],"f13":[3,3,3],"local":8}}`, // field is presented but doesn't fit to requirements, local fieldmask is presented
			`{"outer":1,"value":{"f02":[8,8,8,8,8,8,8,8],"f03":[1],"local":8}}`,               // field is absent and default value of it doesn't fit to requirements, local fieldmask is presented
			`{"outer":0,"value":{"f03":[],"f10":239}}`,                                        // field is presented, external fieldmask's bit is 0
			`{"outer":2,"value":{"f02":[2,2],"f21":true}}`,                                    // field is presented, incorrect case
			`{"outer":0,"value":{"f02":[2,2],"f22":[2,2],"local":2}}`,                         // field is presented, external fieldmask's bit is 0
			`{"outer":4,"value":{"f02":[2,2],"f03":[4,4,4,4],"local":2}}`,                     // field is absent and default value of it doesn't fit to requirements, external fieldmask's bit is 1
			`{"outer":0,"value":{"f03":[8,8,8,8,8,8,8,8],"f23":[8,8,8,8,8,8,8,8]}}`,           // field is presented, external fieldmask's bit is 0
			`{"outer":8,"value":{"f03":[8,8,8,8,8,8,8,8],"f23":[3,3,3]}}`,                     // field is presented but doesn't fit to requirements, external fieldmask's bit is 1
			`{"outer":8,"value":{"f03":[8,8,8,8,8,8,8,8]}}`,                                   // field is absent and default value of it doesn't fit to requirements, external fieldmask's bit is 1
		},
	}
	runMappingTest(t, mt)
}
