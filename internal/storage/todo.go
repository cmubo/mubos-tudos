package storage

import (
	"todo/internal/model"
	"todo/internal/types"
	"todo/internal/utils"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type TodoStore interface {
	GetTodo(int, *model.User) (*model.Todo, error)
	GetTodos(*model.User, types.Pagination, string, []types.Filter) ([]*model.Todo, int, error)
	CreateTodo(*model.Todo, *model.User) (*model.Todo, error)
	UpdateTodo(*model.Todo, *model.User) (*model.Todo, error)
	DeleteTodo(int, *model.User) error
}

func (s *Store) GetTodo(id int, user *model.User) (*model.Todo, error) {
	todo := model.Todo{}

	err := s.Db.Get(&todo, "SELECT * FROM todos WHERE id = $1 LIMIT 1", id)

	if err != nil {
		return nil, utils.DbErrorSinglularResource(err)
	}

	if todo.CreatedBy != user.Id {
		return nil, fiber.NewError(fiber.StatusUnauthorized, fiber.ErrUnauthorized.Message)
	}

	return &todo, nil
}

func (s *Store) GetTodos(user *model.User, pagination types.Pagination, orderByString string, filters []types.Filter) ([]*model.Todo, int, error) {
	todos := []*model.Todo{}
	offset := utils.GetPaginationOffset(pagination)
	limit := utils.GetPaginationLimit(pagination)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Select("*").From("todos").OrderBy(orderByString).Limit(uint64(limit)).Offset(uint64(offset)).Where("created_by = ?", user.Id)
	query = utils.AddFiltersToSquirrelQuery(query, filters)
	sql, args, err := query.ToSql()

	if err != nil {
		return utils.DbErrorMultiResource(err, todos)
	}

	err = s.Db.Select(&todos, sql, args...)

	if err != nil {
		return utils.DbErrorMultiResource(err, todos)
	}

	count, countErr := utils.DbGetCount(s.Db, "todos")
	if countErr != nil {
		return utils.DbErrorMultiResource(err, todos)
	}

	return todos, count, nil
}

func (s *Store) CreateTodo(todo *model.Todo, user *model.User) (*model.Todo, error) {
	connString := "INSERT INTO todos (title, description, completed, created_by) VALUES ($1,$2,$3,$4) RETURNING *"
	res := model.Todo{}
	err := s.Db.Get(&res, connString, todo.Title, todo.Description, todo.Completed, user.Id)

	if err != nil {
		log.Error(err)
		// TODO: error (if data is wrong, explicitly say it is)
		return nil, fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return &res, nil
}

func (s *Store) UpdateTodo(todo *model.Todo, user *model.User) (*model.Todo, error) {
	connString := `
	UPDATE todos 
	SET title = $1, description = $2, completed = $3 
	WHERE id = $4 AND created_by = $5
	RETURNING *
	`
	todores := model.Todo{}
	err := s.Db.Get(&todores, connString, todo.Title, todo.Description, todo.Completed, todo.Id, user.Id)
	if err != nil {
		return nil, utils.DbErrorSinglularResource(err)
	}

	return &todores, nil
}

func (s *Store) DeleteTodo(id int, user *model.User) error {
	connString := "DELETE FROM todos WHERE id = $1 AND created_by = $2"
	res, err := s.Db.Exec(connString, id, user.Id)

	if err != nil {
		return utils.DbErrorSinglularResource(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return utils.DbErrorSinglularResource(err)
	}

	if rowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "No resource was found that meet the required expectations.")
	}

	return nil
}
