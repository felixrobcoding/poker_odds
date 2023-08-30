/*
功能：动作类型map
说明：
*/
package ACTION_TYPE

type Value struct {
	t         TYPE
	txt_short string
	txt_eng   string
	txt_chs   string
}

var type_map map[TYPE]Value

func init() {
	values := []Value{
		{HIT, "H", "hit", "要牌"},
		{DOUBLE_DOWN, "D", "double", "加倍停牌"},
		{STAND, "S", "stand", "停牌"},
		{SPLIT, "P", "split", "分牌"},
		{INSURANCE, "", "insurance", "保险"},
		{SURRENDER, "Rh", "surrender", "投降"},
	}

	type_map = make(map[TYPE]Value, 0)
	for k, v := range values {
		type_map[v.t] = values[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
