/*
功能：形态-
说明：
*/
package style

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/strategy_bet_area/suggestion"
	"fmt"
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
		if nodes[i].Result_win_bet_area == target_bet_area {
			target_cnt++
			continue
		}
		break
	}
	err = nil
	return
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
