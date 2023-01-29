package kcg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidGroup(t *testing.T) {
	assert.True(t, validGroup("valid", []string{"valid"}))
	assert.False(t, validGroup("invalid", []string{"valid"}))
}
