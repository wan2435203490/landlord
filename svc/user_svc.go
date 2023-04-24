package svc

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"landlord/db"
	"landlord/sdk/service"
	"strings"
	"time"
)

type UserSvc struct {
	service.Service
}

func (s *UserSvc) InsertUser(user *db.User) error {
	return db.MySQL.Create(&user).Error
}

func (s *UserSvc) GetUser(user *db.User) error {

	err := db.MySQL.Find(&user, user.Id).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("用户不存在")
	}

	if err != nil {
		return err
	}
	return nil
}

func (s *UserSvc) GetOrInsertUser(user *db.User) error {

	err := db.MySQL.First(&user, "username=?", user.UserName).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		//第一次进来 用户不存在 自动插入
		user.Timestamp = time.Now()
		user.Id = strings.ReplaceAll(uuid.NewString(), "-", "")
		err = db.MySQL.Create(&user).Error
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *UserSvc) UpdateUser(user *db.User) error {
	return db.MySQL.Updates(&user).Error
}
