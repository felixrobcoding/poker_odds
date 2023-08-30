/*
功能：下注类型map
说明：
*/
package BETTING_TYPE

type Value struct {
	t       TYPE
	txt_eng string
	txt_chs string
}

var type_map map[TYPE]Value

func init() {
	values := []Value{
		{ALL_IN, "all_in", "全下"},
		{FIXED_AMOUNT, "fixed_amount", "固定额度"},
		{MARTEGAL, "martegal", "马丁格尔策略:输了加倍"},
		{FIBONACCI, "fibonacci", "斐波那契策略:累加"},
		{KELLY, "kelly", "凯利策略:每次投注总筹码的某一百分比"},
	}

	type_map = make(map[TYPE]Value, 0)
	for k, v := range values {
		type_map[v.t] = values[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
