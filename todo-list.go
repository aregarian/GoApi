package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID        string
	Item      string
	Completed bool
}

var todos = []Todo{
	{ID: "1", Item: "HEhe boy", Completed: false},
	{ID: "2", Item: "HEhe girl", Completed: false},
	{ID: "3", Item: "HEhe child", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)

}

func addTodos(context *gin.Context) {
	var newTodo Todo

	if error := context.BindJSON(&newTodo); error != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodos)
	router.GET("/todo/:id", getTodo)
	router.Run("localhost:9090")

}
