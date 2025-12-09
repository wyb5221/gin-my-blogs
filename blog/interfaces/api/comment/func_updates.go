package comment

import (
	"gin-my-blogs/blog/interfaces/service/comment"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Updates() func(ctx *gin.Context) {
	req := &comment.UpdatesRequest{}
	return func(ctx *gin.Context) {
		err1 := ctx.ShouldBind(req)
		if err1 != nil {
			return
		}
		err2 := h.commentService.Updates(ctx, req)

		ctx.JSON(http.StatusOK, err2)
	}
}
