package codegen_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/vkcom/tl/internal/utils"

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
	FunctionName      string
	FunctionBodyBytes string
	FunctionBodyJson  string
	ResultBytes       string
	ResultJson        string
}

type writeTest struct {
	TestName          string
	TLType            string
	TestingObjectData map[string]any
	ExpectedOutput    string
}

func TestWriteArgs(t *testing.T) {
	const PathToJsonData = "data/test-functions-bytes.json"
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

			bytes_ := utils.ParseHexToBytes(test.ExpectedOutput)
			_, _ = obj.ReadBoxed(bytes_)

			objBytes, err := obj.WriteBoxedGeneral(nil)
			if err != nil {
				t.Fatalf("%s: error: %v\n", test.TestName, err)
			}

			require.Equal(t, test.ExpectedOutput, utils.SprintHexDump(objBytes))
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

	_, jsonResult, jsonErr := fn.ReadResultWriteResultJSON(utils.ParseHexToBytes(resultBytesValue), nil)
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
	if !assert.Equal(t, resultBytesValue, utils.SprintHexDump(bytesResult)) {
		t.Fatalf("Write bytes failed, difference:v\n%s\n", cmp.Diff(resultBytesValue, utils.SprintHexDump(bytesResult)))
	}
}

func TestReadResult(t *testing.T) {
	const PathToJsonData = "data/test-functions-bytes.json"
	data, readErr := os.ReadFile(PathToJsonData)

	if readErr != nil {
		t.Fatalf("testing data is not provided")
		return
	}

	newTests := allTests{}
	_ = json.Unmarshal(data, &newTests)

	for _, test := range newTests.TestReadFunction {
		t.Run(test.FunctionName, func(t *testing.T) {
			checkFunctionReadWrite(t, test.FunctionName, test.FunctionBodyJson, test.ResultBytes, test.ResultJson)
		})
	}
}
