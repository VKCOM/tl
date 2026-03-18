package linter

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/VKCOM/tl/internal/tlast"
	"github.com/VKCOM/tl/internal/tlcodegen"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

const SamplesDir = "../../tls/backward_compatibility_samples"
const CorrectSamplesDir = SamplesDir + "/correct-changes"
const IncorrectSamplesDir = SamplesDir + "/incorrect-changes"
const PrototypeSample = "prototype.tl"

func RunCompatibilityTest(t *testing.T, prevStateFile, newStateFile string) *tlast.ParseError {
	prevTL, err := readTL(prevStateFile)
	if err != nil {
		t.Fatal(err)
	}
	newTL, err := readTL(newStateFile)
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
	protoPath := filepath.Join(SamplesDir, PrototypeSample)

	t.Run("correct-changes", func(t *testing.T) {
		generalCheckTest(t, protoPath, CorrectSamplesDir, func(t *testing.T, parseError *tlast.ParseError) bool {
			return assert.Nil(t, parseError)
		})
	})
	t.Run("incorrect-changes", func(t *testing.T) {
		generalCheckTest(t, protoPath, IncorrectSamplesDir, func(t *testing.T, parseError *tlast.ParseError) bool {
			return assert.NotNil(t, parseError)
		})
	})
}

func generalCheckTest(t *testing.T, protoPath string, samplesDir string, assertion func(*testing.T, *tlast.ParseError) bool) {
	entries, err := os.ReadDir(samplesDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range entries {
		sampleFileName := e.Name()
		samplePath := filepath.Join(samplesDir, sampleFileName)

		if e.IsDir() {
			continue
		}

		t.Run(e.Name(), func(t *testing.T) {
			checkErr := RunCompatibilityTest(t, protoPath, samplePath)
			assertResult := assertion(t, checkErr)
			if !assertResult {
				t.Logf("provided assertion failed\n")
				correctData, err := os.ReadFile(protoPath)
				if err != nil {
					t.Fatal(err)
				}
				incorrectData, err := os.ReadFile(samplePath)
				if err != nil {
					t.Fatal(err)
				}
				t.Logf("file diff: %s\n", cmp.Diff(string(correctData), string(incorrectData)))
			}
		})
	}
}
