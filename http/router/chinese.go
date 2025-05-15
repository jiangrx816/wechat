package router

import (
	"github.com/jiangrx816/wechat/http/handler/api_handler"

	"github.com/gin-gonic/gin"
)

func ChineseBookApi(r *gin.RouterGroup) {

	prefixRouter := r.Group("v1").Use()
	apiHandler := api_handler.NewApiHandler()
	{
		// 中文绘本列表
		prefixRouter.GET("/chinese/getList", apiHandler.ApiChineseBookList)
	}
}
