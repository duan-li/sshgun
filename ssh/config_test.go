package ssh

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig("host", 22, "user", "pass", "supass")

	assert.NotNil(t, config)

	assert.Equal(t, "host", config.Host)
	assert.Equal(t, 22, config.Port)
	assert.Equal(t, "user", config.User)
	assert.Equal(t, "pass", config.Password)
	assert.Equal(t, "supass", config.SuPassword)
}

func TestNewSSHConfig(t *testing.T) {
	config := NewConfig("host", 22, "user", "pass", "supass")

	sshConfig := NewSSHConfig(config)

	assert.Equal(t, "user", sshConfig.User)
	assert.NotNil(t, sshConfig.Auth)
	assert.NotEmpty(t, sshConfig.Auth)

}
