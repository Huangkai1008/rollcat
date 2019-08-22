package utils

import (
	"github.com/dgrijalva/jwt-go"
	"rollcat/pkg/constants"
	"rollcat/pkg/setting"
	"time"
)

type Claims struct {
	UserId   uint
	Username string
	jwt.StandardClaims
}

// 生成token
func GenerateToken(userId uint, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(constants.JwtExpireDuration) // token失效时间

	claims := Claims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  nowTime.Unix(),
			ExpiresAt: expireTime.Unix(),
			Issuer:    constants.JwtIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(setting.SecretKey))
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.SecretKey), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
