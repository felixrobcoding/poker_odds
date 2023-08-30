/*
功能：花色map
说明：
*/
package SUIT

type Value struct {
	t       TYPE
	txt_eng string
	txt_chs string
}

var type_map map[TYPE]Value

func init() {
	values := []Value{
		{DIAMOND, "Diamond", "方块"},
		{CLUB, "Club", "梅花"},
		{HEART, "Heart", "红心"},
		{SPADE, "Spade", "黑桃"},
	}

	type_map = make(map[TYPE]Value, 0)
	for k, v := range values {
		type_map[v.t] = values[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
