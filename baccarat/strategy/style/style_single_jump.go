/*
功能：单跳形态
说明：
*/
package style

import (
	"Odds/baccarat/define"
	"Odds/baccarat/define/STYLE"
	"Odds/baccarat/strategy/big_road"
)

const (
	SINGLE_JUMP_NODE_CNT = 3
)

// 单跳形态检测
func check_single_jump_style(nodes []*define.StrategyNode) (bool, *define.StrategySuggestion) {
	nodes_len := len(nodes)
	if nodes_len < SINGLE_JUMP_NODE_CNT {
		return false, nil
	}

	start_index := nodes_len - SINGLE_JUMP_NODE_CNT
	part_nodes := nodes[start_index : start_index+SINGLE_JUMP_NODE_CNT]
	big_road := big_road.NewBigRoadWithNodes(part_nodes)
	if big_road.Col_cnt() != SINGLE_JUMP_NODE_CNT {
		return false, nil
	}

	return true, &define.StrategySuggestion{
		Style:     STYLE.SINGLE_JUMP,
		Bet_area:  big_road.Get_col(big_road.Col_cnt() - 2).Bet_area(),
		Bet_times: style_option.single_jump_bet_times,
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
