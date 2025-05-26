package wechat_handler

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jiangrx816/wechat/common"
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

// ApiPoetryBookList 古诗绘本-列表
func (ph *WechatHandler) ApiPoetryBookList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	typeId, _ := strconv.Atoi(ctx.Query("type_id"))
	response, err := ph.service.FindPoetryBookList(page, limit, typeId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiPoetryBookInfo 古诗绘本-详情
func (ph *WechatHandler) ApiPoetryBookInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("book_id"))
	response, err := ph.service.FindPoetryBookInfo(id)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}

// ApiWechatSignature 微信签名
func (ph *WechatHandler) ApiWechatSignature(ctx *gin.Context) {
	rawURL := ctx.Query("url")
	if rawURL == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "缺少url参数"})
		return
	}

	// url参数需要解码一次
	urlDecoded, err := url.QueryUnescape(rawURL)
	if err != nil {
		urlDecoded = rawURL // 失败则用原始
	}

	ticket, err := ph.service.ApiServiceGetJsapiTicket()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取jsapi_ticket失败：" + err.Error()})
		return
	}

	nonceStr := ph.service.ApiServiceGenerateNonceStr(16)
	timestamp := time.Now().Unix()
	signature := ph.service.ApiServiceSalcSignature(ticket, nonceStr, timestamp, urlDecoded)

	ctx.JSON(http.StatusOK, gin.H{
		"appId":     common.AppID,
		"nonceStr":  nonceStr,
		"timestamp": timestamp,
		"signature": signature,
	})
}
