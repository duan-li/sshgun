package ssh

import (
	"golang.org/x/crypto/ssh"
)

func Runner(config Config, command string) (string, error) {
	sshConfig := NewSSHConfig(config)
	conn, err := ssh.Dial("tcp", HostAddress(config), sshConfig)
	if err != nil {
		return "", err
	}

	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	out, err := session.CombinedOutput(command)
	if err != nil {
		panic(err)
	}

	return string(out), err
}
