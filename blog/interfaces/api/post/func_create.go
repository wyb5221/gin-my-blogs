package post

import (
	"gin-my-blogs/blog/interfaces/service/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Create() func(ctx *gin.Context) {
	req := &post.CreateRequest{}

	return func(ctx *gin.Context) {
		err1 := ctx.ShouldBind(req)
		if err1 != nil {
			return
		}
		id, err2 := h.postService.Create(ctx, req)
		if err2 != nil {
			ctx.JSON(http.StatusOK, err2)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	}
}
