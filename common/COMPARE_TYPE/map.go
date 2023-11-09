/*
功能：比较类型map
说明：
*/
package COMPARE_TYPE

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
		{SMALLER, "smaller", "更小"},
		{EQUAL, "equal", "相等"},
		{BIGGER, "bigger", "更大"},
	}
	for k, v := range items {
		type_map[v.t] = items[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
