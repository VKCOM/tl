package codegen_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/constants"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/tl"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/tlservice1"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/tlservice2"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/tltasks"
	"github.com/vkcom/tl/pkg/basictl"
)

// Этот файл тестирует код, сгенерированный tlgen.

// Эти тесты могут быстро устаревать и, возможно,
// в дальнейшем их можно будет чем-то заменить.
// Пока они нужны для предварительной валидации,
// которую планируется выполнить перед интеграцией кода,
// сгенерированного tlgen в более-менее серьёзное окружение.

func checkTestCase2(t *testing.T, x interface{}, y interface{}, goldStandard []byte, write func([]byte) ([]byte, error), read func([]byte) ([]byte, error)) {
	buf, err := write(nil)
	if err != nil {
		t.Fatalf("write error: %v\n", err)
	}
	if goldStandard != nil { // We skip non-deterministic writes
		require.Equal(t, goldStandard, buf)
	}
	buf, err = read(buf)
	if err != nil {
		t.Fatalf("read error: %v\n", err)
	}
	if len(buf) != 0 {
		t.Fatalf("excess bytes: %v\n", buf)
	}
	// Пока в кодогенераторе нет чётких правил для
	// nil slice -vs- empty slice, поэтому при
	// сравнении объектов считаем пустой и nil слайсы эквивалентными.
	diff := cmp.Diff(x, y, cmpopts.EquateEmpty(), cmp.AllowUnexported(tl.MyValue{}, tlservice1.Value{}, tltasks.TaskStatus{}))
	if diff != "" {
		t.Fatalf("diff:\n%s", diff)
	}
}

func checkJSONCase2(t *testing.T, x interface{}, z interface{}, jsonStandard string, write1 func([]byte) ([]byte, error), write2 func([]byte) ([]byte, error), read func(*basictl.JsonLexer) error) {
	w1, err := write1(nil)
	if err != nil {
		t.Log("Failed to JSON encode source value\n")
	}
	w2, err := write2(nil)
	if err != nil {
		t.Log("Failed to JSON encode destination value\n")
	}
	j1 := strings.TrimSpace(string(w1))
	j2 := strings.TrimSpace(string(w2))
	if jsonStandard != "" {
		require.Equal(t, jsonStandard, j1)
		if jsonStandard != j1 {
			t.Logf("json standard violated %s \n", j1)
			t.Logf("             should be %s \n", jsonStandard)
		}
	}
	if jsonStandard != "" && j1 != j2 {
		t.Logf("json dst is %s \n", j1)
		t.Logf("  should be %s \n", j2)
	}
	err = read(&basictl.JsonLexer{Data: []byte(j1)})
	if err != nil {
		t.Fatalf("read JSON error: %v\n", err)
	}
	diff := cmp.Diff(x, z, cmpopts.EquateEmpty(), cmp.AllowUnexported(tl.MyValue{}, tlservice1.Value{}, tltasks.TaskStatus{}))
	if diff != "" {
		t.Fatalf("json diff:\n%s", diff)
	}
}

func checkJSONCase3(t *testing.T, x interface{}, z interface{}, jsonStandard string, write1 func(bool, bool, []byte) ([]byte, error), write2 func(bool, bool, []byte) ([]byte, error), read func(*basictl.JsonLexer) error) {
	checkJSONCase2(t, x, z, jsonStandard, func(w []byte) ([]byte, error) {
		return write1(false, false, w)
	}, func(w []byte) ([]byte, error) {
		return write2(false, false, w)
	}, read)
}

