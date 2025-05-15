package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jiangrx816/wechat/http/handler/chinese_handler"
)

func ChineseBookApi(r *gin.RouterGroup) {

	prefixRouter := r.Group("v1").Use()
	apiHandler := chinese_handler.NewChineseHandler()
	{
		//绘本栏目
		prefixRouter.GET("/init/getCategoryList", apiHandler.ApiGetCategoryList)
		// 中文绘本列表
		prefixRouter.GET("/chinese/getList", apiHandler.ApiChineseBookList)
	}
}
