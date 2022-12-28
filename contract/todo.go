package contract

type CreateTodoRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateTodoResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetTodosResponse []*CreateTodoResponse

type GetTodoByIdResponse CreateTodoResponse
