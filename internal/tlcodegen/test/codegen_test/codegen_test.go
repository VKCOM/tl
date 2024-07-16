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
	"github.com/stretchr/testify/require"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/factory"
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
