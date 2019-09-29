package todo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository(t *testing.T)  {
	repo := NewInMemRepository()

	t.Run("Create", func(t *testing.T) {
		todo := New("hello", "world")
		res, _ := repo.Create(todo)
		assert.NotNil(t, res)
	})

	t.Run("Update", func(t *testing.T) {
		todo := New("hello", "world")
		res, _ := repo.Create(todo)
		res.Title = "updated"
		res2, _ := repo.Update(res)
		assert.NotNil(t, res2)
	})

	t.Run("Find", func(t *testing.T) {
		todo := New("hello", "world")
		res, _ := repo.Create(todo)
		res2, _ := repo.Find(res.ID)
		assert.NotNil(t, res2)
	})

	t.Run("FindAll", func(t *testing.T) {
		todo := New("hello", "world")
		_, _ = repo.Create(todo)
		todos, _ := repo.FindAll()
		assert.NotNil(t, todos)
		assert.True(t, len(todos) > 0)
	})

	t.Run("Delete", func(t *testing.T) {
		todo := New("hello", "world")
		res, _ := repo.Create(todo)
		err := repo.Delete(res.ID)
		assert.Nil(t, err)

		res2, _ := repo.Find(res.ID)
		assert.Nil(t, res2)
	})
}

