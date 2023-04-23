package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
	r "landlord/common/response"
	"landlord/core/logger"
	"landlord/core/tools/language"
	"landlord/sdk/pkg"
	"landlord/sdk/pkg/response"
	"landlord/sdk/service"
)

const (
	DefaultLanguage = "zh-CN"
	UserSessionKey  = "curUser"
)

type Api struct {
	Context *gin.Context
	Logger  *logger.Helper
	Svc     *service.Service
	Orm     *gorm.DB
	Err     error
}

func (a *Api) AddError(err error) {
	if a.Err == nil {
		a.Err = err
	} else if err != nil {
		a.Logger.Error(err)
		a.Err = fmt.Errorf("%v; %w", a.Err, err)
	}
}

func (a *Api) Build(c *gin.Context, s service.IService, d interface{}, bindings ...binding.Binding) error {
	return a.MakeContext(c).MakeOrm().Bind(d, bindings...).MakeService(s.Get()).Err
}

// MakeContext 设置http上下文
func (a *Api) MakeContext(c *gin.Context) *Api {
	a.Context = c
	a.Logger = GetRequestLogger(c)
	return a
}

// GetLogger 获取上下文提供的日志
func (a Api) GetLogger() *logger.Helper {
	return GetRequestLogger(a.Context)
}

// ErrorIfExists err 不为空 return
func (a *Api) ErrorIfExists() {
	if a.Err != nil {
		a.Error(a.Err)
	}
}

func (a *Api) MakeService(c *service.Service) *Api {
	c.Log = a.Logger
	c.Orm = a.Orm
	return a
}

// Error 通常错误数据处理
func (a Api) Error(err error) {
	r.ErrorInternal(err.Error(), a.Context)
}
func (a Api) ErrorMsg(msg string) {
	r.ErrorInternal(msg, a.Context)
}

//func (a Api) Err(code int, err error, msg string) {
//	//responsa.Err(a.Context, code, err, msg)
//}

// Error 通常错误数据处理
func (a Api) Success(data any) {
	if a.Err == nil {
		r.Success(data, a.Context)
	} else {
		a.ErrorIfExists()
	}
}

// OK 通常成功数据处理
func (a Api) OK(data interface{}, msg string) {
	response.OK(a.Context, data, msg)
}

// PageOK 分页数据处理
func (a Api) PageOK(result interface{}, count int, pageIndex int, pageSize int, msg string) {
	response.PageOK(a.Context, result, count, pageIndex, pageSize, msg)
}

// Custom 兼容函数
func (a Api) Custom(data gin.H) {
	response.Custum(a.Context, data)
}

func (a Api) Translate(form, to interface{}) {
	pkg.Translate(form, to)
}

// getAcceptLanguage 获取当前语言
func (a *Api) getAcceptLanguage() string {
	languages := language.ParseAcceptLanguage(a.Context.GetHeader("Accept-Language"), nil)
	if len(languages) == 0 {
		return DefaultLanguage
	}
	return languages[0]
}
