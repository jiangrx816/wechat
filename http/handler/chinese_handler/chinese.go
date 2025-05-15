package chinese_handler

import (
	"strconv"

	"github.com/jiangrx816/wechat/utils/errs"

	"github.com/gin-gonic/gin"
)

/**
 * @Description 获取初始的栏目列表
 **/
func (ph *ChineseHandler) ApiGetCategoryList(ctx *gin.Context) {
	typeId, _ := strconv.Atoi(ctx.Query("type"))
	response, err := ph.service.ApiServiceGetCategoryList(ctx, typeId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

/**
 * @Description 获取中文绘本列表
 **/
func (ph *ChineseHandler) ApiChineseBookList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	level, _ := strconv.Atoi(ctx.Query("level"))

	response, err := ph.service.ApiServiceChineseBookList(ctx, page, level)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}
