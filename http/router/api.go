package router

import (
	"github.com/jiangrx816/wechat/http/handler/api_handler"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.RouterGroup) {

	prefixRouter := r.Group("v1").Use()
	apiHandler := api_handler.NewApiHandler()
	{
		// 默认的api示例Demo
		prefixRouter.GET("/api/index", apiHandler.ApiIndex)
	}
}
