package goldmastertests

import (
	"testing"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/tl"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/tlab"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/tlcd"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/tlcurl"
)

func FuzzCurl(f *testing.F) {
	f.Fuzz(func(t *testing.T, original []byte) {
		req := tlcurl.Request{}
		_, err := req.ReadTL2(original, nil)
		if err != nil {
			return // t.Fatalf("read failed: %v", err)
		}
		_ = req.WriteTL2(nil, nil)
	})
}

func FuzzUnion(f *testing.F) {
	f.Fuzz(func(t *testing.T, original []byte) {
		req := tl.UnionArgsUse{}
		_, err := req.ReadTL2(original, nil)
		if err != nil {
			return // t.Fatalf("read failed: %v", err)
		}
		_ = req.WriteTL2(nil, nil)
	})
}

func FuzzUseCycle(f *testing.F) {
	f.Fuzz(func(t *testing.T, original []byte) {
		req := tlcd.UseCycle{}
		_, err := req.ReadTL2(original, nil)
		if err != nil {
			return // t.Fatalf("read failed: %v", err)
		}
		_ = req.WriteTL2(nil, nil)
	})
}

func FuzzUseDictString(f *testing.F) {
	f.Fuzz(func(t *testing.T, original []byte) {
		req := tlab.UseDictString{}
		_, err := req.ReadTL2(original, nil)
		if err != nil {
			return // t.Fatalf("read failed: %v", err)
		}
		_ = req.WriteTL2(nil, nil)
	})
}
