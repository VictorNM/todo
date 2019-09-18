package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/victornm/todo/model/todo"
)

var db = make(map[int]*todo.Todo)

type todoForm struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Complete bool   `json:"complete"`
}

func getTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{
			"data": nil,
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
	var tf todoForm
	err = json.Unmarshal(body, &tf)
	fmt.Println(tf)
	if err != nil {
		c.JSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	t := todo.New(tf.Title, tf.Text)
	db[t.ID] = t
	c.JSON(200, gin.H{
		"data":  t,
		"error": nil,
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/todo/:id", getTodo)
	r.POST("/todos", createTodo)

	return r
}

func main() {
	r := setupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
