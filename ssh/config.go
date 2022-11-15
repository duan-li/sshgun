package ssh

import "golang.org/x/crypto/ssh"

type Config struct {
	Host       string
	Port       int
	User       string
	Password   string
	SuPassword string
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
