package middleware

import "github.com/gin-gonic/gin"

type cors struct{}

func NewCorsMiddlewrerAccessControll() *cors {
	return &cors{}
}

func (corsAccessControll *cors) CorsMiddlewrerAccessControll() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin","ใส่ Domain ที่อนุญาต ตรงนี้")
		c.Writer.Header().Set("Access-Control-Allow-Credentials","true")
		c.Writer.Header().Set("Access-Control-Allow-Headers","Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods","POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		}

		c.Next()		
		
	}
}