func TestWriteArgs(t *testing.T) {
	writeBoxedVector32 := func(x []int32) func(w []byte) ([]byte, error) {
		fn := tl.BoxedVector32{X: x}
		return fn.WriteBoxedGeneral
	}

	writeBoxedVector32BoxedElem := func(x []int32) func(w []byte) ([]byte, error) {
		fn := tl.BoxedVector32BoxedElem{X: x}
		return fn.WriteBoxedGeneral
	}

	writeBoxedVector64 := func(x []int64) func(w []byte) ([]byte, error) {
		fn := tl.BoxedVector64{X: x}
		return fn.WriteBoxedGeneral
	}

	writeBoxedString := func(x string) func(w []byte) ([]byte, error) {
		fn := tl.BoxedString{X: x}
		return fn.WriteBoxedGeneral
	}

	writeBoxedInt := func(x int32) func(w []byte) ([]byte, error) {
		fn := tl.BoxedInt{X: x}
		return fn.WriteBoxedGeneral
	}

	writeBoxedTuple := func(x [3]int32) func(w []byte) ([]byte, error) {
		fn := tl.BoxedTuple{X: x}
		return fn.WriteBoxedGeneral
	}

	writeGetQueueSize := func(typeName string, queueID []int32, fieldsMask uint32) func(w []byte) ([]byte, error) {
		fn := tltasks.GetQueueSize{
			TypeName:   typeName,
			QueueId:    queueID,
			FieldsMask: fieldsMask,
		}
		return fn.WriteBoxedGeneral
	}

	writeGetAnyTask := func() func(w []byte) ([]byte, error) {
		fn := tltasks.GetAnyTask{}
		return fn.WriteBoxedGeneral
	}

	writeTaskAdd := func(typeName string, queueID []int32, task tltasks.Task) func(w []byte) ([]byte, error) {
		fn := tltasks.AddTask{
			TypeName: typeName,
			QueueId:  queueID,
			Task:     task,
		}
		return fn.WriteBoxedGeneral
	}

	writeGetQueueTypes := func(settingsMask, statsMask uint32) func(w []byte) ([]byte, error) {
		fn := tltasks.GetQueueTypes{
			SettingsMask: settingsMask,
			StatsMask:    statsMask,
		}
		return fn.WriteBoxedGeneral
	}

	writeGetTaskFromQueue := func(typeName string, queueID []int32) func(w []byte) ([]byte, error) {
		fn := tltasks.GetTaskFromQueue{
			TypeName: typeName,
			QueueId:  queueID,
		}
		return fn.WriteBoxedGeneral
	}

	writeGetMyValue := func(x tl.MyValue) func(w []byte) ([]byte, error) {
		fn := tl.GetMyValue{X: x}
		return fn.WriteBoxedGeneral
	}

	writeGetMyDictOfInt := func(x tl.MyDictOfInt) func(w []byte) ([]byte, error) {
		fn := tl.GetMyDictOfInt{X: x}
		return fn.WriteBoxedGeneral
	}

	writeGetDouble := func(x float64) func(w []byte) ([]byte, error) {
		fn := tl.GetDouble{X: x}
		return fn.WriteBoxedGeneral
	}

	writeGetFloat := func(x float32) func(w []byte) ([]byte, error) {
		fn := tl.GetFloat{X: x}
		return fn.WriteBoxedGeneral
	}

	writeGetNonOptNat := func(n uint32, xs []int32) func(w []byte) ([]byte, error) {
		fn := tl.GetNonOptNat{
			N:  n,
			Xs: xs,
		}
		return fn.WriteBoxed
	}

	tests := []struct {
		name  string
		want  string
		write func(w []byte) ([]byte, error)
	}{
		{
			"tasks.getAnyTask",
			"4a9c7dbb",
			writeGetAnyTask(),
		},

		{
			"tasks.getQueueTypes0",
			"5434457a 0000000c 00064e32",
			writeGetQueueTypes(12, 413234),
		},
		{
			"tasks.getQueueTypes1",
			"5434457a 00000000 00000000",
			writeGetQueueTypes(0, 0),
		},
		{
			"tasks.getQueueTypes2",
			"5434457a ffffffff ffffffff",
			writeGetQueueTypes(0xFFFFFFFF, 0xFFFFFFFF),
		},

		{
			"tasks.getTaskFromQueue0",
			"6a52b698 00000000 00000000",
			writeGetTaskFromQueue("", nil),
		},
		{
			"tasks.getTaskFromQueue1",
			"6a52b698 70797404 00000065 00000009 00000384 00000320 000002bc 00000258 000001f4 00000258 000002bc 00000320 00000384",
			writeGetTaskFromQueue("type", []int32{900, 800, 700, 600, 500, 600, 700, 800, 900}),
		},

		{
			"tasks.addTask0",
			"2ca073d5 6f6f6606 00726162 00000003 00000001 00000002 00000003 00000000 11bbccdd 00000001 00000004 33323103",
			writeTaskAdd(
				"foobar",
				[]int32{1, 2, 3},
				tltasks.Task{
					FieldsMask: 0,
					Flags:      0x11BBCCDD,
					Tag:        []int32{4},
					Data:       "123",
				},
			),
		},

		{
			"tasks.addTask1",
			"2ca073d5 00007801 00000001 00000003 0000000a 00000000 00000000 000100fe 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 78787878 0000000a 000927c0",
			writeTaskAdd(
				"x",
				[]int32{3},
				tltasks.Task{
					// Включение retries и deadline.
					FieldsMask: (1 << 1) | (1 << 3),
					Flags:      0,
					Data:       strings.Repeat("x", 256),
					Tag:        []int32{},
					Retries:    10,
					Deadline:   600000,
				},
			),
		},

		{
			"tasks.addTask2",
			"2ca073d5 00000000 00000001 000032e5 0000000f 000004d0 00000000 00000000 00000309 00000000 00000002 00000066 000927c0",
			writeTaskAdd(
				"",
				[]int32{13029},
				tltasks.Task{
					// Включение всех опциональных полей.
					FieldsMask:    (1 << 0) | (1 << 1) | (1 << 2) | (1 << 3),
					Flags:         1232,
					Data:          "",
					Tag:           []int32{},
					Retries:       2,
					Deadline:      600000,
					ScheduledTime: 102,
					Id:            777,
				},
			),
		},

		{
			"tasks.addTask3",
			"2ca073d5 00000000 00000000 00000001 00000000 00000000 00000000 00000000 00000000",
			writeTaskAdd(
				"",
				nil,
				tltasks.Task{
					// Включение только ID.
					FieldsMask: 1 << 0,
				},
			),
		},

		{
			"tasks.getQueueSize0",
			"d8fcda03 61206107 61206120 00000001 00000000 00000000",
			writeGetQueueSize("a a a a", []int32{0}, 0),
		},
		{
			"tasks.getQueueSize1",
			"d8fcda03 00000000 00000000 00000001",
			writeGetQueueSize("", nil, 1<<0),
		},
		{
			"tasks.getQueueSize2",
			"d8fcda03 00000000 00000000 00000007",
			writeGetQueueSize("", []int32{}, (1<<0)|(1<<1)|(1<<2)),
		},

		{
			"boxedVector32_0",
			"bbadef07 1cb5c415 00000000",
			writeBoxedVector32([]int32{}),
		},
		{
			"boxedVector32_1",
			"bbadef07 1cb5c415 00000003 00000001 00000002 00000003",
			writeBoxedVector32([]int32{1, 2, 3}),
		},

		{
			"boxedVector32BoxedElem_0",
			"591cecd4 1cb5c415 00000000",
			writeBoxedVector32BoxedElem([]int32{}),
		},
		{
			"boxedVector32BoxedElem_1",
			"591cecd4 1cb5c415 00000003 a8509bda 00000001 a8509bda 00000002 a8509bda 00000003",
			writeBoxedVector32BoxedElem([]int32{1, 2, 3}),
		},

		{
			"boxedVector64_0",
			"83659ba8 1cb5c415 00000000",
			writeBoxedVector64([]int64{}),
		},
		{
			"boxedVector64_1",
			"83659ba8 1cb5c415 00000003 22076cba 00000001 00000000 22076cba 00000002 00000000 22076cba 00000003 00000000",
			writeBoxedVector64([]int64{1, 2, 3}),
		},

		{
			"boxedTuple0",
			"30c9d533 9770768a 00000000 00000000 00000000",
			writeBoxedTuple([3]int32{0, 0, 0}),
		},
		{
			"boxedTuple1",
			"30c9d533 9770768a 00000001 00000002 00000003",
			writeBoxedTuple([3]int32{1, 2, 3}),
		},

		{
			"boxedString0",
			"548994db b5286e24 00000000",
			writeBoxedString(""),
		},
		{
			"boxedString1",
			"548994db b5286e24 63626104 00000064",
			writeBoxedString("abcd"),
		},

		{
			"boxedInt0",
			"5688ebaf a8509bda 00000000",
			writeBoxedInt(0),
		},
		{
			"boxedInt1",
			"5688ebaf a8509bda 00114b6f",
			writeBoxedInt(1133423),
		},

		{"getMyValueInt0", "b3df27fe c12375b7 a8509bda 00000000", writeGetMyValue(tl.MyInt{Val1: 0}.AsUnion())},
		{"getMyValueInt1", "b3df27fe c12375b7 a8509bda 00000001", writeGetMyValue(tl.MyInt{Val1: 1}.AsUnion())},
		{
			"getMyValueString0",
			"b3df27fe c8bfa969 b5286e24 00000000",
			writeGetMyValue(tl.MyString{Val2: ""}.AsUnion()),
		},
		{
			"getMyValueString1",
			"b3df27fe c8bfa969 b5286e24 6473610b 61646373 33646473",
			writeGetMyValue(tl.MyString{Val2: "asdscdasdd3"}.AsUnion()),
		},

		// Для словарей можем иметь тесты только для 0 и 1 элементов,
		// потому что для остальных длин будет расходиться сериализация
		// из-за рандомного порядка обхода мапов.
		{
			"getMyDictOfInt0",
			"166f962c b8019a3d 00000000",
			writeGetMyDictOfInt(tl.MyDictOfInt{}),
		},
		{
			"getMyDictOfInt1",
			"166f962c b8019a3d 00000001 00006101 00000000",
			writeGetMyDictOfInt(tl.MyDictOfInt{"a": 0}),
		},

		{
			"getDouble0",
			"39711d7b 2210c154 00000000 00000000",
			writeGetDouble(0),
		},
		{
			"getDouble1",
			"39711d7b 2210c154 8f5c28f6 40a94542",
			writeGetDouble(3234.63),
		},
		{
			"getDouble2",
			"39711d7b 2210c154 29f16b12 c05e0000",
			writeGetDouble(-120.00001),
		},

		{"getNonOptNat0/nil", "67665961 00000000 9770768a", writeGetNonOptNat(0, nil)},
		{"getNonOptNat0/[]", "67665961 00000000 9770768a", writeGetNonOptNat(0, []int32{})},
		{
			"getNonOptNat2",
			"67665961 00000002 9770768a 00000001 00000002",
			writeGetNonOptNat(2, []int32{1, 2}),
		},

		{"getFloat/0", "25a7bc68 00000000", writeGetFloat(0)},
		{"getFloat/small", "25a7bc68 c2280000", writeGetFloat(-42.0)},
		{"getFloat/big", "25a7bc68 51ae12fc", writeGetFloat(93455349934.33111214930)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w, err := test.write(nil)
			if err != nil {
				t.Fatalf("%s: error: %v\n", test.name, err)
			}
			have := sprintHexDump(w)
			require.Equal(t, test.want, have)
		})
	}
}

