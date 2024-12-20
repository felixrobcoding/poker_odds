/*
功能：形态-查询
说明：
*/
package style

import (
	"github.com/felixrobcoding/poker_oddsbaccarat/strategy_bet_area/suggestion"

	"github.com/poker-x-studio/x/xdebug"
)

type HandlerCheckStyle func([]*suggestion.FeedbackNode) (bool, *suggestion.BetAreaSuggestion)

// 查询形态
func Style_query(nodes []*suggestion.FeedbackNode) *suggestion.BetAreaSuggestion {
	handlers := []HandlerCheckStyle{
		check_mm_jump_style,
		check_mn_jump_style,
		check_long_style,
		check_reverse_style,
		check_follow_style,
	}

	for _, v := range handlers {
		if is, suggestion := v(nodes); is {
			return suggestion
		}
	}
	panic(xdebug.Funcname())
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
