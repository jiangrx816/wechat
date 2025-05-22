package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jiangrx816/wechat/http/handler/chinese_handler"
)

func ChineseBookApi(r *gin.RouterGroup) {

	prefixRouter := r.Group("v1").Use()
	apiHandler := chinese_handler.NewChineseHandler()
	{
		// 绘本栏目
		prefixRouter.GET("/init/getCategoryList", apiHandler.ApiGetCategoryList)
		// 绘本列表
		prefixRouter.GET("/chinese/getBookList", apiHandler.ApiChineseBookList)
		// 绘本详情
		prefixRouter.GET("/chinese/getBookInfo", apiHandler.ApiChineseBookInfo)
		// 绘本搜索
		prefixRouter.GET("/chinese/getBookSearch", apiHandler.ApiChineseBookSearch)
		// 英文绘本列表
		prefixRouter.GET("/english/getBookList", apiHandler.ApiEnglishBookList)
		// 英文绘本详情
		prefixRouter.GET("/english/getBookInfo", apiHandler.ApiEnglishBookInfo)
	}
}
