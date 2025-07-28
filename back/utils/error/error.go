package errVo

import "net/http"

type CustomError string

type ErrorMetadata struct {
	Type     CustomError `json:"type,omitempty"`
	HttpCode int         `json:"httpCode,omitempty"`
	Message  string      `json:"message,omitempty"`
}

const (
	InternalError       = "E_INTERNAL_ERROR"
	UnauthorizedError   = "E_UNAUTHORIZED_ERROR"
	BadRequestError     = "E_BAD_REQUEST"
	InvalidParams       = "E_INVALID_PARAMS"
	ValidateParamsError = "E_VALIDATE_PARAMS"
)

var Errors = map[CustomError]ErrorMetadata{
	InternalError: {
		HttpCode: http.StatusInternalServerError,
	},
	UnauthorizedError: {
		HttpCode: http.StatusUnauthorized,
	},
	BadRequestError: {
		HttpCode: http.StatusBadRequest,
	},
	InvalidParams: {
		HttpCode: http.StatusBadRequest,
	},
	ValidateParamsError: {
		HttpCode: http.StatusBadRequest,
	},
}
