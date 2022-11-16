package ssh

import (
	"golang.org/x/crypto/ssh"
	"strconv"
)

type Config struct {
	Host       string
	Port       int
	User       string
	Password   string
	SuPassword string
}

func HostAddress(config Config) string {
	return config.Host + ":" + strconv.Itoa(config.Port)
}

func NewConfig(host string, port int, user string, pass string, supass string) Config {
	return Config{
		Host:       host,
		Port:       port,
		User:       user,
		Password:   pass,
		SuPassword: supass,
	}
}

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
