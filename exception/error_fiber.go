package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-mongo-clean/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationErr)
	if ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
