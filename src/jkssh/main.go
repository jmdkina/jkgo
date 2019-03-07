package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"jk/jklog"
	"net"
	"os"
	//"strings"
	"time"
)

type JKSSH struct {
	addr     string
	port     int
	username string
	password string
	session  *ssh.Session
	c        *ssh.Client
}

func NewJKSSH(addr string, port int, username, password string) (*JKSSH, error) {
	//var hostKey ssh.PublicKey
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	cConfig := &ssh.ClientConfig{
		User:    username,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	fulladdr := fmt.Sprintf("%s:%d", addr, port)
	c, err := ssh.Dial("tcp", fulladdr, cConfig)
	if err != nil {
		return nil, err
	}

	session, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	jkssh := &JKSSH{
		addr:     addr,
		port:     port,
		username: username,
		password: password,
		session:  session,
		c:        c,
	}
	return jkssh, nil
}

func (s *JKSSH) Send(data string) (string, error) {
	var b bytes.Buffer
	ss, _ := s.c.NewSession()
	ss.Stdout = &b
	err := ss.Run(data)
	if err != nil {
		return "", err
	}
	ss.Close()
	return b.String(), nil
}

func (s *JKSSH) Close() {
	s.session.Close()
	s.c.Close()
}

func main() {
	var addr = flag.String("addr", "", "remote address")
	var port = flag.Int("port", 22, "remote port")
	var username = flag.String("username", "root", "remote user")
	var password = flag.String("password", "", "remote password")

	flag.Parse()
	jkssh, err := NewJKSSH(*addr, *port, *username, *password)
	if err != nil {
		jklog.L().Errorln("Error create ssh ", err)
		return
	}

	read := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		inputcmd2, _, _ := read.ReadLine()
		inputcmd := string(inputcmd2)
		if inputcmd == "exit" {
			break
		}
		buf, err := jkssh.Send(inputcmd)
		if err != nil {
			jklog.L().Errorln("Error send data ", err)
			break
		}
		fmt.Println(buf)
	}

	jkssh.Close()
}
