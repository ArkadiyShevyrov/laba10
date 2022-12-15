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

type UsersOnline struct {
	Online bool `json:"online"`
	Count  int  `json:"count"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	server := gin.New()

	server.GET("/", UsersOnlineHandler)

	err := server.Run(":8098")
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

func getUserOnline() (users UsersOnline) {
	names := getUsersFromResponse(getResponse())
	users = UsersOnline{
		Online: false,
		Count:  0,
	}
	for _, name := range names {
		if name == "iu9lab" {
			users.Online = true
			users.Count += 1
		}
	}
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
	output, err := sess.Output("w")
	if err != nil {
		log.Println(err)
	}
	response = string(output)
	return
}

func getUsersFromResponse(str string) (res []string) {
	split := strings.Split(str, "\n")
	for i, s := range split {
		if i == 0 || i == 1 || s == "" {
			continue
		}
		name := strings.Split(s, " ")[0]
		if name == "" {
			continue
		}
		res = append(res, name)
	}
	return
}
