package controller

import (
	"OceanLearn/common"
	"OceanLearn/model"
	"OceanLearn/service"
	"OceanLearn/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(ctx *gin.Context) {
	//获取参数
	name := ctx.PostForm("name")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//数据验证
	if len(username) > 15 || len(username) < 8 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "账号长度必须为8-15位"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能不少于6位"})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	user := model.User{}
	//判断账号是否存在
	isExist := user.IsExist(username)
	if isExist {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "账号已存在"})
		return
	}

	//创建用户
	user.Name = name
	user.Username = username
	user.Password = password
	service.CreateUser(user)

	common.Ok(ctx, nil, "注册成功")
}

func LoginUser(ctx *gin.Context) {
	//获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "账号密码不能为空")
	}
	token := service.UserLogin(username, password)
	common.Ok(ctx, token, "登陆成功")
}

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	common.Ok(ctx, user, "当前用户信息")
}
