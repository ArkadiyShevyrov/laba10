package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"log"
	"net/http"
	"strings"
)

const ip = "151.248.113.144"
const port = "443"
const login = "iu9lab"
const password = "12345678990iu9iu9"

type Message struct {
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	server := gin.New()

	server.GET("/", UsersOnlineHandler)

	err := server.Run(":8091")
	if err != nil {
		log.Fatalln(err)
	}
}

func UsersOnlineHandler(ctx *gin.Context) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	for {
		err = ws.WriteJSON(getUserOnline())
		if err != nil {
			fmt.Println(err)
			break
		}

	}
}

func getUserOnline() (users Message) {
	users = Message{
		Message: "",
	}
	res := getFilenamesFromResponse(getResponse())
	for _, re := range res {
		if re == "achtung.txt" {
			users.Message = getFile()
		}
	}
	return
}

func getFile() (res string) {
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

	output, err := sess.Output("cat achtung.txt")
	if err != nil {
		log.Println(err)
	}
	res = string(output)
	return
}

func getResponse() (response string) {
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

	output, err := sess.Output("ls")
	if err != nil {
		log.Println(err)
	}
	response = string(output)
	return
}

func getFilenamesFromResponse(str string) (res []string) {
	split := strings.Split(str, "\n")
	for _, s := range split {
		if s == "" {
			continue
		}
		res = append(res, s)
	}
	return
}
