package chinese_service

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jiangrx816/wechat/common/response"
	"github.com/jiangrx816/wechat/core/server/api"
	"github.com/jiangrx816/wechat/model"
)

/**
 * @Description 处理数据
 */
func (ps *ChineseService) ApiServiceEnglishHandleData(ctx *gin.Context, level int, filePath string) (resp string, apiErr api.Error) {

	resp = "start english handle data"
	// 异步处理数据
	go func() {
		ps.HandleEnglishData(ctx, level, filePath)
	}()

	return
}

func (ps *ChineseService) HandleEnglishData(ctx *gin.Context, level int, filePath string) (err error) {
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

	// 3. 处理数据
	var englishPictureTempList []model.SEnglishPicture
	itemList := handleData.Data.Ent.Items
	for _, item := range itemList {
		// 处理每一项数据
		englishPictureTempList = append(englishPictureTempList, model.SEnglishPicture{
			BookId: strconv.FormatInt(item.Bookid, 10),
			Title:  item.Title,
			Icon:   item.Cover.Tiny,
			Type:   level,
			Status: 1,
		})
	}
	err = model.Default().Model(&model.SEnglishPicture{}).Debug().Create(&englishPictureTempList).Error
	if condition := err != nil; condition {
		panic(fmt.Sprintf("入库 失败: %v", err))
	}
	return
}
