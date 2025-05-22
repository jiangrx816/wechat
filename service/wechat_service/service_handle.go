package wechat_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jiangrx816/wechat/common/request"
	"github.com/jiangrx816/wechat/common/response"
	"github.com/jiangrx816/wechat/core/server/api"
	"github.com/jiangrx816/wechat/model"
)

/**
 * @Description 处理数据
 */
func (ps *WechatService) ApiServiceEnglishHandleData(ctx *gin.Context, level int, filePath string) (resp string, apiErr api.Error) {

	resp = "start english handle data"
	// 异步处理数据
	go func() {
		ps.HandleEnglishDataInfo(ctx, level, filePath)
	}()

	return
}

func (ps *WechatService) HandleEnglishDataInfo(ctx *gin.Context, level int, filePath string) (err error) {
	var bookInitList []model.SEnglishPicture
	db := model.Default().Model(&model.SEnglishPicture{}).Debug()
	db = db.Where("id >0")
	db = db.Order("id asc")
	db.Find(&bookInitList)

	var bookIdList []string
	for _, item := range bookInitList {
		bookIdList = append(bookIdList, item.BookId)
	}

	for _, bookId := range bookIdList {
		bookIdInt, _ := strconv.Atoi(bookId)
		ps.ApiServiceEnglishBookSearch(bookIdInt)
	}
	return
}

func (ps *WechatService) ApiServiceEnglishBookSearch(bookId int) {
	// 创建要发送的数据
	data := request.RequestEnglishInfoParam{
		HDt:    3,
		HDid:   "16851689799320000",
		Did:    "16851689799320000",
		HCH:    "miniprogram",
		HTs:    "1685168991352",
		HM:     0,
		Zone:   28800,
		Token:  "",
		HLc:    "zh",
		Cate:   1,
		Atype:  3,
		BookId: bookId,
		Limit:  200,
	}

	// 将数据编码为 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	// 创建 POST 请求
	url := "https://www.ipalfish.com/klian/ugc/picturebook/official/product/bookid/page/list"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	respH, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer respH.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(respH.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}
	jsonFilePath := fmt.Sprintf("/Users/jiang/小程序/info/%d.json", bookId)

	// 将 拓扑数据类型数据的 JSON 内容写入文件
	if err = os.WriteFile(jsonFilePath, []byte(body), 0644); err != nil {
		log.Fatalf("数据写入文件失败: %v", err)
	}
	fmt.Printf("数据写入文件成功: %v \n", jsonFilePath)
}

func (ps *WechatService) HandleEnglishData(ctx *gin.Context, level int, filePath string) (err error) {
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
