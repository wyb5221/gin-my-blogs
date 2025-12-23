package user

import (
	"gin-my-blogs/blog/interfaces/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) ResetPassword() func(ctx *gin.Context) {
	req := &user.UpdatesRequest{}
	return func(ctx *gin.Context) {
		err1 := ctx.ShouldBind(req)
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
			return
		}
		err2 := h.userService.UpdatePassword(*ctx, req)
		if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "密码重置成功"})
	}
}
