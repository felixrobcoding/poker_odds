/*
功能：形态map
说明：
*/
package STYLE

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
		{FOLLOW, "Follow", "跟从形态"},
		{LONG, "Long", "龙形态"},
		{MM_JUMP, "MM Jump", "MM跳形态"},
		{MN_JUMP, "MN Jump", "MN跳形态"},
	}
	for k, v := range items {
		type_map[v.t] = items[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
