package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var source = "./dashboard/src/resources"

func main() {
	server := gin.New()

	server.LoadHTMLGlob(source + "/templates/*.html")

	server.GET("/dashboard", Dashboard)

	err := server.Run(":8090")
	if err != nil {
		log.Fatalln(err)
	}
}

func Dashboard(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
