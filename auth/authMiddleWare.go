package auth

import (
	"blog/service"
	"blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var ignorePaths = []string{"/user/register", "/user/login"}

// jwt 认证中间件
func AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		valid := validContains(c.FullPath())
		if !valid {
			// 解析jwt
			tokenStr := c.GetHeader("Authorization")
			claims, err := util.ParseJWT(tokenStr)
			if err != nil {
				util.JSON(c, http.StatusUnauthorized, util.TokenFail)
				c.Abort()
				return
			}
			user := service.GetUserById((*claims).UserID)
			if user.Password != "" && (*claims).Password != "" && strings.Compare(user.Password, (*claims).Password) != 0 {
				util.JSON(c, http.StatusUnauthorized, util.UserInfoError)
				c.Abort()
				return
			}
			c.Set("user", user)
		}
		c.Next()
	}
}

func validContains(reqPaht string) bool {
	for _, path := range ignorePaths {
		if strings.Contains(reqPaht, path) {
			return true
		}
	}
	return false
}
