package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"landlord/core/logger"
	"landlord/sdk/pkg"
)

const (
	DefaultLanguage = "zh-CN"
)

type Api struct {
	Context *gin.Context
	//Logger  *logger.Helper
	Orm *gorm.DB
	Err error
}

func (a *Api) AddError(err error) {
	if err == nil {
		return
	}
	if a.Err == nil {
		a.Err = err
	} else {
		a.Err = fmt.Errorf("%v; %w", a.Err, err)
	}
	return
}

// IsError 如果有error则write json response
func (a *Api) IsError(err error) bool {
	if err != nil {
		a.ErrorInternal(err.Error())
		return true
	}
	return false
}

//func (a *Api) Build(c *gin.Context, s service.IService, d interface{}, bindings ...binding.Binding) error {
//	return a.MakeContext(c).MakeOrm().Bind(d, bindings...).MakeService(s.Get()).Err
//}

// MakeContext 设置http上下文
func (a *Api) MakeContext(c *gin.Context) error {
	a.Context = c
	//a.Logger = GetRequestLogger(c)

	err := a.MakeOrm()
	if err != nil {
		return err
	}

	return nil
}

// MakeOrm 设置Orm DB
func (a *Api) MakeOrm() error {
	var err error
	db, err := pkg.GetOrm(a.Context)
	if err != nil {
		//s.Api.Logger.Error(http.StatusInternalServerError, err, "数据库连接获取失败")
		return err
	}
	a.Orm = db
	return nil
}

// GetLogger 获取上下文提供的日志
func (a *Api) GetLogger() *logger.Helper {
	return GetRequestLogger(a.Context)
}

//
//func (a *Api) MakeService(c *service.Service) *Api {
//	c.Log = a.Logger
//	c.Orm = a.Orm
//	c.Context = a.Context
//	return a
//}
