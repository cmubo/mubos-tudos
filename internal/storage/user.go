package storage

import (
	"fmt"
	"todo/internal/model"
	"todo/internal/types"
	"todo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UserStore interface {
	GetUser(int) (*model.User, error)
	GetUserFromEmail(string) (*model.User, error)
	GetUsers(types.Pagination, string, []types.Filter) ([]*model.User, int, error)
	CreateUser(*model.User) (*model.User, error)
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(id int) error
}

func (s *Store) GetUser(id int) (*model.User, error) {
	fmt.Println(id)
	user := model.User{}

	err := s.Db.Get(&user, "SELECT * FROM users WHERE id = $1 LIMIT 1", id)

	if err != nil {
		return nil, utils.DbErrorSinglularResource(err)
	}

	return &user, nil
}

func (s *Store) GetUserFromEmail(email string) (*model.User, error) {
	user := model.User{}

	err := s.Db.Get(&user, "SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		return nil, utils.DbErrorSinglularResource(err)
	}

	return &user, nil
}

func (s *Store) GetUsers(_ types.Pagination, _ string, _ []types.Filter) ([]*model.User, int, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Store) CreateUser(user *model.User) (*model.User, error) {
	connString := "INSERT INTO users (email,password) VALUES ($1,$2) RETURNING *"
	res := model.User{}
	err := s.Db.Get(&res, connString, user.Email, user.Password)

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			return nil, fiber.NewError(fiber.StatusConflict, "That email account is already in use")
		}

		log.Error(err)
		// TODO: error (if data is wrong, explicitly say it is)
		return nil, fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return &res, nil
}

func (s *Store) UpdateUser(_ *model.User) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Store) DeleteUser(id int) error {
	panic("not implemented") // TODO: Implement
}
