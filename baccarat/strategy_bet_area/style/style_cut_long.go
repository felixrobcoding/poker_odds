/*
功能：断龙形态
说明：

1 形态举例
🔴🔵🔴[断龙后的第二颗]
🔴
x[x>=0]

2 形态举例
🔵🔴🔵[断龙后的第二颗]
🔵
x[x>=0]
*/
package style

import (
	"Odds/baccarat/define/STYLE"
	"Odds/baccarat/strategy_bet_area/big_road"
	"Odds/baccarat/strategy_bet_area/suggestion"
)

const (
	CUT_LONG_MIN_NODE_CNT = 3 //最少节点数
	CUT_LONG_MIN_COL_CNT  = 2 //最少列数
)

// 断龙检测
func check_cut_long_style(nodes []*suggestion.FeedbackNode) (bool, *suggestion.BetAreaSuggestion) {
	//最少节点数校验
	nodes_cnt := len(nodes)
	if nodes_cnt < CUT_LONG_MIN_NODE_CNT {
		return false, nil
	}

	//最少列数校验
	big_road_all := big_road.NewBigRoadWithNodes(nodes)
	cols_cnt := big_road_all.Col_cnt()
	if cols_cnt < CUT_LONG_MIN_COL_CNT {
		return false, nil
	}

	//最后一列只有一颗
	last_col := big_road_all.Last_col()
	last_col_node_cnt := last_col.Cnt()
	if last_col_node_cnt != 1 {
		return false, nil
	}

	//倒数第二列是龙[颗数>=2]
	second_last_col := big_road_all.Get_col(cols_cnt - 2)
	second_last_col_node_cnt := second_last_col.Cnt()
	if second_last_col_node_cnt < 2 {
		return false, nil
	}

	bet_area := second_last_col.Result_area()

	return true, &suggestion.BetAreaSuggestion{
		Style:    STYLE.CUT_LONG,
		Bet_area: bet_area,
		Comment:  "断龙下注",
		Alart:    true,
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
