package mysql_model

import (
	"landlord/db"
)

func GetAchievementByUserId(userId string) (*db.Achievement, error) {
	var achievement db.Achievement
	err := db.DB.Mysql.DB.Table("achievement").Where("user_id=?", userId).First(&achievement).Error

	if err != nil {
		return nil, err
	}
	return &achievement, nil
}

func InsertAchievement(achievement *db.Achievement) error {

	err := db.DB.Mysql.DB.Table("achievement").Create(&achievement).Error

	return err
}

func UpdateAchievement(achievement *db.Achievement) {

	err := db.DB.Mysql.DB.Table("achievement").Updates(&achievement).Error

	if err != nil {
		panic(err)
	}
}
