package codegen_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/factory"
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

type allTests struct {
	TestWriteBoxed   []writeTest
	TestReadFunction []readFunctionTest
}

type readFunctionTest struct {
	FunctionName string
	FunctionBody string
	ResultBytes  string
	ResultJson   string
}

type writeTest struct {
	TestName          string
	TLType            string
	TestingObjectData map[string]any
	ExpectedOutput    string
}

func checkFunctionReadWrite(t *testing.T, fnType string, fnJsonValue string, resultBytesValue string, resultJsonValue string) {
	fn := factory.CreateFunctionFromName(fnType)
	initFnErr := fn.ReadJSON(false, &basictl.JsonLexer{Data: []byte(fnJsonValue)})
	if initFnErr != nil {
		t.Fatalf("Function initilization error: %v\n", initFnErr)
		return
	}

	_, jsonResult, jsonErr := fn.ReadResultWriteResultJSON(parseHexToBytes(resultBytesValue), nil)
	if jsonErr != nil {
		t.Fatalf("Write json error: %v\n", jsonErr)
		return
	}
	if !assert.Equal(t, resultJsonValue, string(jsonResult)) {
		t.Fatalf("Write json failed, difference:v\n%s\n", cmp.Diff(resultJsonValue, jsonResult))
	}
	_, bytesResult, bytesErr := fn.ReadResultJSONWriteResult([]byte(resultJsonValue), nil)
	if bytesErr != nil {
		t.Fatalf("Write bytes error: %v\n", jsonErr)
		return
	}
	if !assert.Equal(t, resultBytesValue, sprintHexDump(bytesResult)) {
		t.Fatalf("Write bytes failed, difference:v\n%s\n", cmp.Diff(resultBytesValue, sprintHexDump(bytesResult)))
	}
}

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

func TestWriteArgs(t *testing.T) {
	const PathToJsonData = "casetests/data/test-bytes.json"
	data, readErr := os.ReadFile(PathToJsonData)

	if readErr != nil {
		t.Fatalf("testing data is not provided")
		return
	}

	newTests := allTests{}
	_ = json.Unmarshal(data, &newTests)

	for _, test := range newTests.TestWriteBoxed {
		t.Run(test.TestName, func(t *testing.T) {
			obj := factory.CreateObjectFromName(test.TLType)

			bytes_ := parseHexToBytes(test.ExpectedOutput)
			_, _ = obj.ReadBoxed(bytes_)

			objBytes, err := obj.WriteBoxedGeneral(nil)
			if err != nil {
				t.Fatalf("%s: error: %v\n", test.TestName, err)
			}

			require.Equal(t, test.ExpectedOutput, sprintHexDump(objBytes))
		})
	}
}

func TestReadResult(t *testing.T) {
	const PathToJsonData = "casetests/data/test-bytes.json"
	data, readErr := os.ReadFile(PathToJsonData)

	if readErr != nil {
		t.Fatalf("testing data is not provided")
		return
	}

	newTests := allTests{}
	_ = json.Unmarshal(data, &newTests)

	for _, test := range newTests.TestReadFunction {
		t.Run(test.FunctionName, func(t *testing.T) {
			checkFunctionReadWrite(t, test.FunctionName, test.FunctionBody, test.ResultBytes, test.ResultJson)
		})
	}

	t.Run("Enum", func(t *testing.T) {
		//t.Skip("unable to get Enum object from JSON in for of byte slice without function object")
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

func parseHexToBytes(data string) []byte {
	var result []byte
	for _, octet := range strings.Split(data, " ") {
		b1, _ := hex.DecodeString(octet[6:8])
		b2, _ := hex.DecodeString(octet[4:6])
		b3, _ := hex.DecodeString(octet[2:4])
		b4, _ := hex.DecodeString(octet[0:2])
		result = append(result, b1...)
		result = append(result, b2...)
		result = append(result, b3...)
		result = append(result, b4...)
	}
	return result
}
