package main

import (
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

const ip = "151.248.113.144"
const port = "443"
const login = "iu9lab"
const password = "12345678990iu9iu9"

func main() {

	config := &ssh.ClientConfig{
		User: login,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", ip+":"+port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	sess, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer func(sess *ssh.Session) {
		err := sess.Close()
		if err != nil {
		}
	}(sess)
	stdin, err := sess.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		_, err := io.Copy(stdin, os.Stdin)
		if err != nil {
		}
	}()
	sess.Stdin = os.Stdin
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr
	modes := ssh.TerminalModes{}
	if err := sess.RequestPty("", 20, 40, modes); err != nil {
		log.Fatal(err)
	}
	if err := sess.Shell(); err != nil {
		log.Fatal(err)
	}
	if err = sess.Wait(); err != nil {
		log.Fatal(err)
	}
}
