// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package casetests

import (
	"encoding/json"
	"testing"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/tlbenchmarks"
	"github.com/vkcom/tl/pkg/basictl"
)

type array = []interface{}
type object = map[string]interface{}

func initJsonValue(volume int) []byte {
	positions := array{}
	for i := 0; i < volume; i++ {
		positions = append(positions, object{
			"commit_bit":     true,
			"meta_block":     true,
			"split_payload":  true,
			"rotation_block": true,
			"canonical_hash": true,
			"payload_offset": i * 100,
			"hash": object{
				"low":  i,
				"high": 2 * i,
			},
			"file_offset": i * i,
			"seq_number":  i,
		})
	}

	source := object{
		"value": object{
			"type": "benchmarks.vrutoytopLevelUnionBig",
			"value": object{
				"next_positions": positions,
			},
		},
	}

	bs, _ := json.Marshal(source)
	return bs
}

func initJsonValueWithDependency(volume int) []byte {
	positions := array{}
	for i := 0; i < volume; i++ {
		positions = append(positions, object{
			"commit_bit":     true,
			"meta_block":     true,
			"split_payload":  true,
			"rotation_block": true,
			"canonical_hash": true,
			"payload_offset": i * 100,
			"hash": object{
				"low":  i,
				"high": 2 * i,
			},
			"file_offset": i * i,
			"seq_number":  i,
		})
	}

	source := object{
		"n": volume,
		"value": object{
			"next_positions": positions,
		},
	}

	bs, _ := json.Marshal(source)
	return bs
}

const InitVolume = 100000

func BenchmarkNewRead(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	jsonValue := initJsonValue(InitVolume)
	dst := tlbenchmarks.VrutoyTopLevelContainer{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := dst.ReadJSON(true, &basictl.JsonLexer{Data: jsonValue})
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkOldRead(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	jsonValue := initJsonValue(InitVolume)
	dst := tlbenchmarks.VrutoyTopLevelContainer{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := dst.UnmarshalJSON(jsonValue)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkNewReadWithSkip(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	jsonValue := initJsonValueWithDependency(InitVolume)
	dst := tlbenchmarks.VrutoyTopLevelContainerWithDependency{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := dst.ReadJSON(true, &basictl.JsonLexer{Data: jsonValue})
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkOldReadWithSkip(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	jsonValue := initJsonValueWithDependency(InitVolume)
	dst := tlbenchmarks.VrutoyTopLevelContainerWithDependency{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := dst.UnmarshalJSON(jsonValue)
		if err != nil {
			b.Fail()
		}
	}
}
