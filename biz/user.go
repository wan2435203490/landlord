package biz

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"landlord/db"
	"landlord/db/mysql_model"
	"time"
)

// GenerateUser Only For Test
func GenerateUser(userId string) (*db.User, error) {
	user := &db.User{Id: userId, Timestamp: time.Now(), UserName: "test" + userId}
	return mysql_model.InsertUser(user)
}

func InsertUser(user *db.User) (*db.User, error) {
	return mysql_model.InsertUser(user)
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

func GetUserByName(userName string) (*db.User, error) {
	user, err := mysql_model.GetUserByName(userName)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(user *db.User) {
	mysql_model.UpdateUser(user)
}
