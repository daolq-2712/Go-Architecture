package common

import (
	"Go-Architecture/common/constants"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	StatusCode int       `json:"status_code"`
	Message    string    `json:"message"`
	Key        string    `json:"error_key"`
	RootCause  RootCause `json:"root_cause"`
}

type RootCause struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func AbortRequest(c *gin.Context, appErr *AppError) {
	if err := c.AbortWithError(appErr.StatusCode, appErr); err != nil {
		return
	}
}

func BadRequest(root error, msg *string) *AppError {
	var message = constants.ERROR_MSG_BAD_REQUEST
	if msg != nil {
		message = *msg
	}

	defaultError := errors.New(message)
	if root == nil {
		root = defaultError
	}
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootCause:  getRootCause(root),
		Message:    message,
		Key:        "ERROR_BAD_REQUEST",
	}
}

func ServerError(root error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		RootCause:  getRootCause(root),
		Message:    "Oops. Có lỗi xảy ra, xin vui lòng thử lại sau.",
		Key:        "ERROR_INTERNAL_SERVER",
	}
}

func CreationFailed(root error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		RootCause:  getRootCause(root),
		Message:    "Tạo bản ghi thất bại",
		Key:        "ERROR_INTERNAL_SERVER",
	}
}

func UpdateFailed(root error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		RootCause:  getRootCause(root),
		Message:    "Cập nhật bản ghi thất bại",
		Key:        "ERROR_INTERNAL_SERVER",
	}
}

func NotFound(root error, msg *string) *AppError {
	var message = constants.ERROR_MSG_NOT_FOUND
	if msg != nil {
		message = *msg
	}

	defaultError := errors.New(message)
	if root == nil {
		root = defaultError
	}
	return &AppError{
		StatusCode: http.StatusNotFound,
		RootCause:  getRootCause(root),
		Message:    message,
		Key:        "ERROR_NOT_FOUND",
	}
}

func (e *AppError) Error() string {
	return e.Message
}

func getRootCause(rootError error) RootCause {
	return RootCause{
		rootError.Error(),
		rootError,
	}
}
