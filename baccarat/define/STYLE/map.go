/*
功能：形态map
说明：
*/
package STYLE

type Value struct {
	t       TYPE
	txt_eng string
	txt_chs string
}

var type_map map[TYPE]Value

func init() {
	values := []Value{
		{FOLLOW, "Follow", "跟从形态"},
		{LONG, "Long", "龙形态"},
		{SINGLE_JUMP, "Single Jump", "单跳形态"},
		{DOUBLE_JUMP, "Double Jump", "双跳形态"},
	}

	type_map = make(map[TYPE]Value, 0)
	for k, v := range values {
		type_map[v.t] = values[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
