package linter

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/tlcodegen"
)

const TestSamplesDir = "../../tls/backward_compatibility_samples"
const CorrectSample = "schema-correct.tl"

func RunCompatibilityTest(t *testing.T, prevStateFile, newStateFile string) *tlast.ParseError {
	prevPath := filepath.Join(TestSamplesDir, prevStateFile)
	newPath := filepath.Join(TestSamplesDir, newStateFile)

	prevTL, err := readTL(prevPath)
	if err != nil {
		t.Fatal(err)
	}
	newTL, err := readTL(newPath)
	if err != nil {
		t.Fatal(err)
	}
	return tlcodegen.CheckBackwardCompatibility(&newTL, &prevTL)
}

func readTL(file string) (tlast.TL, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return tlast.ParseTLFile(string(data), file, tlast.LexerOptions{
		AllowBuiltin: false,
		AllowDirty:   false,
	})
}

func TestCompatibilityLinter(t *testing.T) {
	entries, err := os.ReadDir(TestSamplesDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range entries {
		sampleFileName := e.Name()

		if e.IsDir() || sampleFileName == CorrectSample {
			continue
		}

		t.Run(e.Name(), func(t *testing.T) {
			checkErr := RunCompatibilityTest(t, CorrectSample, sampleFileName)
			notNil := assert.NotNil(t, checkErr)
			if !notNil {
				correctData, err := os.ReadFile(filepath.Join(TestSamplesDir, CorrectSample))
				if err != nil {
					t.Fatal(err)
				}
				incorrectData, err := os.ReadFile(filepath.Join(TestSamplesDir, sampleFileName))
				if err != nil {
					t.Fatal(err)
				}
				t.Logf("file diff: %s\n", cmp.Diff(string(correctData), string(incorrectData)))
			}
		})
	}
}
