/*
功能：hand type map
说明：
*/
package HAND_TYPE

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
		{HARD, "Hard", "不包含A"},
		{SOFT, "Soft", "包含A"},
		{SPLITS, "Splits", "对子,可分牌"},
	}
	for k, v := range items {
		type_map[v.t] = items[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
