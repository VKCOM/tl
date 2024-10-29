package codegen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import "github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/meta"

func TestFunctionHasUnion(t *testing.T) {
	{
		// test type which contains union in its recursive definition
		fun := meta.FactoryItemByTLName("service5.insertList")
		if assert.NotNil(t, fun) {
			assert.True(t, fun.IsFunction())
			assert.True(t, fun.HasUnionTypesInResult())
		}
	}
	{
		// test type which doesn't contain union in its recursive definition
		fun := meta.FactoryItemByTLName("usefulService.getUserEntity")
		if assert.NotNil(t, fun) {
			assert.True(t, fun.IsFunction())
			assert.True(t, !fun.HasUnionTypesInResult())
		}
	}
	{
		// test type which contains enum in its recursive definition
		fun := meta.FactoryItemByTLName("ab.call10")
		if assert.NotNil(t, fun) {
			assert.True(t, fun.IsFunction())
			assert.True(t, fun.HasUnionTypesInResult())
		}
	}
}

func TestFunctionHasUnionInArguments(t *testing.T) {
	{
		// test type which contains union in its recursive definition
		fun := meta.FactoryItemByTLName("service5.insertList")
		if assert.NotNil(t, fun) {
			assert.True(t, fun.IsFunction())
			assert.True(t, !fun.HasUnionTypesInArguments())
		}
	}
	{
		// test type which doesn't contain union in its recursive definition
		fun := meta.FactoryItemByTLName("ab.call11")
		if assert.NotNil(t, fun) {
			assert.True(t, fun.IsFunction())
			assert.True(t, fun.HasUnionTypesInArguments())
		}
	}
}
