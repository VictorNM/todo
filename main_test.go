package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	body := strings.NewReader(`{"title":"hello","text":"world"}`)
	req, _ := http.NewRequest("POST", "/todos", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var m map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &m)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "hello", m["data"].(map[string]interface{})["title"])
}
