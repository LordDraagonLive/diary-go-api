package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Entry struct {
	Id              int64     `json:"id"`
	Title           string    `json:"title"`
	Body            string    `json:"body"`
	Dedication      string    `json:"dedication"`
	CreatedDateTime time.Time `json:"createdDatTime"`
	UpdatedDateTime time.Time `json:"updatedDateTime"`
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return // if error, return
	}

	newEntry.CreatedDateTime = time.Now().Local() // set default value for CreatedDateTime
	newEntry.UpdatedDateTime = time.Now().Local() // set default value for UpdatedDateTime

	entries = append(entries, newEntry)          // append newEntry to entries
	c.IndentedJSON(http.StatusCreated, newEntry) // return newEntry with status code 201
}

func getEntryByIdHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64) // get id from url parameter and convert it to int64
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entry, err := getEntryById(id) // get entry by id
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, entry) // return entry with status code 200
}

/*
internal method to get entry by id
*/
func getEntryById(id int64) (*Entry, error) {
	for index, entry := range entries {
		if entry.Id == id {
			return &entries[index], nil // return pointer to entry `&entry` if you want to avoid modifying the original entry in entries array
		}
	}

	return nil, errors.New("Entry not found")
}

func main() {
	router := gin.Default()
	router.GET("/diary", getEntrys)
	router.POST("/diary", createEntry)
	router.GET("/diary/:id", getEntryByIdHandler)
	router.Run("localhost:8080")
}
