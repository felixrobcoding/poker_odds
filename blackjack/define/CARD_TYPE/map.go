/*
功能：牌型map
说明：
*/
package CARD_TYPE

type Item struct {
	t       TYPE
	odds    float64 //赔率
	txt_eng string
	txt_chs string
}

var (
	type_map = make(map[TYPE]Item, 0)
)

func init() {
	items := []Item{
		{POINT, 1, "Point", "点牌[计算点数]"},
		{SURRENDER, 0.5, "Surrender", "投降"},
		{BUST, 0, "Bust", "爆点"},
		{BLACK_JACK, 1.5, "Blackjack", "Blackjack"},
	}
	for k, v := range items {
		type_map[v.t] = items[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
