package interfaces

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test1() {
	fmt.Println("---gin-my-blogs--blogs--interfaces---")
}

func Test1RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/test1", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "SUCCESS",
			"data":    "Hello word ! test1",
		})
	})

	r.GET("/test2", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "SUCCESS",
			"data":    "Hello word ! test2",
		})
	})
}
