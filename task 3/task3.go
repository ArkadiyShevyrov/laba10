package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const (
	password string = "Je2dTYr6"
	login    string = "iu9networkslabs"
	host     string = "students.yss.su"
	dbname   string = "iu9networkslabs"
)

type Product struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	DatePublished string `json:"datePublished"`
	TimeRecords   string `json:"timeRecords"`
	Link          string `json:"link"`
}

type Products struct {
	Products []Product `json:"products"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	server := gin.New()

	server.GET("/", UsersOnlineHandler)

	err := server.Run(":8095")
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

func getUserOnline() (users Products) {
	users = Products{
		Products: getResponse(),
	}
	return
}

func getResponse() (products []Product) {
	db, err := sql.Open("mysql", login+":"+password+"@tcp("+host+")/"+dbname)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from iu9networkslabs.iu9Shevyrov")
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		p := Product{}
		err4 := rows.Scan(&p.Id, &p.Title, &p.Content, &p.DatePublished, &p.TimeRecords, &p.Link)
		if err4 != nil {
			fmt.Println(err4)
			continue
		}
		products = append(products, p)
	}
	return
}
