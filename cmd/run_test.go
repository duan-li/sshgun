package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	f := Run()
	assert.NotNil(t, f)
}
