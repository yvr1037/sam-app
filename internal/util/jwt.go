package util

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type jwtInfo struct {
	UserId string
	Role   int8
	jwt.StandardClaims
}

func CreateJwt(uid string, role int8) (string, error) {
	temp := jwtInfo{
		UserId: uid,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(time.Hour)*24*1,
			Issuer:    "dian",
		},
	}
	tokenCla := jwt.NewWithClaims(jwt.SigningMethodHS256, temp)
	if token, err := tokenCla.SignedString([]byte("temp")); err == nil {
		token = "Bearer " + token
		return token, nil
	} else {
		return "", err
	}
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "auth no exist",
			})
			return
		}
		arr := strings.Fields(auth)
		if len(arr) < 2 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "token wrong",
			})
			return
		}
		auth = arr[1]
		token, err := jwt.ParseWithClaims(auth, &jwtInfo{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("temp"), nil
		})
		if err != nil {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			return
		}
		context.Set("uid", token.Claims.(*jwtInfo).UserId)
		context.Set("role", token.Claims.(*jwtInfo).Role)
		context.Next()
	}
}
