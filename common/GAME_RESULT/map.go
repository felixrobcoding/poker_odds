/*
功能：游戏结果
说明：
*/
package GAME_RESULT

type Value struct {
	t       TYPE
	txt_eng string
	txt_chs string
}

var type_map map[TYPE]Value

func init() {
	values := []Value{
		{LOSE, "Lose", "输"},
		{PUSH, "Push", "和"},
		{WIN, "Win", "赢"},
	}

	type_map = make(map[TYPE]Value, 0)
	for k, v := range values {
		type_map[v.t] = values[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
