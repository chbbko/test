package middlewares

import "github.com/gin-gonic/gin"

func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
	ctx.Header("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT, PATCH")
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(200)
		return
	}

	ctx.Next()
}
