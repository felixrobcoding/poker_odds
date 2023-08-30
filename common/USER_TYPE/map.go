/*
功能：用户类型map
说明：
*/
package USER_TYPE

type Value struct {
	t       TYPE
	txt_eng string
	txt_chs string
}

var type_map map[TYPE]Value

func init() {
	values := []Value{
		{PLAYER, "player", "闲家"},
		{BANKER, "banker", "庄家"},
	}

	type_map = make(map[TYPE]Value, 0)
	for k, v := range values {
		type_map[v.t] = values[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
