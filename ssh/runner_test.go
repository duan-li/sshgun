package ssh

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnection(t *testing.T) {
	config := NewConfig("host", 22, "user", "pass", "supass")

	out, err := Runner(config, "ls")

	// session will not be established because of wrong host
	assert.NotNil(t, err)
	assert.Equal(t, "", out)
}
