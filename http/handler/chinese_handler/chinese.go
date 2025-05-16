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

/**
 * @Description 获取中文绘本详情
 **/
func (ph *ChineseHandler) ApiChineseBookInfo(ctx *gin.Context) {
	bookId := ctx.Query("book_id")
	response, err := ph.service.ApiServiceChineseBookInfo(bookId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

/**
 * @Description 获取中文绘本搜索列表
 **/
func (ph *ChineseHandler) ApiChineseBookSearch(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	value := ctx.Query("value")
	response, err := ph.service.ApiServiceChineseBookSearch(page, value)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}
