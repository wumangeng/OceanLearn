package service

import (
	"OceanLearn/model"
	"OceanLearn/util"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateUser(user model.User) {
	password := user.Password
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("密码加密错误！")
	}
	user.Password = string(newPassword)
	user.CreateUser(user)
}

func UserLogin(username string, password string) string {
	//校验密码
	user := model.User{}
	user.GetUserByUsername(username)
	if user.ID == 0 {
		panic("账号不存在")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		panic("密码错误")
	}

	//发放令牌
	token, err := util.GenerateToken(user)
	if err != nil {
		panic("Fail generate token ")
		log.Panicf("Token generate error:", err)
	}
	return token
}
