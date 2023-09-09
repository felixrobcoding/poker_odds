/*
功能：游戏结果
说明：
*/
package GAME_RESULT

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
		{LOSE, "Lose", "输"},
		{PUSH, "Push", "和"},
		{WIN, "Win", "赢"},
	}
	for k, v := range items {
		type_map[v.t] = items[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
