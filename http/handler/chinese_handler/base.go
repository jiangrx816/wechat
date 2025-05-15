package chinese_handler

import (
	"github.com/jiangrx816/wechat/service/chinese_service"
)

func NewChineseHandler() *ChineseHandler {
	return &ChineseHandler{
		service: chinese_service.NewChineseService(),
	}
}

type ChineseHandler struct {
	service *chinese_service.ChineseService
}
