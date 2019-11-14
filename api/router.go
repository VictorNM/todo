package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/victornm/todo"
)

// InitRouter init api
func InitRouter() *gin.Engine {
	controller := NewTodoController(todo.NewInMemRepository())

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/todos", controller.getTodos)
	r.GET("/todos/:id", controller.getTodo)
	r.POST("/todos", controller.createTodo)
	r.PUT("/todos/:id", controller.updateTodo)
	r.DELETE("/todos/:id", controller.deleteTodo)

	return r
}

