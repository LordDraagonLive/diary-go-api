package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Entry struct {
	Id              int64     `json:"id"`
	Title           string    `json:"title"`
	Body            string    `json:"body"`
	Dedication      string    `json:"dedication"`
	CreatedDateTime time.Time `json:"created_date_time"`
	UpdatedDateTime time.Time `json:"updated_date_time"`
}

var entries = []Entry{
	{Id: 1, Title: "Entry 1", Body: "This is entry 1.", Dedication: "To my wife", CreatedDateTime: time.Now().Local(), UpdatedDateTime: time.Now().Local()},
	{Id: 2, Title: "Entry 2", Body: "This is entry 2.", Dedication: "To my kids", CreatedDateTime: time.Now().Local(), UpdatedDateTime: time.Now().Local()},
	{Id: 3, Title: "Entry 3", Body: "This is entry 3.", Dedication: "To my dog", CreatedDateTime: time.Now().Local(), UpdatedDateTime: time.Now().Local()},
}

func getEntrys(c *gin.Context) { // c is a Context of gin
	c.IndentedJSON(http.StatusOK, entries)
}

func createEntry(c *gin.Context) {
	var newEntry Entry

	if err := c.BindJSON(&newEntry); err != nil { // BindJSON binds the request body to newEntry. & is a pointer to newEntry
		return // if error, return
	}

	newEntry.CreatedDateTime = time.Now().Local() // set default value for CreatedDateTime
	newEntry.UpdatedDateTime = time.Now().Local() // set default value for UpdatedDateTime

	entries = append(entries, newEntry)          // append newEntry to entries
	c.IndentedJSON(http.StatusCreated, newEntry) // return newEntry with status code 201
}

func getEntryById(c *gin.Context) {

}

func main() {
	router := gin.Default()
	router.GET("/diary", getEntrys)
	router.POST("/diary", createEntry)
	router.Run("localhost:8080")
}
