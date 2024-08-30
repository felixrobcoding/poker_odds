/*
功能：龙形态
说明：

https://bbs.boniu123.cc/thread-834529-1-1.html
长庄：在大路中连续有 4 个或 4 个以上的庄（如下图所示）
长闲：在大路中连续有 4 个或 4 个以上的闲（如下图所示）
*/
package style

import (
	"Odds/baccarat/define/STYLE"
	"Odds/baccarat/strategy_bet_area/big_road"
	"Odds/baccarat/strategy_bet_area/suggestion"
)

const (
	LONG_MIN_COL_CNT  = 1 //最少列数
	LONG_MIN_NODE_CNT = 3 //最少节点数
)

// 龙形态检测
func check_long_style(nodes []*suggestion.FeedbackNode) (bool, *suggestion.BetAreaSuggestion) {
	nodes_cnt := len(nodes)
	if nodes_cnt < LONG_MIN_NODE_CNT {
		return false, nil
	}

	start_index := nodes_cnt - LONG_MIN_NODE_CNT
	part_nodes := nodes[start_index : start_index+LONG_MIN_NODE_CNT]
	big_road := big_road.NewBigRoadWithNodes(part_nodes)
	if big_road.Col_cnt() != 1 {
		return false, nil
	}

	return true, &suggestion.BetAreaSuggestion{
		Style:    STYLE.LONG,
		Bet_area: big_road.Last_col().Result_area(),
		Comment:  "检测到_长龙_形态",
		Alart:    true,
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
