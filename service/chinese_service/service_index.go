package chinese_service

import (
	"math"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/jiangrx816/wechat/common"
	"github.com/jiangrx816/wechat/common/response"
	"github.com/jiangrx816/wechat/core/server/api"
	"github.com/jiangrx816/wechat/model"
	"github.com/jiangrx816/wechat/utils"
)

/**
 * @Description 获取初始的栏目列表
 */
func (ps *ChineseService) ApiServiceGetCategoryList(ctx *gin.Context, typeId int) (resp response.ChineseBookNavNameResponse, apiErr api.Error) {
	utils.DefaultIntOne(&typeId)
	db := model.Default().Model(&model.SBookName{}).Debug()
	db = db.Where("status = 1 and s_type = ?", typeId)
	db = db.Order("s_sort asc").Order("id asc")
	db.Find(&resp.List)
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

	var bookList []model.SChinesePicture
	db := model.Default().Model(&model.SChinesePicture{}).Debug()
	db = db.Where("type = ? and status = 1", level).Count(&resp.Total)
	db = db.Order("position desc").Limit(size).Offset(offset)
	db.Find(&bookList)

	var bookInfoCountList []response.ResponseBookInfoCount
	db1 := model.Default().Model(&model.SChinesePictureInfo{}).Debug()
	db1.Raw("SELECT book_id,count(id) as book_count FROM s_chinese_picture_info where status = 1 GROUP BY book_id").Scan(&bookInfoCountList)

	var temp response.ResponseChineseBook
	for _, item := range bookList {
		temp.Id = item.Id
		temp.BookId = item.BookId
		temp.Title = item.Title
		temp.Icon = item.Icon
		temp.Level = item.Type
		temp.Position = item.Position
		resp.List = append(resp.List, temp)
	}
	for index, item := range resp.List {
		for _, it := range bookInfoCountList {
			if item.BookId == it.BookId {
				resp.List[index].BookCount = it.BookCount
			}
		}
	}
	sort.Slice(resp.List, func(i, j int) bool {
		if resp.List[i].Position > resp.List[j].Position {
			return true
		}
		return resp.List[i].Position == resp.List[j].Position && resp.List[i].Id < resp.List[j].Id
	})
	resp.Page = page
	resp.TotalPage = int(math.Ceil(float64(resp.Total) / float64(size)))
	return
}
