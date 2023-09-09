/*
功能：牌型map
说明：
*/
package CARD_TYPE

type Item struct {
	t       TYPE
	txt_eng string
	txt_chs string
}

var (
	type_map = make(map[TYPE]Item, 0)
)

func init() {
	items := []Item{
		{PAIR, "Pair", "对子"},
		{NATURAL, "Natural", "天牌"},
	}
	for k, v := range items {
		type_map[v.t] = items[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
