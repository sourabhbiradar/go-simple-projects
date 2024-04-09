package middleware

import "github.com/gin-gonic/gin"

// request MW

// one way to write MW
func Authenticate(c *gin.Context) {
	if !(c.Request.Header.Get("token") == "auth") {
		c.AbortWithStatusJSON(500, gin.H{
			"msg": "Token not found",
		})
		return
	}
	c.Next()

}

// second way to write MW
// func Authenticate() gin.HandlerFunc {
// 	// real time use case
// 	return func(c *gin.Context) {
// 		if !(c.Request.Header.Get("token") == "auth") {
// 			c.AbortWithStatusJSON(500, gin.H{
// 				"msg": "Token not found",
// 			})
// 			return
// 		}
// 		c.Next()

// 	}
// }

// response MW
func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("key1", "val2")
	c.Next()
}
