package api

import (
	"errors"
	"gorm.io/gorm"
	"landlord/sdk/pkg"
	"net/http"
)

// GetOrm 获取Orm DB
func (a Api) GetOrm() (*gorm.DB, error) {
	db, err := pkg.GetOrm(a.Context)
	if err != nil {
		a.Logger.Error(http.StatusInternalServerError, err, "数据库连接获取失败")
		return nil, err
	}
	return db, nil
}

// MakeOrm 设置Orm DB
func (a *Api) MakeOrm() *Api {
	var err error
	if a.Logger == nil {
		err = errors.New("at MakeOrm logger is nil")
		a.AddError(err)
		return a
	}
	db, err := pkg.GetOrm(a.Context)
	if err != nil {
		a.Logger.Error(http.StatusInternalServerError, err, "数据库连接获取失败")
		a.AddError(err)
	}
	a.Orm = db
	return a
}
