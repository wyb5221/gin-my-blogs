package post

import (
	"fmt"
	"gin-my-blogs/blog/interfaces/service/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) AutoTable() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		h.postService.AutoTable(ctx)
		ctx.JSON(http.StatusOK, gin.H{"mes": "SUCCESS"})
	}
}

func (h *handler) Create() func(ctx *gin.Context) {
	fmt.Printf("=== 请求到达Create Handler ===\n")
	return func(ctx *gin.Context) {
		req := &post.CreateRequest{}
		err1 := ctx.ShouldBind(req)
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
			return
		}
		id, err2 := h.postService.Create(*ctx, req)
		if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	}
}
