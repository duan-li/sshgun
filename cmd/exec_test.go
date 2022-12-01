package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExec(t *testing.T) {
	ex := Exec()
	assert.NotNil(t, ex)
}
