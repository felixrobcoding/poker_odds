/*
功能：跟从形态
说明：
*/
package style

import (
	"github.com/felixrobcoding/poker_oddsbaccarat/define/STYLE"
	"github.com/felixrobcoding/poker_oddsbaccarat/strategy_bet_area/suggestion"
)

// 跟从形态检测
// 默认的形态
func check_follow_style(nodes []*suggestion.FeedbackNode) (bool, *suggestion.BetAreaSuggestion) {
	nodes_cnt := len(nodes)
	if nodes_cnt <= 0 {
		return true, suggestion.NewBetAreaSuggestion()
	}

	return true, &suggestion.BetAreaSuggestion{
		Style:    STYLE.FOLLOW,
		Bet_area: nodes[nodes_cnt-1].Result_area,
		Comment:  "默认跟随下注",
		Alart:    false,
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
