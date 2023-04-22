package db

import (
	"fmt"
	"github.com/dtm-labs/rockscache"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"landlord/common/config"
	"landlord/db/driver"
	"log"
	"time"
)

var DB DataBases

type DataBases struct {
	Mysql       driver.MySqlDB
	mgoSession  *mgo.Session
	mongoClient *mongo.Client
	RDB         *redis.UniversalClient
	Rc          *rockscache.Client
	WeakRc      *rockscache.Client
}

func init() {
	initMysqlDB()
}

type Writer struct{}

func (w Writer) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func initMysqlDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Config.MySQL.UserName, config.Config.MySQL.Password, config.Config.MySQL.Address[0], "mysql")
	log.Println("------------------------", dsn)
	var db *gorm.DB
	var err1 error
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		time.Sleep(time.Duration(30) * time.Second)
		db, err1 = gorm.Open(mysql.Open(dsn), nil)
		if err1 != nil {
			panic(err1.Error() + " open failed " + dsn)
		}
	}

	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default charset utf8 COLLATE utf8_general_ci;", config.Config.MySQL.DatabaseName)
	err = db.Exec(sql).Error
	if err != nil {
		panic(err.Error() + " Exec failed " + sql)
	}
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Config.MySQL.UserName, config.Config.MySQL.Password, config.Config.MySQL.Address[0], config.Config.MySQL.DatabaseName)
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             time.Duration(config.Config.MySQL.SlowThreshold) * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.LogLevel(config.Config.MySQL.LogLevel),                       // Log level
			IgnoreRecordNotFoundError: true,                                                                // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                                                                // Disable color
		},
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err.Error() + " Open failed " + dsn)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error() + " db.DB() failed ")
	}

	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.Config.MySQL.MaxLifeTime))
	sqlDB.SetMaxOpenConns(config.Config.MySQL.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.Config.MySQL.MaxIdleConns)

	db.AutoMigrate(&User{}, &Achievement{})
	db.Set("gorm:table_options", "CHARSET=utf8")
	db.Set("gorm:table_options", "collation=utf8_unicode_ci")

	DB.Mysql.DB = db
}

func DefaultGormDB() *gorm.DB {
	return DB.Mysql.DB
}
