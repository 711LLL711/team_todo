package middleware

import (
	"net/http"
	"team_todo/util"

	"github.com/gin-gonic/gin"
)

// 设置中间件,验证鉴权--判断是否登录
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头部获取 Bearer JWT
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		// 解析 JWT 并验证
		tokenString := authHeader[len("Bearer "):]
		claims, err := util.CheckToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		//将claims存入context
		c.Set("claims", claims)
		c.Next()
	}
}
