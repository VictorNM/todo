package todo

// Todo present a todo
type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Complete bool   `json:"complete"`
}

var curID = 0

// New todo
func New(title string, text string) *Todo {
	curID++
	return &Todo{
		ID:       curID,
		Title:    title,
		Text:     text,
		Complete: false,
	}
}
