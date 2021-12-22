package model

import (
	"OceanLearn/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Username string `gorm:"type:varchar(100);unique_index;not null"`
	Password string `gorm:"size:255;not null"`
}

// IsExist 存在返回true
func (user User) IsExist(username string) bool {
	model := User{}
	common.GetDB().Where("username = ?", username).First(&model)
	if model.ID != 0 {
		return true
	} else {
		return false
	}
}

func (u User) CreateUser(user User) {
	common.GetDB().Create(&user)
}

func (user *User) GetUserByUsername(username string) {
	common.GetDB().Where("username = ?", username).First(&user)
}
