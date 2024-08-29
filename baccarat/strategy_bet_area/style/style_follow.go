/*
功能：跟从形态
说明：
*/
package style

import (
	"Odds/baccarat/define/STYLE"
	"Odds/baccarat/strategy_bet_area/suggestion"
)

// 跟从形态检测
// 默认的形态
func check_follow_style(nodes []*suggestion.ResultNode) (bool, *suggestion.BetAreaSuggestion) {
	len := len(nodes)
	if len <= 0 {
		return true, suggestion.NewBetAreaSuggestion()
	}

	return true, &suggestion.BetAreaSuggestion{
		Style:    STYLE.FOLLOW,
		Bet_area: nodes[len-1].Current_win_bet_area,
		Comment:  "默认跟随下注",
		Alart:    false,
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
