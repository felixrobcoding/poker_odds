/*
功能：hand type map
说明：
*/
package HAND_TYPE

type Value struct {
	t       TYPE
	txt_eng string
	txt_chs string
}

var type_map map[TYPE]Value

func init() {
	values := []Value{
		{HARD, "Hard", "不包含A"},
		{SOFT, "Soft", "包含A"},
		{SPLITS, "Splits", "对子,可分牌"},
	}

	type_map = make(map[TYPE]Value, 0)
	for k, v := range values {
		type_map[v.t] = values[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
