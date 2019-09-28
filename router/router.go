package router

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/victornm/todo"
	"io/ioutil"
	"strconv"
)

var db = make(map[int]*todo.Todo)

func getTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response(c, 400, nil, err)
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
		response(c, 404, nil, errors.New("not found"))
		return
	}

	var t_ *todo.Todo
	err = json.Unmarshal(body, &t_)
	if err != nil {
		response(c, 400, nil, err)
		return
	}

	t.Title = t_.Title
	t.Text = t_.Text
	t.Complete = t_.Complete

	db[t.ID] = t
	response(c, 200, t, nil)
}

func response(c *gin.Context, code int, data interface{}, err error) {
	c.JSON(code, gin.H{
		"data":  data,
		"error": err,
	})
}

// Init init router
func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/todos/:id", getTodo)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:id", updateTodo)

	return r
}

