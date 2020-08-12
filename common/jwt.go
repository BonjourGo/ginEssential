package common

import (
	"ginEssential/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("s_gdjJ_jsbj")

type Claim struct {
	UserId uint
	jwt.StandardClaims
}

func GetToken(user model.User) (string, error) {
	// 设置过期时间
	uTime := time.Now().Add(7 * 25 * time.Hour)
	claim := &Claim{
		UserId:         user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: uTime.Unix(), // 过期时间
			IssuedAt: time.Now().Unix(), // 发放时间
			Issuer: "Bonjour",
			Subject: "token",
		},
	}
	// 获取token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claim, error) {
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})
	return token, claim, err
}