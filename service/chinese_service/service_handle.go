package chinese_service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jiangrx816/wechat/common/response"
	"github.com/jiangrx816/wechat/core/server/api"
)

/**
 * @Description 处理数据
 */
func (ps *ChineseService) ApiServiceEnglishHandleData(ctx *gin.Context, filePath string) (resp string, apiErr api.Error) {

	resp = "start english handle data"
	// 异步处理数据
	go func() {
		ps.HandleEnglishData(ctx, filePath)
	}()

	return
}

func (ps *ChineseService) HandleEnglishData(ctx *gin.Context, filePath string) (err error) {
	// 1. 读取 JSON 文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("读取文件失败: %v", err))
	}

	// 2. 解析 JSON 到结构体
	var handleData response.BYJSONData
	err = json.Unmarshal(data, &handleData)
	if err != nil {
		panic(fmt.Sprintf("解析 JSON 失败: %v", err))
	}

	return
}
