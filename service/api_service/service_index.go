package api_service

import (
	"github.com/jiangrx816/wechat/common"
	"github.com/jiangrx816/wechat/common/response"
	"github.com/jiangrx816/wechat/core/server/api"

	"github.com/gin-gonic/gin"
)

func (ps *ApiService) ApiServiceIndex(ctx *gin.Context, page int, size int) (resp response.IndexResponse, apiErr api.Error) {
	if size < 0 {
		size = common.DEFAULT_PAGE_SIZE
	}

	return
}
