package util

import (
	"OceanLearn/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// jwt密钥
var secretKey = []byte("ocean_learn_secre_key")

type Claims struct {
	UserId uint
	Name   string
	jwt.StandardClaims
}

func GenerateToken(user model.User) (string, error) {
	//设置token的过期时间
	expirationTime := time.Now().Add(6 * time.Hour)

	claims := &Claims{
		UserId: user.ID,
		Name:   user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),     //token发放时间
			Issuer:    "oceanlearn",          //签发人
			Subject:   "user token",          //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用密钥生成token
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil

}

func ParesToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	return token, claims, err
}
