/*
功能：动作类型map
说明：
*/
package ACTION_TYPE

type Item struct {
	t         TYPE
	txt_short string
	txt_eng   string
	txt_chs   string
}

var (
	type_map = make(map[TYPE]Item, 0)
)

func init() {
	items := []Item{
		{HIT, "H", "hit", "要牌"},
		{DOUBLE_DOWN, "D", "double", "加倍停牌"},
		{STAND, "S", "stand", "停牌"},
		{SPLIT, "P", "split", "分牌"},
		{INSURANCE, "", "insurance", "保险"},
		{SURRENDER, "Rh", "surrender", "投降"},
	}
	for k, v := range items {
		type_map[v.t] = items[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
