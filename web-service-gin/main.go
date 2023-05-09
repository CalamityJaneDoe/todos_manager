// Please note that this is an API for purpose of a mini demo project
// This is not a production ready code
// What is missing / have to be corrected for a production ready API :
// 1. Implement PATCH route (to have a full CRUD API)
// 2. Dialog with db instead of seeding data into a slice
// 3. The logic within PUT endpoints is a little bit "Hackish" + error handling is missing

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ------- 1. Define data struct and seed data for project purpose (normally API'd interact with DB) ------

// todo represents data struct of a todo.
type todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Urgent      bool   `json:"urgent"`
	Subject     string `json:"subject"`
}

// todo slice to seed todos data.
var todos = []todo{
	{ID: "1", Title: "Buy bread", Completed: true, Urgent: false, Subject: "perso"},
	{ID: "2", Title: "Call my sister", Description: "Organize next family event", Completed: false, Urgent: false, Subject: "perso"},
	{ID: "3", Title: "Read deep C secrets", Description: "Gain some knowledge on C", Completed: false, Urgent: false, Subject: "pro"},
	{ID: "4", Title: "Appointment Doctor", Completed: true, Urgent: false, Subject: "perso"},
	{ID: "5", Title: "Buy flowers", Description: "Offer them to grandma", Completed: false, Urgent: true, Subject: "perso"},
	{ID: "6", Title: "Read bio Evan You", Description: "Gain some knowledge on Vue.JS", Completed: false, Urgent: false, Subject: "pro"},
	{ID: "7", Title: "Go to the mall", Description: "We need some food !", Completed: false, Urgent: true, Subject: "perso"},
	{ID: "8", Title: "Buy concert tickets", Description: "This is gonna be awesome", Completed: false, Urgent: false, Subject: "perso"},
	{ID: "9", Title: "Try the new udemy course", Description: "Have to learn go", Completed: false, Urgent: true, Subject: "pro"},
	{ID: "10", Title: "Make some smash potatoes", Description: "Crap ! We are having friends over to dinner", Completed: true, Urgent: true, Subject: "perso"},
	{ID: "11", Title: "Go to see Dartagnan", Description: "'Cos everybody loves Louis Garrel", Completed: true, Urgent: false, Subject: "perso"},
	{ID: "12", Title: "Buy some shoes", Description: "Nothing to wear for my uncle's wedding", Completed: false, Urgent: false, Subject: "perso"},
	{ID: "13", Title: "Write a card to Lea", Completed: false, Urgent: false, Subject: "perso"},
	{ID: "14", Title: "Write a blog post", Description: "Subject : is it easy to code APi with go ?", Completed: false, Urgent: false, Subject: "pro"},
	{ID: "15", Title: "Clean the garage", Description: "Upcomming home hackaton", Completed: false, Urgent: false, Subject: "perso"},
	{ID: "16", Title: "Find a cute bow", Description: "I love to disguise my cat", Completed: false, Urgent: false, Subject: "perso"},
	{ID: "17", Title: "Get some cofee", Description: "Oh boy I need it !", Completed: false, Urgent: true, Subject: "pro"},
	{ID: "18", Title: "Read a snobish book", Description: "Maybe a Duras ?", Completed: false, Urgent: false, Subject: "perso"},
	{ID: "19", Title: "Clean up the Kanban", Description: "Prepare next sprint !", Completed: false, Urgent: true, Subject: "pro"},
	{ID: "20", Title: "Finish the Vue.js/Go mini project", Completed: false, Urgent: true, Subject: "pro"},
}

// ------ 2. Assign handlers (see bellow) to endpoint paths ------

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todo/:id", getTodoByID) //:smth = this is a path parameter
	router.POST("/todo", postTodo)
	router.PUT("/todo/:id", putTodoByID)
	router.DELETE("/todo/:id", deleteTodoByID)
	//@TODO : add a PATCH endpoint

	router.Run("0.0.0.0:8080") // listen on all network interfaces
}

// ------ 3. Define handlers ------

// getTodos responds with the list of all todos as JSON.
func getTodos(c *gin.Context) { // gin.Context carries request details, validates and serializes JSON, and more.
	c.IndentedJSON(http.StatusOK, todos)
}

// getTodoByID locates the Todo whose ID value matches the id
// parameter sent by the client, then returns that todo as a response.
func getTodoByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of todos, looking for
	// an album whose ID value matches the parameter.
	for _, a := range todos { //Note that in a real world project, there'd be a database query instead of this logic
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

// postTodo adds a todo from JSON received in the request body.
func postTodo(c *gin.Context) {
	var newTodo todo

	// Call BindJSON to bind the received JSON to
	// newTodo.
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	// Add the new todo to the slice.
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// putTodoByID locates the Todo whose ID value matches the id
// parameter sent by the client, then update that todo from JSON received in the request body.
func putTodoByID(c *gin.Context) {
	id := c.Param("id")
	// Locate and delete the todo to modify
	for _, a := range todos { //Note that in a real world project, there'd be a database query instead of this logic
		if a.ID == id {
			if s, err := strconv.Atoi(a.ID); err == nil { // id is a string so we have to change it to an int
				todos = append(todos[:s-1], todos[s:]...) // Trick with append() in order to remove the todo from the slice
				// see https://stackoverflow.com/questions/25025409/delete-element-in-a-slice for explanation
			}
		}
	}
	// add the modified todo received in the request body to the slice
	var newTodo todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusCreated, newTodo)
	//@ TODO : error handling is missing !
}

// deleteTodoByID locates the Todo whose ID value matches the id
// parameter sent by the client, then delete that todo.
func deleteTodoByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of todos, looking for
	// a todo whose ID value matches the parameter.
	for _, a := range todos { //Note that in a real world project, there'd be a database query instead of this logic
		if a.ID == id {
			var idx = SliceIndex(len(todos), func(i int) bool { return todos[i] == a }) // define idx for the todo
			todos = append(todos[:idx], todos[idx+1:]...)                               // suppress the todo
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

// 4. ------ Handlers Utils ------

// SliceIndex calculate the index of an element of a Slice
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
