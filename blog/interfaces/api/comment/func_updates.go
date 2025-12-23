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
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
			return
		}
		err2 := h.commentService.Updates(*ctx, req)
		if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "更新成功"})
	}
}

func (h *handler) UpdateStatus() func(ctx *gin.Context) {
	req := &comment.UpdatesRequest{}
	return func(ctx *gin.Context) {
		err1 := ctx.ShouldBind(req)
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
			return
		}
		err2 := h.commentService.UpdateStatus(*ctx, req)
		if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "更新成功"})
	}
}
