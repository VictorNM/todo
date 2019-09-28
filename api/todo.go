package api

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
		response(c, 200, t, nil)
		return
	}

	response(c, 404, nil, errors.New("not found"))
}

func createTodo(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		response(c, 400, nil, err)
		return
	}
	var t *todo.Todo
	err = json.Unmarshal(body, &t)
	if err != nil {
		response(c, 400, nil, err)
		return
	}

	t = todo.New(t.Title, t.Text)
	db[t.ID] = t
	response(c, 201, t, nil)
}

func updateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response(c, 400, nil, err)
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		response(c, 400, nil, err)
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

func deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response(c, 400, nil, err)
		return
	}

	_, ok := db[id]
	if ok {
		delete(db, id)
	}

	response(c, 204, nil, nil)
}

func response(c *gin.Context, code int, data interface{}, err error) {
	c.JSON(code, gin.H{
		"data":  data,
		"error": err,
	})
}

