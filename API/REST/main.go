package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// event represents data about a record album.
type event struct {
	ID  int `json:"ID"`
	X   int `json:"X"`
	Y   int `json:"Y"`
	RES int `json:"RES"`
}

// events slice to seed record album data.
var events = []event{
	{ID: 0, X: 0, Y: 0, RES: 0},
	{ID: 1, X: 1, Y: 1, RES: 2},
}

func main() {
	router := gin.Default()
	router.GET("/events/show", getEvents)
	router.GET("/events/show/:ID", getEventsByID)
	router.POST("/events/add", addEvents)

	router.Run("localhost:8080")
}

// getEvents responds with the list of all events as JSON.
func getEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

// addEvents adds an event from JSON received in the request body.
func addEvents(c *gin.Context) {
	var newEvent event
	// ID, _ := strconv.Atoi(events[len(events)-1].ID)
	// ID = ID + 1
	// X, _ := strconv.Atoi(c.Param("X"))
	// Y, _ := strconv.Atoi(c.Param("X"))
	// RES := X + Y

	// Call BindJSON to bind the received JSON to
	// newEvent.
	if err := c.BindJSON(&newEvent); err != nil {
		return
	}
	// newEvent.ID = strconv.Itoa(ID)
	// newEvent.RES = strconv.Itoa(RES)
	// Add the new event to the slice.
	events = append(events, newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

// getEventsByID locates the event whose ID value matches the id
// parameter sent by the client, then returns that event as a response.
func getEventsByID(c *gin.Context) {
	id := c.Param("ID")
	ID, _ := strconv.Atoi(id)

	// Loop through the list of events, looking for
	// an event whose ID value matches the parameter.
	for _, a := range events {
		if a.ID == ID {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found"})
}
