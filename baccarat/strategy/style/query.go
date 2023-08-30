/*
功能：形态-查询
说明：
*/
package style

import (
	"Odds/baccarat/define"

	"github.com/poker-x-studio/x/xdebug"
)

type HandlerCheckStyle func([]*define.StrategyNode) (bool, *define.StrategySuggestion)

// 查询形态
func Style_query(nodes []*define.StrategyNode) *define.StrategySuggestion {
	if style_option == nil {
		panic(xdebug.Funcname())
	}
	handlers := []HandlerCheckStyle{
		check_long_style,
		check_follow_style,
		check_single_jump_style,
		check_double_jump_style,
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
