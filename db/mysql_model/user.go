package mysql_model

import (
	"landlord/db"
	"time"
)

func InsertUser(user *db.User) (*db.User, error) {

	user.Timestamp = time.Now()
	err := db.DB.Mysql.DB.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return GetUser(user.Id)
}

func GetUser(userId string) (*db.User, error) {
	var user *db.User
	err := db.DB.Mysql.DB.Table("user").Where("id=?", userId).Take(&user).Error

	//err := db.DB.MySQL.DB.Find(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByName(userName string) (*db.User, error) {
	var user *db.User
	err := db.DB.Mysql.DB.Table("user").Where("username=?", userName).Take(&user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *db.User) {
	
	err := db.DB.Mysql.DB.Table("user").Updates(&user).Error

	if err != nil {
		panic(err)
	}
}
