/*
功能：双跳形态
说明：
*/
package style

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/define/STYLE"
	"Odds/baccarat/strategy_bet_area/big_road"
	"Odds/baccarat/strategy_bet_area/suggestion"
)

const (
	DOUBLE_JUMP_NODE_CNT = 5 //最少需要5个节点
)

// 双跳形态检测
func check_double_jump_style(nodes []*suggestion.ResultNode) (bool, *suggestion.BetAreaSuggestion) {
	nodes_len := len(nodes)
	if nodes_len < DOUBLE_JUMP_NODE_CNT {
		return false, nil
	}
	start_index := nodes_len - DOUBLE_JUMP_NODE_CNT
	part_nodes := nodes[start_index : start_index+DOUBLE_JUMP_NODE_CNT]
	big_road := big_road.NewBigRoadWithNodes(part_nodes)
	if big_road.Col_cnt() != 3 {
		return false, nil
	}
	for i := 0; i < big_road.Col_cnt(); i++ {
		if big_road.Get_col(i).Cnt() > 2 {
			return false, nil
		}
	}

	//中间一列必须是2
	if big_road.Get_col(1).Cnt() != 2 {
		return false, nil
	}

	//最后一列元素个数
	bet_area := BET_AREA.ERROR
	last_col_cnt := big_road.Last_col().Cnt()
	if last_col_cnt == 1 {
		bet_area = big_road.Last_col().Bet_area()
	} else if last_col_cnt == 2 {
		bet_area = big_road.Get_col(1).Bet_area()
	}

	return true, &suggestion.BetAreaSuggestion{
		Style:    STYLE.DOUBLE_JUMP,
		Bet_area: bet_area,
		Comment:  "检测到_双跳_形态",
		Alart:    true,
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
