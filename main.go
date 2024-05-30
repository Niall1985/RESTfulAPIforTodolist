package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}

var todo_list = []todo{
	{ID: "1", Task: "Clean House", Status: false},
	{ID: "2", Task: "Fill Water", Status: false},
	{ID: "3", Task: "Tidy up the kitchen", Status: false},
}

func get_Todo_list(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todo_list)
}

func add_todo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todo_list = append(todo_list, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func delete_todo(context *gin.Context) {
	id := context.Param("id")
	for i, t := range todo_list {
		if t.ID == id {
			todo_list = append(todo_list[:i], todo_list[i+1:]...)
			context.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	for i, t := range todo_list {
		if t.ID == id {
			todo_list[i].Status = !todo_list[i].Status
			context.JSON(http.StatusOK, gin.H{"message": "Todo status toggled"})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func main() {
	router := gin.Default()
	router.GET("/todo_list", get_Todo_list)
	router.POST("/todo_list", add_todo)
	router.DELETE("/todo_list/:id", delete_todo)
	router.PUT("/todo_list/:id/toggle", toggleTodoStatus)
	router.Run("localhost:8000")

}
