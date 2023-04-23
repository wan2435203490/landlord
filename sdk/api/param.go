package api

import (
	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/gin-gonic/gin/binding"
)

// Param maybe must？
func (a *Api) Param(key string) string {

	value := a.Context.Param(key)
	if value == "" {
		a.ErrorMsg(key + " is empty")
	}
	return value
}

// Bind 参数校验
func (a *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	if d == nil {
		return a
	}
	var err error
	if len(bindings) == 0 {
		bindings = constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = a.Context.ShouldBindUri(d)
		} else {
			err = a.Context.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			a.Logger.Warn("request body is not present anymore. ")
			err = nil
			continue
		}
		if err != nil {
			a.AddError(err)
			break
		}
	}
	//vd.SetErrorFactory(func(failPath, msg string) error {
	//	return fmt.Errorf(`"validation failed: %s %s"`, failPath, msg)
	//})
	if err1 := vd.Validate(d); err1 != nil {
		a.AddError(err1)
	}
	a.ErrorIfExists()
	return a
}