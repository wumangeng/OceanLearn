package main

import (
	"OceanLearn/common"
	"OceanLearn/model"
	"OceanLearn/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig() //加载配置文件

	db := common.InitDB() //连接数据库
	db.AutoMigrate(&model.User{})
	defer db.Close()

	r := gin.Default() //加载gin框架
	router.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		r.Run(":" + port)
	}
	r.Run()
}

func InitConfig() {
	//获取当前工作目录
	workDir, _ := os.Getwd()

	viper.SetConfigName("application")         //读取文件名设置
	viper.SetConfigType("yml")                 //读取文件类型设置
	viper.AddConfigPath(workDir + "/resource") //设置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprint("配置文件读取失败 err=", err))
	}
}
