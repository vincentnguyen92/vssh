package client

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"golang.org/x/crypto/ssh"
)

type SSHClient struct {
	config   *ssh.ClientConfig
	host     string
	port     int
	username string
	filePem  string
}

func NewSSHClient(host string, port int, username string, filePem string) *SSHClient {
	s := SSHClient{host: host, port: port, username: username, filePem: filePem}

	return &s
}

func (s *SSHClient) Dial(method string) (*ssh.Client, error) {
	// Use filepem to create new config SSHClient
	// TODO: We should detect and change when using only username and password
	// or attach with tunnel
	s.setConfigPublicKey(s.filePem)

	conn, err := ssh.Dial(method, s.host+":"+strconv.Itoa(s.port), s.config)
	if err != nil {
		fmt.Errorf("cannot connect: %v", err)
		os.Exit(1)
	}

	return conn, err
}

func (s *SSHClient) getPublicFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}

func (s *SSHClient) setConfigPublicKey(file string) {
	s.config = &ssh.ClientConfig{
		User: s.username,
		Auth: []ssh.AuthMethod{
			s.getPublicFile(file),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}
