package controller

import (
	"fmt"
	"strconv"
	"todo/internal/constants"
	"todo/internal/model"
	"todo/internal/types"
	"todo/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// GetTodos godoc
//
//	@Summary		Get Todos
//	@Description	Get a list of todos
//	@Tags			todo
//	@Accept			json
//	@Produce		json
//	@Param			page	query	integer		false	"Page number, defaults to 1"
//	@Param			per_page	query	integer		false	"Results per request, defaults to 20, max 50."
//	@Param			filter	query	string		false	"Filters seperated via &. example: ?filter[eq]=completed=true&balance[gte]=50."
//	@Param			sort_by		query	string		false	"sort the results by. string value of the db column name and then either DESC or ASC for direction. Example: created_at DESC."
//	@Success		200	{array}	model.Todo
//	@Failure		500	{object}	string
//	@Router			/todo [get]
func (h *Handler) GetTodos(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.User)

	page := utils.GetPaginationQuery(c.Query("page"), 1)
	perPage := utils.GetPaginationQuery(c.Query("per_page"), constants.PAGINATION_PERPAGE_DEFAULT)
	filters := utils.GetFiltersFromQuery(c.Query("filter"), constants.TodoAcceptedFilters)
	sortByString := utils.GetSortByString(c.Query("sort_by"), "created_at DESC", constants.AcceptedSortMethodsTodo)

	todos, count, err := h.Storage.GetTodos(user, types.Pagination{Page: page, PerPage: perPage}, sortByString, filters)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	c.Set("Results-Count", fmt.Sprint(count))

	return utils.JsonResponse(c, todos, constants.RESPONSE_SUCCESFULLY_RETRIEVED)
}

// GetTodo godoc
//
//	@Summary		Get a single todo
//	@Description	Get a single todo
//	@Tags			todo
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Todo
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/todo/{id} [get]
func (h *Handler) GetTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.User)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	todo, err := h.Storage.GetTodo(id, user)
	if err != nil {
		return err
	}

	return utils.JsonResponse(c, todo, constants.RESPONSE_SUCCESFULLY_RETRIEVED)
}

// CreateTodo godoc
//
//	@Summary		Create a Todo
//	@Description	Create a new todo item
//	@Tags			todo
//	@Accept			json
//	@Produce		json
//	@Param			user	body	model.CreateTodo	true	"Add todo"
//	@Success		200	{object}	model.Todo
//	@Failure		400	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/todo [post]
func (h *Handler) CreateTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.User)

	body := model.Todo{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	err = h.Validator.Validate(body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	todo, err := h.Storage.CreateTodo(&body, user)
	if err != nil {
		return err
	}

	return utils.JsonResponse(c, todo, constants.RESPONSE_SUCCESFULLY_CREATED)
}

// GetTodo godoc
//
//	@Summary		Delete a todo
//	@Description	Delete a single todo
//	@Tags			todo
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Failure		500	{object}	string
//	@Router			/todo/{id} [delete]
func (h *Handler) DeleteTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.User)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	err = h.Storage.DeleteTodo(id, user)
	if err != nil {
		return err
	}

	return utils.JsonResponse(c, "", constants.RESPONSE_SUCCESFULLY_DELETED)
}

// CreateTodo godoc
//
//	@Summary		Update a Todo
//	@Description	Update a new todo item
//	@Tags			todo
//	@Accept			json
//	@Produce		json
//	@Param			user	body	model.CreateTodo	true	"Update todo"
//	@Success		200	{object}	model.Todo
//	@Failure		400	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/todo/{id} [put]
func (h *Handler) UpdateTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.User)

	body := model.Todo{}
	err := c.BodyParser(&body)
	if err != nil {
		// TODO: This would return: json: cannot unmarshal string into Go struct field Todo.id of type int
		// Could potentially check the field thats of the wrong format using a regex or something and implement a nicer message.
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if body.Id == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "[Id integer] You need to provide a valid \"Id\" in a json format inside the body of the request.")
	}

	err = h.Validator.Validate(body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	todo, err := h.Storage.UpdateTodo(&body, user)
	if err != nil {
		return err
	}

	return utils.JsonResponse(c, todo, constants.RESPONSE_SUCCESFULLY_UPDATED)
}