func TestReadResult(t *testing.T) {
	largeService2 := tlservice2.AddOrIncrMany{
		ObjectIdLength:   4,
		IntCountersNum:   3,
		FloatCountersNum: 2,
		ObjectsNum:       2,
		IntCounters:      []int32{100, 200, 300},
		FloatCounters:    []int32{5, 8},
		Deltas: []tlservice2.DeltaSet{
			{
				Id: tlservice2.ObjectId{Id: []int32{1, 2, 3, 4}},
				Counters: tlservice2.CounterSet{
					IntCounters:   []int32{500, 600, 700},
					FloatCounters: []float64{-0.5, 11.5},
				},
			}, {
				Id: tlservice2.ObjectId{Id: []int32{5, 6, 7, 8}},
				Counters: tlservice2.CounterSet{
					IntCounters:   []int32{5000, 6000, 7000},
					FloatCounters: []float64{-0.05, 110.5},
				},
			},
		},
	}

	smallService2 := tlservice2.AddOrIncrMany{
		ObjectIdLength:   0,
		IntCountersNum:   0,
		FloatCountersNum: 1,
		ObjectsNum:       1,
		IntCounters:      nil,
		FloatCounters:    []int32{8},
		Deltas: []tlservice2.DeltaSet{
			{
				Id: tlservice2.ObjectId{Id: nil},
				Counters: tlservice2.CounterSet{
					IntCounters:   nil,
					FloatCounters: []float64{-0.05},
				},
			},
		},
	}

	t.Run("GetMaybeIface", func(t *testing.T) {
		testGetMaybeIface := func(x tlservice1.ValueBoxedMaybe, goldStandard []byte, jsonStandard string, jsonStandardOld string) {
			var y, z tlservice1.ValueBoxedMaybe
			var fn tl.GetMaybeIface
			checkTestCase2(t, &x, &y, goldStandard, x.WriteBoxedGeneral, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, x.WriteJSONGeneral, y.WriteJSONGeneral, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
			checkJSONCase3(t, &x, &z, jsonStandardOld, func(b bool, b2 bool, i []byte) ([]byte, error) {
				return x.WriteJSONOpt(b, b2, i), nil
			}, func(b bool, b2 bool, i []byte) ([]byte, error) {
				return y.WriteJSONOpt(b, b2, i), nil
			}, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
		}

		testGetMaybeIface(tlservice1.ValueBoxedMaybe{
			Value: tlservice1.Value{},
			Ok:    false,
		}, []byte{0x7b, 0xa, 0x93, 0x27},
			`{}`,
			`{}`)
		testGetMaybeIface(tlservice1.ValueBoxedMaybe{
			Value: tlservice1.Longvalue{Value: 102}.AsUnion(),
			Ok:    true,
		}, []byte{0xf8, 0x8e, 0x9c, 0x3f, 0x45, 0x9, 0x2e, 0x8, 0x66, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			`{"ok":true,"value":{"type":"service1.longvalue","value":{"value":102}}}`,
			`{"ok":true,"value":{"type":"service1.longvalue#082e0945","value":{"value":102}}}`)
		testGetMaybeIface(tlservice1.ValueBoxedMaybe{
			Value: tlservice1.Strvalue{Value: "hello, world!"}.AsUnion(),
			Ok:    true,
		}, []byte{0xf8, 0x8e, 0x9c, 0x3f, 0x52, 0xc, 0xaa, 0x5f, 0xd, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x21, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			`{"ok":true,"value":{"type":"service1.strvalue","value":{"value":"hello, world!"}}}`,
			`{"ok":true,"value":{"type":"service1.strvalue#5faa0c52","value":{"value":"hello, world!"}}}`)
	})

	t.Run("GetStats", func(t *testing.T) {
		testGetStats := func(x tltasks.QueueTypeStats, goldStandard []byte, jsonStandard string) {
			var y, z tltasks.QueueTypeStats
			var fn tl.GetStats
			checkTestCase2(t, &x, &y, goldStandard, x.WriteBoxedGeneral, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, x.WriteJSONGeneral, y.WriteJSONGeneral, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
		}

		testGetStats(tltasks.QueueTypeStats{}, []byte{0xf2, 0x85, 0xb7, 0xe1, 0x0, 0x0, 0x0, 0x0}, `{}`)
		testGetStats(tltasks.QueueTypeStats{
			FieldsMask:  1 << 0,
			WaitingSize: 0xFFFFFFFFF1,
		}, []byte{0xf2, 0x85, 0xb7, 0xe1, 0x1, 0x0, 0x0, 0x0, 0xf1, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0}, `{"fields_mask":1,"waiting_size":1099511627761}`)
		testGetStats(tltasks.QueueTypeStats{
			FieldsMask:    1 << 1,
			ScheduledSize: 0xFFFFFFFFF2,
		}, []byte{0xf2, 0x85, 0xb7, 0xe1, 0x2, 0x0, 0x0, 0x0, 0xf2, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0}, `{"fields_mask":2,"scheduled_size":1099511627762}`)
		testGetStats(tltasks.QueueTypeStats{
			FieldsMask:     1 << 2,
			InProgressSize: 0xFFFFFFFFF3,
		}, []byte{0xf2, 0x85, 0xb7, 0xe1, 0x4, 0x0, 0x0, 0x0, 0xf3, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0}, `{"fields_mask":4,"in_progress_size":1099511627763}`)
		testGetStats(tltasks.QueueTypeStats{
			FieldsMask: 1 << 3,
			NumQueues:  0xFFFFF4,
		}, []byte{0xf2, 0x85, 0xb7, 0xe1, 0x8, 0x0, 0x0, 0x0, 0xf4, 0xff, 0xff, 0x0}, `{"fields_mask":8,"num_queues":16777204}`)
		testGetStats(tltasks.QueueTypeStats{
			FieldsMask:     (1 << 0) | (1 << 1) | (1 << 2) | (1 << 3),
			WaitingSize:    0xFFFFFFFFF1,
			ScheduledSize:  0xFFFFFFFFF2,
			InProgressSize: 0xFFFFFFFFF3,
			NumQueues:      0xFFFFF4,
		}, []byte{0xf2, 0x85, 0xb7, 0xe1, 0xf, 0x0, 0x0, 0x0, 0xf1, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0xf2, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0xf3, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0xf4, 0xff, 0xff, 0x0}, `{"fields_mask":15,"waiting_size":1099511627761,"scheduled_size":1099511627762,"in_progress_size":1099511627763,"num_queues":16777204}`)
	})

	t.Run("TasksGetQueueSize", func(t *testing.T) {
		testTasksGetQueueSize := func(x tltasks.QueueStats, field_mask uint32, goldStandard []byte, jsonStandard string) {
			var y, z tltasks.QueueStats
			var fn tltasks.GetQueueSize
			fn.FieldsMask = field_mask
			checkTestCase2(t, &x, &y, goldStandard, func(w []byte) ([]byte, error) { return x.WriteBoxed(w, field_mask), nil }, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, func(w []byte) ([]byte, error) { return x.WriteJSON(w, field_mask), nil }, func(w []byte) ([]byte, error) { return y.WriteJSON(w, field_mask), nil }, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
		}

		testTasksGetQueueSize(tltasks.QueueStats{}, 0, []byte{0x43, 0x25, 0x94, 0x1d}, `{}`)
		testTasksGetQueueSize(tltasks.QueueStats{
			WaitingSize: 1,
		}, 1<<0, []byte{0x43, 0x25, 0x94, 0x1d, 0x1, 0x0, 0x0, 0x0}, `{"waiting_size":1}`)
		testTasksGetQueueSize(tltasks.QueueStats{
			ScheduledSize: 2,
		}, 1<<1, []byte{0x43, 0x25, 0x94, 0x1d, 0x2, 0x0, 0x0, 0x0}, `{"scheduled_size":2}`)
		testTasksGetQueueSize(tltasks.QueueStats{
			InProgressSize: 3,
		}, 1<<2, []byte{0x43, 0x25, 0x94, 0x1d, 0x3, 0x0, 0x0, 0x0}, `{"in_progress_size":3}`)
		testTasksGetQueueSize(tltasks.QueueStats{
			WaitingSize:    1,
			InProgressSize: 3,
		}, (1<<0)|(1<<2), []byte{0x43, 0x25, 0x94, 0x1d, 0x1, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0}, `{"waiting_size":1,"in_progress_size":3}`)
		testTasksGetQueueSize(tltasks.QueueStats{}, (1<<0)|(1<<1)|(1<<2), []byte{0x43, 0x25, 0x94, 0x1d, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, `{"waiting_size":0,"scheduled_size":0,"in_progress_size":0}`)
		testTasksGetQueueSize(tltasks.QueueStats{
			WaitingSize:    1,
			ScheduledSize:  2,
			InProgressSize: 3,
		}, (1<<0)|(1<<1)|(1<<2), []byte{0x43, 0x25, 0x94, 0x1d, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0}, `{"waiting_size":1,"scheduled_size":2,"in_progress_size":3}`)
	})

	t.Run("TasksGetAnyTask", func(t *testing.T) {
		testTasksGetAnyTask := func(x tltasks.TaskInfoMaybe, goldStandard []byte, jsonStandard string) {
			var y, z tltasks.TaskInfoMaybe
			var fn tltasks.GetAnyTask
			checkTestCase2(t, &x, &y, goldStandard, x.WriteBoxedGeneral, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, func(w []byte) ([]byte, error) { return fn.WriteResultJSON(w, x) }, func(w []byte) ([]byte, error) { return fn.WriteResultJSON(w, y) }, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
		}

		testTasksGetAnyTask(tltasks.TaskInfoMaybe{Ok: false}, []byte{0x7b, 0xa, 0x93, 0x27}, `{}`)
		testTasksGetAnyTask(tltasks.TaskInfoMaybe{Ok: true}, []byte{0xf8, 0x8e, 0x9c, 0x3f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, `{"ok":true,"value":{"task":{}}}`)
		testTasksGetAnyTask(tltasks.TaskInfoMaybe{
			Ok:    true,
			Value: tltasks.TaskInfo{QueueId: nil},
		}, []byte{0xf8, 0x8e, 0x9c, 0x3f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, `{"ok":true,"value":{"task":{}}}`)
		testTasksGetAnyTask(tltasks.TaskInfoMaybe{
			Ok:    true,
			Value: tltasks.TaskInfo{QueueId: []int32{}},
		}, []byte{0xf8, 0x8e, 0x9c, 0x3f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, `{"ok":true,"value":{"task":{}}}`)
		testTasksGetAnyTask(tltasks.TaskInfoMaybe{
			Ok: true,
			Value: tltasks.TaskInfo{
				TypeName: "some type name",
				Task: tltasks.Task{
					Flags: 1,
					Tag:   []int32{0, 0, 0},
					Data:  "the data string",
				},
			},
		}, []byte{0xf8, 0x8e, 0x9c, 0x3f, 0xe, 0x73, 0x6f, 0x6d, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0x74, 0x68, 0x65, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67}, `{"ok":true,"value":{"type_name":"some type name","task":{"flags":1,"tag":[0,0,0],"data":"the data string"}}}`)
	})

	t.Run("GetMyValueInt", func(t *testing.T) {
		testGetMyValueInt := func(x tl.MyValue, goldStandard []byte, jsonStandard string, jsonStandardOld string) {
			y := tl.MyInt{}.AsUnion()
			z := tl.MyInt{}.AsUnion()
			var fn tl.GetMyValue
			checkTestCase2(t, &x, &y, goldStandard, x.WriteBoxedGeneral, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, x.WriteJSONGeneral, y.WriteBoxedGeneral, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
			checkJSONCase3(t, &x, &z, jsonStandardOld, func(b bool, b2 bool, i []byte) ([]byte, error) { return x.WriteJSONOpt(b, b2, i), nil }, func(b bool, b2 bool, i []byte) ([]byte, error) { return y.WriteJSONOpt(b, b2, i), nil }, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
		}

		testGetMyValueInt(tl.MyInt{Val1: 0}.AsUnion(), []byte{0xb7, 0x75, 0x23, 0xc1, 0xda, 0x9b, 0x50, 0xa8, 0x0, 0x0, 0x0, 0x0},
			`{"type":"myInt","value":{}}`,
			`{"type":"myInt#c12375b7","value":{}}`)
		testGetMyValueInt(tl.MyInt{Val1: -1}.AsUnion(), []byte{0xb7, 0x75, 0x23, 0xc1, 0xda, 0x9b, 0x50, 0xa8, 0xff, 0xff, 0xff, 0xff},
			`{"type":"myInt","value":{"val1":-1}}`,
			`{"type":"myInt#c12375b7","value":{"val1":-1}}`)
		testGetMyValueInt(tl.MyInt{Val1: 123129491}.AsUnion(), []byte{0xb7, 0x75, 0x23, 0xc1, 0xda, 0x9b, 0x50, 0xa8, 0x93, 0xce, 0x56, 0x7},
			`{"type":"myInt","value":{"val1":123129491}}`,
			`{"type":"myInt#c12375b7","value":{"val1":123129491}}`)
	})

	t.Run("GetMyValueString", func(t *testing.T) {
		testGetMyValueString := func(x tl.MyValue, goldStandard []byte, jsonStandard string, jsonStandardOld string) {
			y := tl.MyString{}.AsUnion()
			z := tl.MyString{}.AsUnion()
			var fn tl.GetMyValue
			checkTestCase2(t, &x, &y, goldStandard, x.WriteBoxedGeneral, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, x.WriteJSONGeneral, y.WriteJSONGeneral, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
			checkJSONCase3(t, &x, &z, jsonStandardOld, func(b bool, b2 bool, i []byte) ([]byte, error) { return x.WriteJSONOpt(b, b2, i), nil }, func(b bool, b2 bool, i []byte) ([]byte, error) { return y.WriteJSONOpt(b, b2, i), nil }, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
		}

		testGetMyValueString(tl.MyString{Val2: ""}.AsUnion(), []byte{0x69, 0xa9, 0xbf, 0xc8, 0x24, 0x6e, 0x28, 0xb5, 0x0, 0x0, 0x0, 0x0},
			`{"type":"myString","value":{}}`,
			`{"type":"myString#c8bfa969","value":{}}`)
		testGetMyValueString(tl.MyString{Val2: "123"}.AsUnion(), []byte{0x69, 0xa9, 0xbf, 0xc8, 0x24, 0x6e, 0x28, 0xb5, 0x3, 0x31, 0x32, 0x33},
			`{"type":"myString","value":{"val2":"123"}}`,
			`{"type":"myString#c8bfa969","value":{"val2":"123"}}`)
		testGetMyValueString(tl.MyString{Val2: strings.Repeat("x", 300)}.AsUnion(), []byte{0x69, 0xa9, 0xbf, 0xc8, 0x24, 0x6e, 0x28, 0xb5, 0xfe, 0x2c, 0x1, 0x0, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78},
			`{"type":"myString","value":{"val2":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}}`,
			`{"type":"myString#c8bfa969","value":{"val2":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}}`)
	})

	t.Run("GetMyDictOfInt", func(t *testing.T) {
		testGetMyDictOfInt := func(x tl.MyDictOfInt, goldStandard []byte, jsonStandard string) {
			var y, z tl.MyDictOfInt
			var fn tl.GetMyDictOfInt
			checkTestCase2(t, &x, &y, goldStandard, x.WriteBoxedGeneral, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, x.WriteJSONGeneral, y.WriteJSONGeneral, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
		}

		testGetMyDictOfInt(nil, []byte{0x3d, 0x9a, 0x1, 0xb8, 0x0, 0x0, 0x0, 0x0}, ``)
		testGetMyDictOfInt(tl.MyDictOfInt{}, []byte{0x3d, 0x9a, 0x1, 0xb8, 0x0, 0x0, 0x0, 0x0}, `{}`)
		testGetMyDictOfInt(tl.MyDictOfInt{"x": 1}, []byte{0x3d, 0x9a, 0x1, 0xb8, 0x1, 0x0, 0x0, 0x0, 0x1, 0x78, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0}, `{"x":1}`)
		testGetMyDictOfInt(tl.MyDictOfInt{
			"1": 1,
			"2": 2,
			"3": 3,
		}, nil /* Write of map with size > 1 is non-deterministic */, `{"1":1,"2":2,"3":3}`)
	})

	t.Run("BoxedTupleSlice2", func(t *testing.T) {
		testBoxedTupleSlice2 := func(x tl.MyBoxedTupleSlice, goldStandard []byte, jsonStandard string) {
			var y, z tl.MyBoxedTupleSlice
			var fn tl.BoxedTupleSlice2
			checkTestCase2(t, &x, &y, goldStandard, x.WriteBoxed, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, x.WriteJSON, y.WriteJSON, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
		}

		testBoxedTupleSlice2(tl.MyBoxedTupleSlice{N: 0, Data: nil}, []byte{0xbe, 0xa1, 0xd1, 0x25, 0x0, 0x0, 0x0, 0x0, 0x8a, 0x76, 0x70, 0x97}, `{}`)
		testBoxedTupleSlice2(tl.MyBoxedTupleSlice{N: 0, Data: []int32{}}, []byte{0xbe, 0xa1, 0xd1, 0x25, 0x0, 0x0, 0x0, 0x0, 0x8a, 0x76, 0x70, 0x97}, `{}`)
		testBoxedTupleSlice2(tl.MyBoxedTupleSlice{N: 3, Data: []int32{1, 2, 3}}, []byte{0xbe, 0xa1, 0xd1, 0x25, 0x3, 0x0, 0x0, 0x0, 0x8a, 0x76, 0x70, 0x97, 0xda, 0x9b, 0x50, 0xa8, 0x1, 0x0, 0x0, 0x0, 0xda, 0x9b, 0x50, 0xa8, 0x2, 0x0, 0x0, 0x0, 0xda, 0x9b, 0x50, 0xa8, 0x3, 0x0, 0x0, 0x0}, `{"n":3,"data":[1,2,3]}`)
		testBoxedTupleSlice2(tl.MyBoxedTupleSlice{N: 3, Data: []int32{3, 2, 1}}, []byte{0xbe, 0xa1, 0xd1, 0x25, 0x3, 0x0, 0x0, 0x0, 0x8a, 0x76, 0x70, 0x97, 0xda, 0x9b, 0x50, 0xa8, 0x3, 0x0, 0x0, 0x0, 0xda, 0x9b, 0x50, 0xa8, 0x2, 0x0, 0x0, 0x0, 0xda, 0x9b, 0x50, 0xa8, 0x1, 0x0, 0x0, 0x0}, `{"n":3,"data":[3,2,1]}`)
	})

	t.Run("BoxedTupleSlice3", func(t *testing.T) {
		testBoxedTupleSlice3 := func(n uint32, x []int32, goldStandard []byte, jsonStandard string) {
			y := make([]int32, n)
			z := make([]int32, n)
			var fn tl.BoxedTupleSlice3
			fn.N = n
			checkTestCase2(t, &x, &y, goldStandard, func(w []byte) ([]byte, error) {
				w = basictl.NatWrite(w, constants.Tuple)
				for i := range x[:n] {
					w = basictl.IntWrite(w, x[i])
				}
				return w, nil
			}, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, func(w []byte) ([]byte, error) { return fn.WriteResultJSON(w, x) }, func(w []byte) ([]byte, error) { return fn.WriteResultJSON(w, y) }, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
		}

		testBoxedTupleSlice3(0, nil, []byte{0x8a, 0x76, 0x70, 0x97}, ``)
		testBoxedTupleSlice3(0, []int32{}, []byte{0x8a, 0x76, 0x70, 0x97}, `[]`)
		testBoxedTupleSlice3(1, []int32{1}, []byte{0x8a, 0x76, 0x70, 0x97, 0x1, 0x0, 0x0, 0x0}, `[1]`)
		testBoxedTupleSlice3(2, []int32{1, 2}, []byte{0x8a, 0x76, 0x70, 0x97, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0}, `[1,2]`)
	})

	t.Run("BoxedArray", func(t *testing.T) {
		testBoxedArray := func(x tl.MyBoxedArray, goldStandard []byte, jsonStandard string) {
			var y, z tl.MyBoxedArray
			var fn tl.BoxedArray
			checkTestCase2(t, &x, &y, goldStandard, x.WriteBoxedGeneral, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, x.WriteJSONGeneral, y.WriteJSONGeneral, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
		}

		testBoxedArray(tl.MyBoxedArray{Data: [2]int32{0, 0}}, []byte{0xf0, 0x64, 0x8f, 0x28, 0x8a, 0x76, 0x70, 0x97, 0xda, 0x9b, 0x50, 0xa8, 0x0, 0x0, 0x0, 0x0, 0xda, 0x9b, 0x50, 0xa8, 0x0, 0x0, 0x0, 0x0}, `{"data":[0,0]}`)
		testBoxedArray(tl.MyBoxedArray{Data: [2]int32{1, 2}}, []byte{0xf0, 0x64, 0x8f, 0x28, 0x8a, 0x76, 0x70, 0x97, 0xda, 0x9b, 0x50, 0xa8, 0x1, 0x0, 0x0, 0x0, 0xda, 0x9b, 0x50, 0xa8, 0x2, 0x0, 0x0, 0x0}, `{"data":[1,2]}`)
		testBoxedArray(tl.MyBoxedArray{Data: [2]int32{2, 1}}, []byte{0xf0, 0x64, 0x8f, 0x28, 0x8a, 0x76, 0x70, 0x97, 0xda, 0x9b, 0x50, 0xa8, 0x2, 0x0, 0x0, 0x0, 0xda, 0x9b, 0x50, 0xa8, 0x1, 0x0, 0x0, 0x0}, `{"data":[2,1]}`)
	})

	t.Run("GetMyDouble", func(t *testing.T) {
		testGetMyDouble := func(x tl.MyDouble, goldStandard []byte, jsonStandard string) {
			var y, z tl.MyDouble
			var fn tl.GetMyDouble
			checkTestCase2(t, &x, &y, goldStandard, x.WriteBoxedGeneral, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, x.WriteJSONGeneral, y.WriteJSONGeneral, func(in *basictl.JsonLexer) error {
				var res tl.MyDouble
				err := fn.ReadResultJSON(true, in, &res)
				z = res
				return err
			})
		}

		testGetMyDouble(0, []byte{0x26, 0xc7, 0xa6, 0x90, 0x54, 0xc1, 0x10, 0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, `0`)
		testGetMyDouble(1.0, []byte{0x26, 0xc7, 0xa6, 0x90, 0x54, 0xc1, 0x10, 0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf0, 0x3f}, `1`)
		testGetMyDouble(0.1, []byte{0x26, 0xc7, 0xa6, 0x90, 0x54, 0xc1, 0x10, 0x22, 0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xb9, 0x3f}, `0.1`)
		testGetMyDouble(1033.12, []byte{0x26, 0xc7, 0xa6, 0x90, 0x54, 0xc1, 0x10, 0x22, 0x14, 0xae, 0x47, 0xe1, 0x7a, 0x24, 0x90, 0x40}, `1033.12`)
		testGetMyDouble(-1033.12, []byte{0x26, 0xc7, 0xa6, 0x90, 0x54, 0xc1, 0x10, 0x22, 0x14, 0xae, 0x47, 0xe1, 0x7a, 0x24, 0x90, 0xc0}, `-1033.12`)
	})

	t.Run("GetFloat", func(t *testing.T) {
		testGetFloat := func(x float32, goldStandard []byte, jsonStandard string) {
			var y, z float32
			var fn tl.GetFloat
			checkTestCase2(t, &x, &y, goldStandard, func(w []byte) ([]byte, error) { return (*tl.Float)(&x).WriteBoxed(w), nil }, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, func(w []byte) ([]byte, error) { return fn.WriteResultJSON(w, x) }, func(w []byte) ([]byte, error) { return fn.WriteResultJSON(w, y) }, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
		}

		testGetFloat(0, []byte{0x22, 0xab, 0x4d, 0x82, 0x0, 0x0, 0x0, 0x0}, `0`)
		testGetFloat(393.6, []byte{0x22, 0xab, 0x4d, 0x82, 0xcd, 0xcc, 0xc4, 0x43}, `393.6`)
		testGetFloat(-1248128.1827, []byte{0x22, 0xab, 0x4d, 0x82, 0x1, 0x5c, 0x98, 0xc9}, `-1248128.1`)
	})

	t.Run("Enum", func(t *testing.T) {
		t.Skip("unable to get Enum object from JSON in for of byte slice without function object")
		// TODO: change z.ReadJSON to another function to run this test correctly
		// testEnum := func(x tl.TasksTaskStatus, goldStandard []byte, jsonStandard string) {
		//	var y, z tl.TasksTaskStatus
		//	checkTestCase2(t, &x, &y, goldStandard, x.WriteBoxed, y.ReadBoxed)
		//	checkJSONCase2(t, &x, &z, jsonStandard, x.WriteJSON, y.WriteJSON, z.ReadJSON)
		// }

		// testEnum(tl.TasksTaskStatusScheduled(), []byte{0xa9, 0x80, 0xca, 0xa}, `"tasks.taskStatusScheduled#0aca80a9"`)
		// testEnum(tl.TasksTaskStatusInProgress(), []byte{0xe7, 0x70, 0xef, 0x6}, `"tasks.taskStatusInProgress#06ef70e7"`)
	})

	t.Run("Service2", func(t *testing.T) {
		t.Log("unfair test")
		// TODO: not fair transformation, incorrect json might be fixed
		testService2 := func(x tlservice2.AddOrIncrMany, goldStandard []byte, jsonStandard string) {
			var y, z tlservice2.AddOrIncrMany
			checkTestCase2(t, &x, &y, goldStandard, x.WriteBoxed, func(w []byte) ([]byte, error) {
				w, err := basictl.NatReadExactTag(w, y.TLTag())
				if err != nil {
					return w, err
				}
				return y.Read(w)
			})
			checkJSONCase2(t, &x, &z, jsonStandard, x.WriteJSON, y.WriteJSON, func(in *basictl.JsonLexer) error {
				return z.ReadJSON(false, in)
			})
		}

		testService2(largeService2, []byte{0x89, 0x24, 0xa5, 0x5a, 0x4, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x64, 0x0, 0x0, 0x0, 0xc8, 0x0, 0x0, 0x0, 0x2c, 0x1, 0x0, 0x0, 0x5, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0xf4, 0x1, 0x0, 0x0, 0x58, 0x2, 0x0, 0x0, 0xbc, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe0, 0xbf, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x27, 0x40, 0x5, 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x7, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x88, 0x13, 0x0, 0x0, 0x70, 0x17, 0x0, 0x0, 0x58, 0x1b, 0x0, 0x0, 0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xa9, 0xbf, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa0, 0x5b, 0x40}, `{"objectIdLength":4,"intCountersNum":3,"floatCountersNum":2,"objectsNum":2,"intCounters":[100,200,300],"floatCounters":[5,8],"deltas":[{"id":{"id":[1,2,3,4]},"counters":{"intCounters":[500,600,700],"floatCounters":[-0.5,11.5]}},{"id":{"id":[5,6,7,8]},"counters":{"intCounters":[5000,6000,7000],"floatCounters":[-0.05,110.5]}}]}`)
		testService2(smallService2, []byte{0x89, 0x24, 0xa5, 0x5a, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xa9, 0xbf}, `{"floatCountersNum":1,"objectsNum":1,"floatCounters":[8],"deltas":[{"id":{},"counters":{"floatCounters":[-0.05]}}]}`)
	})

	t.Run("Service2Result", func(t *testing.T) {
		testService2Result := func(fn tlservice2.AddOrIncrMany, x []tlservice2.CounterSet, goldStandard []byte, jsonStandard string) {
			var y []tlservice2.CounterSet
			var z []tlservice2.CounterSet
			checkTestCase2(t, &x, &y, goldStandard, func(w []byte) ([]byte, error) { return fn.WriteResult(w, x) }, func(r []byte) ([]byte, error) { return fn.ReadResult(r, &y) })
			checkJSONCase2(t, &x, &z, jsonStandard, func(w []byte) ([]byte, error) { return fn.WriteResultJSON(w, x) }, func(w []byte) ([]byte, error) { return fn.WriteResultJSON(w, y) }, func(in *basictl.JsonLexer) error { return fn.ReadResultJSON(true, in, &z) })
		}

		testService2Result(largeService2, []tlservice2.CounterSet{{
			IntCounters:   []int32{1, 2, 3},
			FloatCounters: []float64{5.5, -111.25},
		}, {
			IntCounters:   []int32{4, 5, 6},
			FloatCounters: []float64{-5.5, 1.25},
		}}, []byte{0x8a, 0x76, 0x70, 0x97, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x16, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0xd0, 0x5b, 0xc0, 0x4, 0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x16, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf4, 0x3f}, `[{"intCounters":[1,2,3],"floatCounters":[5.5,-111.25]},{"intCounters":[4,5,6],"floatCounters":[-5.5,1.25]}]`)
		testService2Result(smallService2, []tlservice2.CounterSet{{
			IntCounters:   nil,
			FloatCounters: []float64{0.1},
		}}, []byte{0x8a, 0x76, 0x70, 0x97, 0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xb9, 0x3f}, `[{"floatCounters":[0.1]}]`)
	})
}

func sprintHexDump(data []byte) string {
	var buf bytes.Buffer
	buf.Grow(len(data) + len(data)/4)
	for i := 0; i < len(data); i += 4 {
		// Печатаем октеты в обратном порядке, чтобы они совпадали
		// с константами из `constant.go`.
		_, _ = fmt.Fprintf(&buf, "%02x%02x%02x%02x ",
			data[i+3],
			data[i+2],
			data[i+1],
			data[i+0])
	}
	return strings.TrimSpace(buf.String())
}
