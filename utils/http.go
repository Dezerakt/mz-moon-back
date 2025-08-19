package utils

import (
	"github.com/gofiber/fiber"
	"mz-moon-back/config"
	errVo "mz-moon-back/utils/error"
	"net/http"
)

func ReturnError(c *fiber.Ctx, customError errVo.CustomError, err error) {
	storedError := errVo.Errors[customError]

	storedError.Message = err.Error()
	storedError.Type = customError

	c.Status(storedError.HttpCode).JSON(storedError)
}

type SuccessResponse struct {
	AppVersion string      `json:"appVersion,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func ReturnOk(c *fiber.Ctx, data interface{}) {
	successReponse := SuccessResponse{
		AppVersion: config.GetAppVersion(),
		Data:       data,
	}

	c.Status(http.StatusOK).JSON(successReponse)
}
