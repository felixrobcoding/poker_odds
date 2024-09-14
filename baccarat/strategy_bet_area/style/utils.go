/*
功能：形态-
说明：
*/
package style

import (
	"Odds/baccarat/strategy_bet_area/big_road"
	"errors"
)

type STYLE_TYPE int //形态类型

const (
	HALF_STYLE STYLE_TYPE = 1 //半形态
	FULL_STYLE STYLE_TYPE = 2 //全形态
)

// 检测项
type CheckItem struct {
	style_type            STYLE_TYPE //形态类型
	check_col_cnt         int        //需要检测的列数
	col_node_cnts         []int      //每列节点数
	min_last_col_node_cnt int        //最后一列最小节点数
	max_last_col_node_cnt int        //最后一列最大节点数
}

// 提取需要的列,列节点数分布情况
func extract_col_nodes(big_road *big_road.BigRoad, check_col_cnt int) ([]int, error) {
	col_nodes := make([]int, 0)
	col_cnt := big_road.Col_cnt()
	if check_col_cnt > col_cnt {
		return nil, errors.New("")
	}

	for i := col_cnt - check_col_cnt; i < col_cnt; i++ {
		col_nodes = append(col_nodes, big_road.Get_col(i).Cnt())
	}
	return col_nodes, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
