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
		{SINGLE_JUMP, "Single Jump", "单跳形态"},
		{DOUBLE_JUMP, "Double Jump", "双跳形态"},
	}
	for k, v := range items {
		type_map[v.t] = items[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
