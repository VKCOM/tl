package codegen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import "github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/meta"

func TestFunctionHasUnion(t *testing.T) {
	{
		fun := meta.FactoryItemByTLName("service5.insertList")
		if assert.NotNil(t, fun) {
			assert.True(t, fun.IsFunction())
			assert.True(t, fun.HasUnionTypesInResult())
		}
	}
	{
		fun := meta.FactoryItemByTLName("usefulService.getUserEntity")
		if assert.NotNil(t, fun) {
			assert.True(t, fun.IsFunction())
			assert.True(t, !fun.HasUnionTypesInResult())
		}
	}
}
