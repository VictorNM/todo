package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/victornm/todo"
	"io/ioutil"
	"strconv"
)

type TodoController struct {
	repo todo.Repository
}

// NewTodoController constructor
func NewTodoController(repo todo.Repository) *TodoController {
	return &TodoController{repo:repo}
}

func (controller *TodoController) getTodos(c *gin.Context) {
	todos, err := controller.repo.FindAll()
	if err != nil {
		response(c, 400, nil, err)
	}

	response(c, 200, todos, nil)
}

func (controller *TodoController) getTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response(c, 400, nil, err)
		return
	}

	t, err := controller.repo.Find(id)
	if err != nil {
		response(c, 404, nil, err)
	}

	response(c, 200, t, nil)
}

func (controller *TodoController) createTodo(c *gin.Context) {
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
	t, err = controller.repo.Create(t)
	if err != nil {
		response(c, 400, nil, err)
		return
	}

	response(c, 201, t, nil)
}

func (controller *TodoController) updateTodo(c *gin.Context) {
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

	t, err := controller.repo.Find(id)
	if err != nil {
		response(c, 404, nil, err)
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

	t, err = controller.repo.Update(t)
	if err != nil {
		response(c, 400, nil, err)
		return
	}

	response(c, 200, t, nil)
}

func (controller *TodoController) deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response(c, 400, nil, err)
		return
	}

	err = controller.repo.Delete(id)
	if err != nil {
		response(c, 400, nil, err)
	}
	response(c, 204, nil, nil)
}

func response(c *gin.Context, code int, data interface{}, err error) {
	c.JSON(code, gin.H{
		"data":  data,
		"error": err,
	})
}

