package controller

import (
	"fmt"
	"strconv"
	"todo/internal/constants"
	"todo/internal/model"
	"todo/internal/types"
	"todo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (h *Handler) GetTodos(c *fiber.Ctx) error {
	page := utils.GetPaginationQuery(c.Query("page"), 1)
	perPage := utils.GetPaginationQuery(c.Query("per_page"), constants.PAGINATION_PERPAGE_DEFAULT)

	todos, count, err := h.storage.GetTodos(types.Pagination{Page: page, PerPage: perPage})
	if err != nil {
		return err
	}

	c.Set("Results-Count", fmt.Sprint(count))

	return utils.JsonResponse(c, todos, constants.RESPONSE_SUCCESFULLY_RETRIEVED)
}

func (h *Handler) GetTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	todo, err := h.storage.GetTodo(id)
	if err != nil {
		return err
	}

	return utils.JsonResponse(c, todo, constants.RESPONSE_SUCCESFULLY_RETRIEVED)
}

func (h *Handler) CreateTodo(c *fiber.Ctx) error {
	body := model.Todo{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	todo, err := h.storage.CreateTodo(&body)
	if err != nil {
		log.Error(err)
		return err
	}

	return utils.JsonResponse(c, todo, constants.RESPONSE_SUCCESFULLY_CREATED)
}

func (h *Handler) DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	err = h.storage.DeleteTodo(id)
	if err != nil {
		return err
	}

	return utils.JsonResponse(c, make([]model.Todo, 0), constants.RESPONSE_SUCCESFULLY_DELETED)
}

func (h *Handler) UpdateTodo(c *fiber.Ctx) error {
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

	todo, err := h.storage.UpdateTodo(&body)
	if err != nil {
		return err
	}

	return utils.JsonResponse(c, todo, constants.RESPONSE_SUCCESFULLY_UPDATED)
}
