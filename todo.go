package todo

import (
	"errors"
)

// Todo present a todo
type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Complete bool   `json:"complete"`
}

var curID = 0

func nextID() int {
	curID++
	return curID
}

// New Todo
func New(title string, text string) *Todo {
	return &Todo{
		ID:       nextID(),
		Title:    title,
		Text:     text,
		Complete: false,
	}
}

type Repository interface {
	Find(id int) (*Todo, error)
	FindAll() ([]*Todo, error)
	Create(entity *Todo) (*Todo, error)
	Delete(id int) error
	Update(entity *Todo) (*Todo, error)
}

type InMemRepository struct{
	db map[int]*Todo
}

func NewInMemRepository() *InMemRepository {
	return &InMemRepository{db:make(map[int]*Todo)}
}

func (r *InMemRepository) Find(id int) (*Todo, error) {
	if todo, ok := r.db[id]; ok {
		return todo, nil
	}

	return nil, errors.New("model not found")
}

func (r *InMemRepository) FindAll() ([]*Todo, error) {
	var todos []*Todo
	for _, todo := range r.db {
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *InMemRepository) Create(entity *Todo) (*Todo, error) {
	if _, ok := r.db[entity.ID]; ok {
		return nil, errors.New("model existed")
	}

	r.db[entity.ID] = entity
	return entity, nil
}

func (r *InMemRepository) Update(entity *Todo) (*Todo, error) {
	if _, ok := r.db[entity.ID]; !ok {
		return nil, errors.New("model not found")
	}

	r.db[entity.ID] = entity
	return entity, nil
}

func (r *InMemRepository) Delete(id int) error {
	delete(r.db, id)
	return nil
}
