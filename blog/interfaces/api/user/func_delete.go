package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) Delete() func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		// 转换为 uint
		id, err := strconv.ParseUint(idStr, 10, 32) // 10进制，32位
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID参数格式错误"})
			return
		}
		err = h.userService.Delete(ctx, uint(id))
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, err)
	}
}
