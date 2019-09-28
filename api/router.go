package api

import (
	"github.com/gin-gonic/gin"
)

// InitRouter init api
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/todos/:id", getTodo)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	return r
}

