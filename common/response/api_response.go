package r

import (
	"github.com/gin-gonic/gin"
	"landlord/common/utils"
	"net/http"
)

const (
	SuccessCode = 0
	ErrorCode   = -1
)

type ApiResponse[T any] struct {
	Code    int    `json:"code"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func Success0[T any](data T) *ApiResponse[T] {

	res := &ApiResponse[T]{
		Data:    data,
		Message: http.StatusText(http.StatusOK),
		Status:  SuccessCode,
		Code:    http.StatusOK,
	}

	return res
}

func Error0(httpStatus int, msg string) *ApiResponse[struct{}] {

	//if httpStatus < 0 {
	//	panic(...)
	//}

	res := &ApiResponse[struct{}]{
		Message: utils.IfThen(len(msg) > 0, msg, http.StatusText(httpStatus)).(string),
		Status:  ErrorCode,
		Code:    httpStatus,
	}

	return res
}

func Success[T any](data T, c *gin.Context) {

	res := &ApiResponse[T]{
		Data:    data,
		Message: http.StatusText(http.StatusOK),
		Status:  SuccessCode,
		Code:    http.StatusOK,
	}

	c.JSON(http.StatusOK, res)

}

func Error(httpStatus int, msg string, c *gin.Context) {

	res := &ApiResponse[struct{}]{
		Message: utils.IfThen(len(msg) > 0, msg, http.StatusText(httpStatus)).(string),
		Status:  ErrorCode,
		Code:    httpStatus,
	}

	c.JSON(httpStatus, res)
}

func ErrorInternal(msg string, c *gin.Context) {

	res := &ApiResponse[struct{}]{
		Message: utils.IfThen(len(msg) > 0, msg, http.StatusText(http.StatusInternalServerError)).(string),
		Status:  ErrorCode,
		Code:    http.StatusInternalServerError,
	}

	c.JSON(http.StatusInternalServerError, res)
}
