package middleware

import "github.com/gin-gonic/gin"

type HeaderParams struct {
	Authorization string `header:"Authorization" binding:"required,min=20"`
}

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerParams := HeaderParams{}

		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := c.ShouldBindHeader(&headerParams); err != nil {

			return
		}
		c.Next()

	}

}
