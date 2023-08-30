/*
功能：龙形态
说明：

https://bbs.boniu123.cc/thread-834529-1-1.html
长庄：在大路中连续有 4 个或 4 个以上的庄（如下图所示）
长闲：在大路中连续有 4 个或 4 个以上的闲（如下图所示）
*/
package style

import (
	"Odds/baccarat/define"
	"Odds/baccarat/define/STYLE"
	"Odds/baccarat/strategy/big_road"
)

// 龙形态检测
func check_long_style(nodes []*define.StrategyNode) (bool, *define.StrategySuggestion) {
	len := len(nodes)
	long_node_cnt := style_option.long_node_cnt
	if len < long_node_cnt {
		return false, nil
	}

	start_index := len - long_node_cnt
	part_nodes := nodes[start_index : start_index+long_node_cnt]
	big_road := big_road.NewBigRoadWithNodes(part_nodes)
	if big_road.Col_cnt() != 1 {
		return false, nil
	}

	return true, &define.StrategySuggestion{
		Style:     STYLE.LONG,
		Bet_area:  big_road.Last_col().Bet_area(),
		Bet_times: style_option.long_bet_times,
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
