package storage

import (
	"todo/internal/model"
	"todo/internal/types"
	"todo/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type TodoStore interface {
	GetTodo(int) (*model.Todo, error)
	GetTodos(types.Pagination) ([]*model.Todo, int, error)
	CreateTodo(*model.Todo) (*model.Todo, error)
	UpdateTodo(*model.Todo) (*model.Todo, error)
	DeleteTodo(id int) error
}

func (s *Store) GetTodo(id int) (*model.Todo, error) {
	todo := model.Todo{}
	err := s.Db.Get(&todo, "SELECT * FROM todos WHERE id = $1", id)

	if err != nil {
		return nil, utils.DbErrorSinglularResource(err)
	}

	return &todo, nil
}

func (s *Store) GetTodos(pagination types.Pagination) ([]*model.Todo, int, error) {
	todos := []*model.Todo{}
	offset := utils.GetPaginationOffset(pagination)
	limit := utils.GetPaginationLimit(pagination)

	// TODO: order by parameter
	err := s.Db.Select(&todos, "SELECT * FROM todos LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return utils.DbErrorMultiResource(err, todos)
	}

	count, countErr := utils.DbGetCount(s.Db, "todos")
	if countErr != nil {
		return utils.DbErrorMultiResource(err, todos)
	}

	return todos, count, nil
}

func (s *Store) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	connString := "INSERT INTO todos (title, description, completed) VALUES ($1,$2,$3) RETURNING *"
	todores := model.Todo{}
	err := s.Db.Get(&todores, connString, todo.Title, todo.Description, todo.Completed)

	if err != nil {
		// TODO: error (if data is wrong, explicitly say it is)
		return nil, fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return &todores, nil
}

func (s *Store) UpdateTodo(todo *model.Todo) (*model.Todo, error) {
	connString := `
	UPDATE todos 
	SET title = $1, description = $2, completed = $3 
	WHERE id = $4 
	RETURNING *
	`
	todores := model.Todo{}
	err := s.Db.Get(&todores, connString, todo.Title, todo.Description, todo.Completed, todo.Id)
	if err != nil {
		// TODO: error (if data is wrong, explicitly say it is)
		return nil, fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return &todores, nil
}

func (s *Store) DeleteTodo(id int) error {
	connString := "DELETE FROM todos WHERE id = $1"
	_, err := s.Db.Exec(connString, id)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}
