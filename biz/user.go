package biz

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"landlord/db"
	"landlord/db/mysql_model"
	"landlord/sdk/service"
	"strings"
	"time"
)

type UserSvc struct {
	service.Service
}

// GenerateUser Only For Test
func (s *UserSvc) GenerateUser(userId string) (*db.User, error) {
	user := &db.User{Id: userId, Timestamp: time.Now(), UserName: "test" + userId}
	return mysql_model.InsertUser(user)
}

func (s *UserSvc) InsertUser(user *db.User) (*db.User, error) {
	return mysql_model.InsertUser(user)
}

func (s *UserSvc) GetUser(userId string) (*db.User, error) {
	user, err := mysql_model.GetUser(userId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户不存在")
		} else {
			return nil, err
		}
	}

	return user, nil
}

func GetUser(userId string) (*db.User, error) {
	user, err := mysql_model.GetUser(userId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户不存在")
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (s *UserSvc) GetOrInsertUser(user *db.User) bool {

	err := s.Orm.First(&user, "username=?", user.UserName).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		//第一次进来 用户不存在 自动插入
		user.Timestamp = time.Now()
		user.Id = strings.ReplaceAll(uuid.NewString(), "-", "")
		err = s.Orm.Create(&user).Error
	}

	if err != nil {
		s.Error(err)
		return false
	}

	return true
}

func (s *UserSvc) UpdateUserBiz(user *db.User) {
	UpdateUser(user)
}

func UpdateUser(user *db.User) {
	mysql_model.UpdateUser(user)
}
