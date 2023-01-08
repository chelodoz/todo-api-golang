package todo

// Create a todo request
type CreateTodoRequest struct {
	// required: true
	Name string `json:"name" validate:"required"`
	// required: true
	Description string `json:"description" validate:"required"`
}

// Update a todo request
type UpdateTodoRequest struct {
	// required: true
	Name string `json:"name" validate:"required"`
	// required: true
	Description string `json:"description" validate:"required"`
}

// Create a todo response
// swagger:response CreateTodoResponse
type CreateTodoResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetTodoResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// List todos returns
// swagger:response GetTodosResponse
type GetTodosResponse []GetTodoResponse

// A single todo returns in the response
// swagger:response GetTodoByIdResponse
type GetTodoByIdResponse CreateTodoResponse

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
