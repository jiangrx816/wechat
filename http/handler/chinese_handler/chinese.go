package chinese_handler

import (
	"strconv"

	"github.com/jiangrx816/wechat/common/request"
	"github.com/jiangrx816/wechat/utils/errs"

	"github.com/gin-gonic/gin"
)

/**
 * @Description 获取初始的栏目列表
 **/
func (ph *WechatHandler) ApiGetCategoryList(ctx *gin.Context) {
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
func (ph *WechatHandler) ApiChineseBookList(ctx *gin.Context) {
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
func (ph *WechatHandler) ApiChineseBookInfo(ctx *gin.Context) {
	bookId := ctx.Query("book_id")
	response, err := ph.service.ApiServiceChineseBookInfo(ctx, bookId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

/**
 * @Description 获取中文绘本搜索列表
 **/
func (ph *WechatHandler) ApiChineseBookSearch(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	value := ctx.Query("value")
	response, err := ph.service.ApiServiceChineseBookSearch(ctx, page, value)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

/**
 * @Description 获取中文绘本列表
 **/
func (ph *WechatHandler) ApiEnglishBookList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	level, _ := strconv.Atoi(ctx.Query("level"))

	response, err := ph.service.ApiServiceEnglishBookList(ctx, page, level)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

/**
 * @Description 获取英文绘本详情
 **/
func (ph *WechatHandler) ApiEnglishBookInfo(ctx *gin.Context) {
	bookId := ctx.Query("book_id")
	response, err := ph.service.ApiServiceEnglishBookInfo(ctx, bookId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

/**
 * @Description 获取英文绘本搜索列表
 **/
func (ph *WechatHandler) ApiEnglishBookSearch(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	value := ctx.Query("value")
	response, err := ph.service.ApiServiceEnglishBookSearch(ctx, page, value)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

/**
 * @Description 处理数据
 **/
func (ph *WechatHandler) ApiEnglishHandleData(ctx *gin.Context) {
	var json request.EnglishHandleDataRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(errs.SucErrResp("", "参数错误"))
		return
	}
	response, err := ph.service.ApiServiceEnglishHandleData(ctx, json.Level, json.FilePath)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}
