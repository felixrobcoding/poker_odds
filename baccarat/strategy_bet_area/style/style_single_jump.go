/*
功能：单跳形态
说明：

1 标准单跳
🔴🔵🔴

2 扩展单跳
🔴🔵🔴
🔴
*/
package style

import (
	"Odds/baccarat/define/STYLE"
	"Odds/baccarat/strategy_bet_area/big_road"
	"Odds/baccarat/strategy_bet_area/suggestion"
)

const (
	SINGLE_JUMP_MIN_NODE_CNT = 3 //最少节点数
	SINGLE_JUMP_MIN_COL_CNT  = 3 //最少列数
)

// 单跳形态检测
func check_single_jump_style(nodes []*suggestion.FeedbackNode) (bool, *suggestion.BetAreaSuggestion) {
	//最少节点数校验
	nodes_cnt := len(nodes)
	if nodes_cnt < SINGLE_JUMP_MIN_NODE_CNT {
		return false, nil
	}

	//最少列数校验
	big_road_all := big_road.NewBigRoadWithNodes(nodes)
	cols_cnt := big_road_all.Col_cnt()
	if cols_cnt < SINGLE_JUMP_MIN_COL_CNT {
		return false, nil
	}

	start_index := nodes_cnt - SINGLE_JUMP_MIN_NODE_CNT
	part_nodes := nodes[start_index : start_index+SINGLE_JUMP_MIN_NODE_CNT]
	big_road := big_road.NewBigRoadWithNodes(part_nodes)
	if big_road.Col_cnt() != SINGLE_JUMP_MIN_NODE_CNT {
		return false, nil
	}

	return true, &suggestion.BetAreaSuggestion{
		Style:    STYLE.SINGLE_JUMP,
		Bet_area: big_road.Get_col(big_road.Col_cnt() - 2).Result_area(),
		Comment:  "检测到_单跳_形态",
		Alart:    true,
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
