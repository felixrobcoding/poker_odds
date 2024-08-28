/*
功能：形态-
说明：
*/
package style

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/strategy_bet_area/suggestion"
	"fmt"

	"github.com/poker-x-studio/x/xdebug"
)

const (
	BANKER_NODE string = "🔴"  //庄
	PLAYER_NODE string = "🔵"  //闲
	TIE_NODE    string = "⚫️" //和
	NIL_NODE    string = "⚪️" //空
)

// 查找
func find(nodes []*suggestion.ResultNode, start_index int, target_bet_area BET_AREA.TYPE) (target_cnt int, err error) {
	target_cnt = 0
	err = fmt.Errorf("")

	len := len(nodes)
	if len <= start_index {
		return
	}

	for i := start_index; i < len; i++ {
		if nodes[i].Current_win_bet_area == target_bet_area {
			target_cnt++
			continue
		}
		break
	}
	err = nil
	return
}

// 转描述字符串
func nodes_2_string(nodes []*suggestion.ResultNode) string {
	txt := ""
	for _, v := range nodes {
		if v.Current_win_bet_area == BET_AREA.BANKER {
			txt += BANKER_NODE
		} else if v.Current_win_bet_area == BET_AREA.PLAYER {
			txt += PLAYER_NODE
		} else {
			panic(xdebug.Funcname())
		}
	}
	return txt
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
