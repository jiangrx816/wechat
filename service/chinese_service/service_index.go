package chinese_service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jiangrx816/wechat/common"
	"github.com/jiangrx816/wechat/common/request"
	"github.com/jiangrx816/wechat/common/response"
	"github.com/jiangrx816/wechat/core/server/api"
	"github.com/jiangrx816/wechat/utils"
	"github.com/jiangrx816/wechat/utils/errs"
)

/**
 * @Description 获取初始的栏目列表
 */
func (ps *ChineseService) ApiServiceGetCategoryList(ctx *gin.Context, typeId int) (resp response.ChineseBookNavNameResponse, apiErr api.Error) {
	utils.DefaultIntOne(&typeId)
	bookNameList, err := ps.ServiceDBFindCategoryList(typeId)
	if condition := err != nil; condition {
		apiErr = errs.NewError(err.Error())
		return
	}
	resp.List = bookNameList
	return
}

/**
 * @Description 获取中文绘本列表
 **/
func (ps *ChineseService) ApiServiceChineseBookList(ctx *gin.Context, page, level int) (resp response.ChineseBookResponse, apiErr api.Error) {
	utils.DefaultIntOne(&page)
	utils.DefaultIntOne(&level)
	size := common.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)

	// 从DB获取绘本列表数据
	total, bookList, err := ps.ServiceDBFindBookList(level, size, offset)
	if condition := err != nil; condition {
		apiErr = errs.NewError(err.Error())
		return
	}

	// 返回数据
	resp.Page = page
	resp.Total = total
	resp.List = bookList
	resp.TotalPage = int(math.Ceil(float64(total) / float64(size)))

	return
}

/**
 * @Description 获取中文绘本详情数据
 */
func (ps *ChineseService) ApiServiceChineseBookInfo(ctx *gin.Context, bookId string) (resp response.ChineseBookInfoResponse, apiErr api.Error) {

	// 从DB获取绘本详情数据
	bookInfo, err := ps.ServiceDBFindBookInfo(bookId)

	if condition := err != nil; condition {
		apiErr = errs.NewError(err.Error())
		return
	}

	// 返回数据
	resp.Info = bookInfo

	return
}

/**
 * @Description 获取中文绘本搜索数据
 */
func (ps *ChineseService) ApiServiceChineseBookSearch(ctx *gin.Context, page int, value string) (resp response.ChineseBookResponse, apiErr api.Error) {
	utils.DefaultIntOne(&page)
	size := 100
	offset := size * (page - 1)

	// 从DB获取绘本搜索数据
	total, bookList, err := ps.ServiceDBFindBookSearch(value, size, offset)
	if condition := err != nil; condition {
		apiErr = errs.NewError(err.Error())
		return
	}

	// 返回数据
	resp.Page = page
	resp.Total = total
	resp.List = bookList
	resp.TotalPage = int(math.Ceil(float64(total) / float64(size)))

	return
}

func (ps *ChineseService) ApiServiceEnglishBookList(ctx *gin.Context, typeId, offset int) (resp response.EnglishBookResponse, apiErr api.Error) {
	// 创建要发送的数据
	data := request.RequestEnglishParam{
		HTs:        "1685021025740e",
		HM:         28800,
		Zone:       0,
		HLc:        "zh",
		Uid:        0,
		Token:      "",
		HCn:        "miniprogram",
		HDt:        3,
		Cate:       1,
		Atype:      3,
		Source:     4,
		Offset:     int32(offset),
		Difficulty: int32(typeId),
	}

	// 将数据编码为 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	// 创建 POST 请求
	url := "https://www.ipalfish.com/klian/ugc/picturebook/level/list"
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

	var contentResult response.ResponseEnglishBook
	if err := json.Unmarshal(body, &contentResult); err != nil {
		return
	}

	resp.Data = contentResult.Data
	return
}

func (ps *ChineseService) ApiServiceEnglishBookInfo(ctx *gin.Context, bookId int) (resp response.EnglishBookInfoResponse, apiErr api.Error) {

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

	var contentResult response.ResponseEnglishBookInfo
	if err := json.Unmarshal(body, &contentResult); err != nil {
		return
	}

	resp.Data = contentResult.Data

	return
}
