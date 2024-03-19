package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"todo/internal/types"

	"github.com/gofiber/fiber/v2"
)

func JsonResponse[T any](ctx *fiber.Ctx, data T, rData types.Response) error {
	return ctx.JSON(fiber.Map{
		"status":  rData.Status,
		"message": rData.Message,
		"data":    data,
	})
}

// Returns format: "created_at DESC", acceptedSortMethods map should be "method": "DESC" | "ASC"
func GetSortByString(query string, defaultSort string, acceptedSortMethods map[string]string) string {
	isWhitespacePresent := regexp.MustCompile(`\s`).MatchString(query)

	if query == "" || isWhitespacePresent {
		// If there is whitespace or the query is empty, return default.
		return defaultSort
	}

	items := strings.Split(query, ",")

	var sortItems []string

	for _, item := range items {
		str, err := createSortString(item, acceptedSortMethods)
		if err != nil {
			return defaultSort
		}

		sortItems = append(sortItems, str)
	}

	return strings.Join(sortItems, ",")
}

func createSortString(item string, acceptedSortMethods map[string]string) (string, error) {
	splitItem := strings.Split(item, ".")

	if len(splitItem) < 2 {
		return "", errors.New("either no period or no direction was provided into the sort parameter")
	}

	if _, ok := acceptedSortMethods[splitItem[0]]; !ok {
		return "", fmt.Errorf("sort method %s not part of accepted methods", splitItem[0])
	}

	splitItem[1] = strings.ToUpper(splitItem[1])
	if splitItem[1] != "DESC" && splitItem[1] != "ASC" {
		splitItem[1] = acceptedSortMethods[splitItem[0]]
	}

	return splitItem[0] + " " + splitItem[1], nil
}
