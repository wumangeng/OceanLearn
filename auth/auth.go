package auth

import (
	"OceanLearn/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenStr := ctx.GetHeader("Authorization")

		//验证是否携带 token 或者token格式
		if tokenStr == "" || strings.HasPrefix(tokenStr, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			//抛弃请求
			ctx.Abort()
			return
		}

		//解析token
		tokenStr = tokenStr[7:] //去掉token头部Bearer
		token, claims, err := util.ParesToken(tokenStr)
		if err != nil || !token.Valid { //解析失败，或者token无效
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			//抛弃请求
			ctx.Abort()
			return
		}

		//token验证通过，将用户信息写入上下文
		ctx.Set("user", claims)
		ctx.Next()
	}
}
