package common

const DEFAULT_PAGE_SIZE = 12

const (
	Math_Add_Symbol_Show = "➕"
	Math_Sub_Symbol_Show = "➖"
	Math_Mul_Symbol_Show = "✖️"
	Math_Div_Symbol_Show = "➗"
)

const (
	Math_Add_Non_Carry = 1 // 加法不进位
	Math_Add_Carry     = 2 // 加法进位
	Math_Sub_Non_Back  = 3 // 不退位减法
	Math_Sub_Back      = 4 // 退位减法
)
