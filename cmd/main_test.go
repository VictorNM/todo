package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/victornm/todo/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTodo(t *testing.T) {
	r := router.Init()
	w := post(r, "hello", "world")
	w = get(r, 1)

	res := parse(w.Body)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "hello", res["data"].(map[string]interface{})["title"])
}

func TestCreateTodo(t *testing.T) {
	r := router.Init()
	w := post(r, "hello", "world")

	res := parse(w.Body)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "hello", res["data"].(map[string]interface{})["title"])
}

func TestUpdateTodo(t *testing.T) {
	t.Run("Update todo succeed", func(t *testing.T) {
		r := router.Init()
		w := post(r, "hello", "world")
		w = put(r, 1, "goodbye", "world", false)

		res := parse(w.Body)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "goodbye", res["data"].(map[string]interface{})["title"])
	})

	t.Run("Update todo 404", func(t *testing.T) {
		r := router.Init()
		w := post(r, "hello", "world")
		w = put(r, 10, "goodbye", "world", false)

		res := parse(w.Body)

		assert.Equal(t, 404, w.Code)
		assert.NotNil(t, res["error"])
	})
}

func TestDeleteTodo(t *testing.T) {
	r := router.Init()
	post(r, "hello", "world")
	w := delete(r, 1)

	assert.Equal(t, 204, w.Code)
}

func get(router *gin.Engine, id int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/todos/%d", id)
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	return w
}

func post(router *gin.Engine, title, text string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	url := "/todos"
	body := strings.NewReader(fmt.Sprintf(`{"title":"%s","text":"%s"}`, title, text))
	req, _ := http.NewRequest("POST", url, body)
	router.ServeHTTP(w, req)

	return w
}

func put(router *gin.Engine, id int, title, text string, complete bool) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/todos/%d", id)
	body := strings.NewReader(fmt.Sprintf(`{"title":"%s","text":"%s","complete":%t}`, title, text, complete))
	req, _ := http.NewRequest("PUT", url, body)
	router.ServeHTTP(w, req)

	return w
}

func delete(router *gin.Engine, id int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/todos/%d", id)
	req, _ := http.NewRequest("DELETE", url, nil)
	router.ServeHTTP(w, req)

	return w
}

func parse(body *bytes.Buffer) map[string]interface{} {
	var m map[string]interface{}
	_ = json.Unmarshal(body.Bytes(), &m)
	return m
}