package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitBullet(t *testing.T) {
	f := InitBullet()
	assert.NotNil(t, f)
}

func TestValidate(t *testing.T) {
	f := Validate()
	assert.NotNil(t, f)
}

func TestCheck(t *testing.T) {
	f := Check()
	assert.NotNil(t, f)
}
