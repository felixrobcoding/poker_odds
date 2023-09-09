/*
功能：花色map
说明：
*/
package SUIT

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
		{DIAMOND, "Diamond", "方块"},
		{CLUB, "Club", "梅花"},
		{HEART, "Heart", "红心"},
		{SPADE, "Spade", "黑桃"},
	}
	for k, v := range items {
		type_map[v.t] = items[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
