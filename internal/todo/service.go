package todo

import "sample-golang-api/internal/entity"

type TodoService interface {
	CreateTodo(todo entity.Todo) (*entity.Todo, error)
	GetTodoById(id uint) (*entity.Todo, error)
	GetTodos() ([]*entity.Todo, error)
}

type todoService struct {
	todoRepository TodoRepository
}

func NewTodoService(repository TodoRepository) TodoService {
	return &todoService{
		todoRepository: repository,
	}
}

func (service *todoService) CreateTodo(todo entity.Todo) (*entity.Todo, error) {
	newTodo, err := service.todoRepository.CreateTodo(todo)

	if err != nil {
		return nil, err
	}

	return newTodo, nil
}

func (service *todoService) GetTodoById(id uint) (*entity.Todo, error) {

	todoById, err := service.todoRepository.GetTodoById(id)

	if err != nil {
		return nil, err
	}

	return todoById, nil
}
func (service *todoService) GetTodos() ([]*entity.Todo, error) {

	todos, err := service.todoRepository.GetTodos()

	if err != nil {
		return nil, err
	}

	return todos, nil
}
