package testify

import (
	"github.com/stretchr/testify/assert"
	"test/src"
	"testing"
)

func TestAssert(t *testing.T) {
	assert.False(t, src.ShouldReturnFalse(""))
}
