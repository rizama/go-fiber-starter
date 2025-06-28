package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizama/go-fiber-starter/app/errors"
	"github.com/rizama/go-fiber-starter/app/response"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if appErr, ok := err.(*errors.AppError); ok {
		return response.ErrorResponse(c, appErr.Code, appErr.Err, appErr.Message)
	}

	if err.Error() == "Cannot "+c.Method()+" "+c.Path() {
		return response.ErrorResponse(c, fiber.StatusNotFound, err, "Route not found")
	}

	// Handle default error
	return response.ErrorResponse(c, fiber.StatusInternalServerError, err, "Internal Server Error")
}
