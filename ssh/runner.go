package ssh

import (
	"bufio"
	"golang.org/x/crypto/ssh"
	"io"
	"strings"
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

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		return "", err
	}

	in, err := session.StdinPipe()
	if err != nil {
		panic(err)
	}

	out, err := session.StdoutPipe()
	if err != nil {
		panic(err)
	}

	var output []byte

	go func(in io.WriteCloser, out io.Reader, output *[]byte) {
		var (
			line string
			r    = bufio.NewReader(out)
		)
		for {
			b, err := r.ReadByte()
			if err != nil {
				break
			}

			*output = append(*output, b)

			if b == byte('\n') {
				line = ""
				continue
			}

			line += string(b)
			if strings.HasPrefix(line, "[sudo] password for ") && strings.HasSuffix(line, ": ") {
				_, err = in.Write([]byte(config.Password + "\n"))
				if err != nil {
					break
				}
			}
		}
	}(in, out, &output)

	_, err = session.Output(command)
	if err != nil {
		return "", err
	}

	return string(output), err
}
