/*
功能：跟从形态
说明：
*/
package style

import (
	"Odds/baccarat/define"
	"Odds/baccarat/define/STYLE"
)

// 跟从形态检测
// 默认的形态
func check_follow_style(nodes []*define.StrategyNode) (bool, *define.StrategySuggestion) {
	len := len(nodes)
	if len <= 0 {
		return true, define.NewStrategySuggestion()
	}

	return true, &define.StrategySuggestion{
		Style:     STYLE.FOLLOW,
		Bet_area:  nodes[len-1].Current_win_bet_area,
		Bet_times: define.DEFAULT_BET_TIMES,
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
