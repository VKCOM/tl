package genphp

import (
	"github.com/VKCOM/tl/internal/tlast"
)

type FileToWrite struct {
	Path string
	Ast  tlast.TL2File
}

func (gen *Gen2) MigrateToTL2(prevState []FileToWrite) (newState []FileToWrite, _ error) {
	return nil, nil
}
