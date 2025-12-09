package comment

import (
	"gin-my-blogs/blog/interfaces/service/comment"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) Detail() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		// 转换为 uint
		id, err := strconv.ParseUint(idStr, 10, 32) // 10进制，32位
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID参数格式错误"})
			return
		}
		user, err := h.commentService.DetailById(ctx, uint(id))
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (h *handler) List() func(ctx *gin.Context) {
	req := &comment.ListRequest{}

	return func(ctx *gin.Context) {
		err1 := ctx.ShouldBind(req)
		if err1 != nil {
			return
		}

		users, err := h.commentService.List(ctx, req)
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, users)
	}
}
