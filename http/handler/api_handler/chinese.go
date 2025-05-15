package api_handler

import (
	"github.com/jiangrx816/wechat/utils/errs"

	"github.com/gin-gonic/gin"
)

/**
 * @Description 获取中文绘本列表
 **/
func (ph *ApiHandler) ApiChineseBookList(ctx *gin.Context) {

	ctx.JSON(errs.SucResp(""))
}
