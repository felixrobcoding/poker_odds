/*
功能：压注区域map
说明：
*/
package BET_AREA

type Value struct {
	t       TYPE
	odds    float64
	txt_eng string
	txt_chs string
}

var type_map map[TYPE]Value

func init() {
	values := []Value{
		{BANKER, 0.95, "Banker", "庄"},
		{PLAYER, 1, "Player", "闲"},
		{TIE, 8, "Tie", "和"},
		{BANKER_PAIR, 11, "Banker Pair", "庄对"},
		{PLAYER_PAIR, 11, "Player Pari", "闲对"},
	}

	type_map = make(map[TYPE]Value, 0)
	for k, v := range values {
		type_map[v.t] = values[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
