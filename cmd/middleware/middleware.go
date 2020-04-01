package middleware

import (
	"github.com/gin-gonic/gin"
)

func AttachInterfaceToGinContext(i interface{}, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, i)
	}
}
func GetInterfaceFromGinContext(c *gin.Context, Interface string) interface{} {
	i, _ := c.Get(Interface)

	return i
}
