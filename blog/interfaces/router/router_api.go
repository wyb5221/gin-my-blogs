package router

import (
	"gin-my-blogs/blog/interfaces/api/comment"
	"gin-my-blogs/blog/interfaces/api/post"
	"gin-my-blogs/blog/interfaces/api/user"
	"gin-my-blogs/blog/interfaces/base"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterRoutes(r *gin.Engine) {
	db := base.InitDb()
	var logger *zap.Logger
	userHandler := user.New(logger, db)
	postHandler := post.New(logger, db)
	commentHandler := comment.New(logger, db)

	//创建分组
	g := r.Group("/blogs")
	{
		g.GET("/user/autoTable", userHandler.AutoTable())
		g.GET("/post/autoTable", postHandler.AutoTable())
		g.GET("/comment/autoTable", commentHandler.AutoTable())
	}
	//创建分组
	gp := r.Group("/blogs")
	{
		gp.POST("/user/add", userHandler.Create())
		gp.GET("/user/delete/:id", userHandler.Delete())
		gp.GET("/user/detail/:id", userHandler.Detail())
		gp.POST("/user/list", userHandler.List())
		gp.POST("/user/resetPassword", userHandler.ResetPassword())

		gp.POST("/post/add", postHandler.Create())
		gp.GET("/post/delete/:id", postHandler.Delete())
		gp.GET("/post/detail/:id", postHandler.Detail())
		gp.POST("/post/list", postHandler.List())
		gp.POST("/post/updates", postHandler.Updates())

		gp.POST("/comment/add", commentHandler.Create())
		gp.GET("/comment/detail/:id", commentHandler.Detail())
		gp.GET("/comment/delete/:id", commentHandler.Delete())
		gp.POST("/comment/list", commentHandler.List())
		gp.POST("/comment/updates", commentHandler.Updates())

	}

}

func Test1RegisterRoutes(r *gin.Engine) {
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
