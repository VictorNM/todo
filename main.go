package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/victornm/todo/model/todo"
)

var db = make(map[int]*todo.Todo)

func getTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	if t, ok := db[id]; ok {
		c.JSON(200, gin.H{
			"data":  t,
			"error": nil,
		})
		return
	}

	c.JSON(404, nil)
}

func createTodo(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	j := string(body)
	_ = j
	defer c.Request.Body.Close()
	if err != nil {
		c.JSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}
	var t *todo.Todo
	err = json.Unmarshal(body, &t)
	if err != nil {
		c.JSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	t = todo.New(t.Title, t.Text)
	db[t.ID] = t
	c.JSON(200, gin.H{
		"data":  t,
		"error": nil,
	})
}

func updateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	j := string(body)
	_ = j
	defer c.Request.Body.Close()
	if err != nil {
		c.JSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	t, ok := db[id]
	if !ok {
		c.JSON(404, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	var t_ *todo.Todo
	err = json.Unmarshal(body, &t_)
	if err != nil {
		c.JSON(400, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	t.Title = t_.Title
	t.Text = t_.Text
	t.Complete = t_.Complete

	db[t.ID] = t
	c.JSON(200, gin.H{
		"data":  t,
		"error": nil,
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/todos/:id", getTodo)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:id", updateTodo)

	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080
}
