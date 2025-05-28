package wechat_service

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jiangrx816/wechat/common"
	"github.com/jiangrx816/wechat/common/response"
	"github.com/jiangrx816/wechat/core/server/api"
)

/**
 * ApiServiceMathCalculationList
 * @Description 获取计算题列表
 */
func (ps *WechatService) ApiServiceMathCalculationList(ctx *gin.Context, forward, value, limit int) (resp []response.MathResponse, apiErr api.Error) {
	var initData [][2]int
	switch forward {
	case common.Math_Add_Non_Carry:
		initData = ps.generateStrictNonCarryAdditions(limit, value)
	case common.Math_Add_Carry:
		initData = ps.generateCarryAdditionProblems(limit, value)
	case common.Math_Sub_Non_Back:
		initData = ps.generateNonBorrowSubtraction(limit, value)
	case common.Math_Sub_Back:
		initData = ps.generateBorrowSubtraction(limit, value)
	}

	resp = ps.FormatMathCalculationData(initData, forward)

	return
}

/**
 * FormatMathCalculationData
 * @Description 格式化计算题数据
 * @param initData 初始数据
 * @param forward 方式
 * @return resp 格式化后的数据
 * @return apiErr 错误
 */
func (ps *WechatService) FormatMathCalculationData(initData [][2]int, forward int) (resp []response.MathResponse) {
	var forwardSymbol string
	switch forward {
	case common.Math_Add_Carry, common.Math_Add_Non_Carry:
		forwardSymbol = common.Math_Add_Symbol_Show
	case common.Math_Sub_Back, common.Math_Sub_Non_Back:
		forwardSymbol = common.Math_Sub_Symbol_Show
	}

	for _, v := range initData {
		var forwardResult int
		if forward == common.Math_Sub_Back || forward == common.Math_Sub_Non_Back {
			forwardResult = v[0] - v[1]
		} else {
			forwardResult = v[0] + v[1]
		}
		resp = append(resp, response.MathResponse{
			NumberOne: v[0],
			NumberTwo: v[1],
			Symbol:    forwardSymbol,
			Result:    forwardResult,
		})
	}
	return
}

/**
 * generateStrictNonCarryAdditions
 * @Description 不进位加法——生成严格不进位加法题目
 * @param count 题目数量
 * @param max 最大值
 * @return [][2]int 题目列表
 */
func (ps *WechatService) generateStrictNonCarryAdditions(count int, max int) [][2]int {
	rand.Seed(time.Now().UnixNano())
	problems := make([][2]int, count)

	for i := 0; i < count; i++ {
		for {
			a := rand.Intn(max) + 1
			b := rand.Intn(max) + 1

			// 根据max值应用不同规则
			if max <= 20 {
				// 20以内：允许个位数，和≤20，不进位
				if a+b <= 20 && (a%10)+(b%10) < 10 {
					problems[i] = [2]int{a, b}
					break
				}
			} else {
				// 超过20：不允许个位数(≥10)，和≤max，不进位
				if a >= 10 && b >= 10 &&
					a+b <= max &&
					(a%10)+(b%10) < 10 {
					problems[i] = [2]int{a, b}
					break
				}
			}
		}
	}
	return problems
}

/**
 * generateCarryAdditionProblems
 * @Description 进位加法-生成进位加法题目（严格限制和的范围）
 * @param count 题目数量
 * @param max 最大值
 * @return [][2]int 题目列表
 */
func (ps *WechatService) generateCarryAdditionProblems(count int, max int) [][2]int {
	rand.Seed(time.Now().UnixNano())
	problems := make([][2]int, count)

	for i := 0; i < count; i++ {
		for {
			var a, b int

			if max <= 20 {
				// 20以内：允许个位数
				a = rand.Intn(max-1) + 1 // 1-19
				b = rand.Intn(max-1) + 1 // 1-19
			} else {
				// 超过20：不允许个位数（必须≥10）
				a = rand.Intn(max-10) + 10 // 10-(max-1)
				b = rand.Intn(max-10) + 10 // 10-(max-1)
			}

			// 共同条件：个位相加≥10（需要进位）
			if (a%10)+(b%10) >= 10 {
				// 20以内额外条件和≤20
				if max <= 20 && a+b <= 20 {
					problems[i] = [2]int{a, b}
					break
				}
				// 超过20时，和不超过max
				if max > 20 && a+b <= max {
					problems[i] = [2]int{a, b}
					break
				}
			}
		}
	}
	return problems
}

/**
 * generateNonBorrowSubtraction
 * @Description 不退位减法——生成不退位减法题目
 * @param count 题目数量
 * @param max 最大值
 * @return [][2]int 题目列表
 */
func (ps *WechatService) generateNonBorrowSubtraction(count int, max int) [][2]int {
	rand.Seed(time.Now().UnixNano())
	problems := make([][2]int, count)

	for i := 0; i < count; i++ {
		for {
			var a, b int

			if max <= 20 {
				// 20以内：允许个位数
				a = rand.Intn(max) + 1 // 1-20
				b = rand.Intn(max) + 1 // 1-20
			} else {
				// 超过20：不允许个位数（必须≥10）
				a = rand.Intn(max-10) + 10 // 10-max
				b = rand.Intn(max-10) + 10 // 10-max
			}

			// 共同条件：被减数≥减数且个位≥减数个位
			if a >= b && (a%10) >= (b%10) {
				// 20以内额外条件和≤20
				if max <= 20 && a <= 20 && b <= 20 {
					problems[i] = [2]int{a, b}
					break
				}
				// 超过20时，和不超过max
				if max > 20 && a <= max && b <= max {
					problems[i] = [2]int{a, b}
					break
				}
			}
		}
	}
	return problems
}

/**
 * generateBorrowSubtraction
 * @Description 退位减法——生成退位减法题目
 * @param count 题目数量
 * @param max 最大值
 * @return [][2]int 题目列表
 */
func (ps *WechatService) generateBorrowSubtraction(count int, max int) [][2]int {
	rand.Seed(time.Now().UnixNano())
	problems := make([][2]int, count)

	for i := 0; i < count; i++ {
		for {
			var a, b int

			if max <= 20 {
				// 20以内：允许个位数
				a = rand.Intn(max) + 1 // 1-20
				b = rand.Intn(max) + 1 // 1-20
			} else {
				// 超过20：不允许个位数（必须≥10）
				a = rand.Intn(max-10) + 10 // 10-max
				b = rand.Intn(max-10) + 10 // 10-max
			}

			// 共同条件：被减数>减数且个位<减数个位（需要退位）
			if a > b && (a%10) < (b%10) {
				// 20以内额外条件和≤20
				if max <= 20 && a <= 20 && b <= 20 {
					problems[i] = [2]int{a, b}
					break
				}
				// 超过20时，和不超过max
				if max > 20 && a <= max && b <= max {
					problems[i] = [2]int{a, b}
					break
				}
			}
		}
	}
	return problems
}
