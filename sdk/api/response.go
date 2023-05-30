package api

import (
	"landlord/common/utils"
	"log"
	"net/http"
)

const (
	SuccessStatus = 0
	ErrorStatus   = -1
)

type Response struct {
	Code    int    `json:"code,omitempty"`
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type response struct {
	Response
	Data any `json:"data"`
}

func (a *Api) OK(data any) {
	res := &response{}
	res.Message = http.StatusText(http.StatusOK)
	res.Status = SuccessStatus
	res.Code = http.StatusOK
	res.Data = data

	a.Context.JSON(http.StatusOK, res)
	log.Printf("%#v\n", data)
}

func (a *Api) Error(httpStatus int, msg string) {
	res := &response{}
	res.Message = utils.IfThen(len(msg) > 0, msg, http.StatusText(httpStatus)).(string)
	res.Status = ErrorStatus
	res.Code = httpStatus

	a.Context.JSON(httpStatus, res)
	log.Printf("%#v\n", msg)
}

func (a *Api) ErrorInternal(msg string) {
	a.Error(http.StatusInternalServerError, msg)
}
