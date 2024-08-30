/*
功能：大路-画
说明：
*/
package big_road

const (
	BANKER_NODE string = "🔴"  //庄
	PLAYER_NODE string = "🔵"  //闲
	TIE_NODE    string = "⚫️" //和
	NIL_NODE    string = "⚪️" //空
)

/*
// 转描述字符串
func nodes_2_string1(nodes []*suggestion.FeedbackNode) string {
	txt := ""
	for _, v := range nodes {
		if v.Result_area == BET_AREA.BANKER {
			txt += BANKER_NODE
		} else if v.Result_area == BET_AREA.PLAYER {
			txt += PLAYER_NODE
		} else {
			panic(xdebug.Funcname())
		}
	}
	return txt
}
*/

///-----------------------------------------------
//					the end
//-----------------------------------------------
