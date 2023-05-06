package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Dedication string `json:"dedication"`
}

var pages = []Page{
	{Id: 1, Title: "Page 1", Body: "This is page 1.", Dedication: "To my wife"},
	{Id: 2, Title: "Page 2", Body: "This is page 2.", Dedication: "To my kids"},
	{Id: 3, Title: "Page 3", Body: "This is page 3.", Dedication: "To my dog"},
}

func getPages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pages)
}

func main() {
	router := gin.Default()
	router.GET("/pages", getPages)
	router.Run("localhost:8080")
}
