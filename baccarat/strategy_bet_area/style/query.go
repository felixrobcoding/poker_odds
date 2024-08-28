/*
功能：形态-查询
说明：
*/
package style

import (
	"Odds/baccarat/strategy_bet_area/suggestion"

	"github.com/poker-x-studio/x/xdebug"
)

type HandlerCheckStyle func([]*suggestion.ResultNode) (bool, *suggestion.BetAreaSuggestion)

// 查询形态
func Style_query(nodes []*suggestion.ResultNode) *suggestion.BetAreaSuggestion {
	if style_option == nil {
		panic(xdebug.Funcname())
	}
	handlers := []HandlerCheckStyle{
		check_long_style,
		check_single_jump_style,
		check_double_jump_style,
		check_follow_style,
	}

	for _, v := range handlers {
		if is, suggestion := v(nodes); is {
			return suggestion
		}
	}
	panic(xdebug.Funcname())
}

// 设置选项
func Style_set_option(option *StyleOption) {
	style_option = option
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
