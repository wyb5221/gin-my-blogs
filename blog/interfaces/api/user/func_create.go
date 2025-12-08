package user

import (
	"gin-my-blogs/blog/interfaces/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Create() func(ctx *gin.Context) {
	req := &user.CreateRequest{}

	return func(ctx *gin.Context) {
		err1 := ctx.ShouldBind(req)
		if err1 != nil {
			return
		}
		id, err2 := h.userService.Create(ctx, req)
		if err2 != nil {
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	}
}
