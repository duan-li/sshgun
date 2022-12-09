package ssh

import (
	"golang.org/x/crypto/ssh"
	"strconv"
)

// Config is the configuration struct for the SSH connection
type Config struct {
	Host       string
	Port       int
	User       string
	Password   string
	SuPassword string
}

// HostAddress returns the host address
func HostAddress(config Config) string {
	return config.Host + ":" + strconv.Itoa(config.Port)
}

// NewConfig returns a new SSH configuration
func NewConfig(host string, port int, user string, pass string, supass string) Config {
	return Config{
		Host:       host,
		Port:       port,
		User:       user,
		Password:   pass,
		SuPassword: supass,
	}
}

// NewSSHConfig create new SSH configuration
func NewSSHConfig(config Config) *ssh.ClientConfig {
	sshConfig := &ssh.ClientConfig{
		User: config.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return sshConfig
}
