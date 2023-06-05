package driver

import (
	"gorm.io/gorm"
)

type MySqlDB struct {
	//sync.RWMutex
	DB *gorm.DB
}
