package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()             // Initialize the database connection and create necessary tables
	server := gin.Default() // Create a new Gin router instance

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) // Bind the JSON request body to the event struct

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

/*

Preparing Statements vs Directly Executing Queries (Prepare() vs Exec()/Query())

- We started sending SQL commands to the SQLite database.
And we did this by following different approaches:

	1. DB.Exec() (when we created the tables)

	2. Prepare() + stmt.Exec() (when we inserted data into the database)

	3. DB.Query() (when we fetched data from the database)

Using Prepare() is 100% optional! You could send all your commands directly via Exec() or Query().

- The difference between those two methods then just is whether you're fetching data from the database (=> use Query())
or your manipulating the database / data in the database (=> use Exec()).

But what's the advantage of using Prepare()?

- Prepare() prepares a SQL statement - this can lead to better performance if the same statement is executed multiple times (potentially with different data for its placeholders).
This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions. In that case, there wouldn't be any advantages.
And, indeed, in this application, we are calling stmt.Close() directly after calling stmt.Exec(). So here, it really wouldn't matter which approach you're using.
But in order to show you the different ways of using the sql package, I decided to also include this preparation approach in this course.

*/
