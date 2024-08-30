/*
功能：双跳形态
说明：

1 标准双跳
🔴🔵
🔴🔵

2 扩展双跳
🔴🔵🔴
🔴🔵
*/
package style

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/define/STYLE"
	"Odds/baccarat/strategy_bet_area/big_road"
	"Odds/baccarat/strategy_bet_area/suggestion"
)

const (
	DOUBLE_JUMP_MIN_COL_CNT  = 2 //最少列数
	DOUBLE_JUMP_MIN_NODE_CNT = 4 //最少节点数
)

// 双跳形态检测
// 前面的两列必须是标准双跳
func check_double_jump_style(nodes []*suggestion.FeedbackNode) (bool, *suggestion.BetAreaSuggestion) {
	//最少节点数校验
	nodes_cnt := len(nodes)
	if nodes_cnt < DOUBLE_JUMP_MIN_NODE_CNT {
		return false, nil
	}

	//最少列数校验
	big_road_all := big_road.NewBigRoadWithNodes(nodes)
	cols_cnt := big_road_all.Col_cnt()
	if cols_cnt < DOUBLE_JUMP_MIN_COL_CNT {
		return false, nil
	}

	last_col := big_road_all.Last_col()
	last_col_node_cnt := last_col.Cnt()
	if last_col_node_cnt > 2 {
		return false, nil
	}

	bet_area := BET_AREA.ERROR

	if last_col_node_cnt == 1 { //最后三列
		if cols_cnt < DOUBLE_JUMP_MIN_COL_CNT+1 {
			return false, nil
		}

		for i := cols_cnt - 3; i < cols_cnt-1; i++ {
			if big_road_all.Get_col(i).Cnt() != 2 {
				return false, nil
			}
		}

		bet_area = last_col.Result_area()
	} else if last_col_node_cnt == 2 { //最后两列
		for i := cols_cnt - 2; i < cols_cnt; i++ {
			if big_road_all.Get_col(i).Cnt() != 2 {
				return false, nil
			}
		}

		if last_col.Result_area() == BET_AREA.BANKER {
			bet_area = BET_AREA.PLAYER
		} else {
			bet_area = BET_AREA.BANKER
		}
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
