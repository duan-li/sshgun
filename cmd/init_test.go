package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	i := Init()
	assert.NotNil(t, i)
}
