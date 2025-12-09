package user

import (
	"fmt"
	"gin-my-blogs/blog/interfaces/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) ResetPassword() func(ctx *gin.Context) {
	req := &user.UpdatesRequest{}
	fmt.Println("---1---")
	return func(ctx *gin.Context) {
		err1 := ctx.ShouldBind(req)
		if err1 != nil {
			return
		}
		err2 := h.userService.UpdatePassword(ctx, req)

		ctx.JSON(http.StatusOK, err2)
	}
}
