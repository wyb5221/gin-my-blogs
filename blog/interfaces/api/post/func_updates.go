package post

import (
	"fmt"
	"gin-my-blogs/blog/interfaces/service/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Updates() func(ctx *gin.Context) {
	req := &post.UpdatesRequest{}
	fmt.Println("---1---")
	return func(ctx *gin.Context) {
		err1 := ctx.ShouldBind(req)
		if err1 != nil {
			return
		}
		err2 := h.postService.Updates(ctx, req)

		ctx.JSON(http.StatusOK, err2)
	}
}
