package entity

type Todo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedOn   string `json:"-"`
}
