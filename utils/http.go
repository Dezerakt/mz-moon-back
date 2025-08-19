package utils

import (
	"github.com/gofiber/fiber"
	"music-streaming/config"
	errVo "music-streaming/utils/error"
	"net/http"
)

func ReturnError(f *fiber.Ctx, customError errVo.CustomError, err error) {
	storedError := errVo.Errors[customError]

	storedError.Message = err.Error()
	storedError.Type = customError

	f.Status(storedError.HttpCode).JSON(storedError)
}

type SuccessResponse struct {
	AppVersion string      `json:"appVersion,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func ReturnOk(f *fiber.Ctx, data interface{}) {
	successReponse := SuccessResponse{
		AppVersion: config.GetAppVersion(),
		Data:       data,
	}

	f.Status(http.StatusOK).JSON(successReponse)
}
