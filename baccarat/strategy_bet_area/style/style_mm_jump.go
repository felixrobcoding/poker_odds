/*
功能：mm跳形态
说明：两列保持相同颗数的跳
*/
package style

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/define/STYLE"
	"Odds/baccarat/strategy_bet_area/big_road"
	"Odds/baccarat/strategy_bet_area/suggestion"
)

// mm跳形态检测
func check_mm_jump_style(nodes []*suggestion.FeedbackNode) (bool, *suggestion.BetAreaSuggestion) {
	handlers := []HandlerCheckStyle{
		_check_single_jump_style,
		_check_double_jump_style,
		_check_3_jump_style,
		_check_4_jump_style,
	}

	for _, v := range handlers {
		if is, suggestion := v(nodes); is {
			return true, suggestion
		}
	}
	return false, nil
}

// 单跳形态检测
func _check_single_jump_style(nodes []*suggestion.FeedbackNode) (bool, *suggestion.BetAreaSuggestion) {
	const MIN_NODE_CNT = 2 //最小节点数
	const MIN_COL_CNT = 2  // 最小列数

	//最少节点数校验
	nodes_cnt := len(nodes)
	if nodes_cnt < MIN_NODE_CNT {
		return false, nil
	}

	//最少列数校验
	big_road_all := big_road.NewBigRoadWithNodes(nodes)
	cols_cnt := big_road_all.Col_cnt()
	if cols_cnt < MIN_COL_CNT {
		return false, nil
	}

	check_items := []CheckItem{
		{
			/*形态举例
			🔴🔵
			*/
			style_type:            FULL_STYLE,
			check_col_cnt:         2,
			col_node_cnts:         []int{1, 1},
			min_last_col_node_cnt: 1,
			max_last_col_node_cnt: 1,
		},
	}

	is_ok := false
	bet_area := BET_AREA.ERROR

	for i := 0; i < len(check_items); i++ {
		check_item := check_items[i]

		col_nodes, err := extract_col_nodes(big_road_all, check_item.check_col_cnt)
		if err != nil {
			continue
		}
		col_nodes_len := len(col_nodes)
		last_col_node_cnt := col_nodes[col_nodes_len-1]

		//最后一列节点个数校验
		if last_col_node_cnt < check_item.min_last_col_node_cnt || last_col_node_cnt > check_item.max_last_col_node_cnt {
			continue
		}

		if check_item.style_type == FULL_STYLE { //全形态
			is_ok = true
			//校验列的节点数必须保持一致
			for j := col_nodes_len - 1; j >= 0; j-- {
				if check_item.col_node_cnts[j] != col_nodes[j] {
					is_ok = false
					break
				}
			}
			if !is_ok {
				continue
			}

			if big_road_all.Last_col().Result_area() == BET_AREA.BANKER {
				bet_area = BET_AREA.PLAYER
			} else {
				bet_area = BET_AREA.BANKER
			}

			return true, &suggestion.BetAreaSuggestion{
				Style:    STYLE.MM_JUMP,
				Bet_area: bet_area,
				Comment:  "检测到_单跳_形态",
				Alart:    true,
			}
		}
	}
	return false, nil
}

// 双跳形态检测
func _check_double_jump_style(nodes []*suggestion.FeedbackNode) (bool, *suggestion.BetAreaSuggestion) {
	const MIN_NODE_CNT = 4 //最小节点数
	const MIN_COL_CNT = 2  // 最小列数

	//最少节点数校验
	nodes_cnt := len(nodes)
	if nodes_cnt < MIN_NODE_CNT {
		return false, nil
	}

	//最少列数校验
	big_road_all := big_road.NewBigRoadWithNodes(nodes)
	cols_cnt := big_road_all.Col_cnt()
	if cols_cnt < MIN_COL_CNT {
		return false, nil
	}

	check_items := []CheckItem{
		{
			/*形态举例
			🔴🔵🔴
			🔴🔵
			*/
			style_type:            HALF_STYLE,
			check_col_cnt:         2,
			col_node_cnts:         []int{2, 2, -1},
			min_last_col_node_cnt: 1,
			max_last_col_node_cnt: 1,
		},
		{
			/*形态举例
			🔴🔵
			🔴🔵
			*/
			style_type:            FULL_STYLE,
			check_col_cnt:         2,
			col_node_cnts:         []int{2, 2},
			min_last_col_node_cnt: 2,
			max_last_col_node_cnt: 2,
		},
	}

	is_ok := false
	bet_area := BET_AREA.ERROR

	for i := 0; i < len(check_items); i++ {
		check_item := check_items[i]

		col_nodes, err := extract_col_nodes(big_road_all, check_item.check_col_cnt)
		if err != nil {
			continue
		}
		col_nodes_len := len(col_nodes)
		last_col_node_cnt := col_nodes[col_nodes_len-1]

		//最后一列节点个数校验
		if last_col_node_cnt < check_item.min_last_col_node_cnt || last_col_node_cnt > check_item.max_last_col_node_cnt {
			continue
		}

		if check_item.style_type == HALF_STYLE { //半形态
			is_ok = true
			//校验列的节点数必须保持一致
			for j := col_nodes_len - 2; j >= 0; j-- {
				if check_item.col_node_cnts[j] != col_nodes[j] {
					is_ok = false
					break
				}
			}
			if !is_ok {
				continue
			}

			return true, &suggestion.BetAreaSuggestion{
				Style:    STYLE.MM_JUMP,
				Bet_area: big_road_all.Last_col().Result_area(),
				Comment:  "检测到_双跳_形态",
				Alart:    true,
			}
		}

		if check_item.style_type == FULL_STYLE { //全形态
			is_ok = true
			//校验列的节点数必须保持一致
			for j := col_nodes_len - 1; j >= 0; j-- {
				if check_item.col_node_cnts[j] != col_nodes[j] {
					is_ok = false
					break
				}
			}
			if !is_ok {
				continue
			}

			if big_road_all.Last_col().Result_area() == BET_AREA.BANKER {
				bet_area = BET_AREA.PLAYER
			} else {
				bet_area = BET_AREA.BANKER
			}

			return true, &suggestion.BetAreaSuggestion{
				Style:    STYLE.MM_JUMP,
				Bet_area: bet_area,
				Comment:  "检测到_双跳_形态",
				Alart:    true,
			}
		}
	}
	return false, nil
}

// 三跳形态检测
func _check_3_jump_style(nodes []*suggestion.FeedbackNode) (bool, *suggestion.BetAreaSuggestion) {
	const MIN_NODE_CNT = 5 //最小节点数
	const MIN_COL_CNT = 2  // 最小列数

	//最少节点数校验
	nodes_cnt := len(nodes)
	if nodes_cnt < MIN_NODE_CNT {
		return false, nil
	}

	//最少列数校验
	big_road_all := big_road.NewBigRoadWithNodes(nodes)
	cols_cnt := big_road_all.Col_cnt()
	if cols_cnt < MIN_COL_CNT {
		return false, nil
	}

	check_items := []CheckItem{
		{
			/*形态举例
			🔴🔵
			🔴🔵
			🔴
			*/
			style_type:            HALF_STYLE,
			check_col_cnt:         2,
			col_node_cnts:         []int{3, 2},
			min_last_col_node_cnt: 2,
			max_last_col_node_cnt: 2,
		},
		{
			/*形态举例
			🔴🔵🔴
			🔴🔵
			🔴🔵
			*/
			/*形态举例
			🔴🔵🔴
			🔴🔵🔴
			🔴🔵
			*/
			style_type:            HALF_STYLE,
			check_col_cnt:         3,
			col_node_cnts:         []int{3, 3, -1},
			min_last_col_node_cnt: 1,
			max_last_col_node_cnt: 2,
		},
		{
			/*形态举例
			🔴🔵
			🔴🔵
			🔴🔵
			*/
			style_type:            FULL_STYLE,
			check_col_cnt:         2,
			col_node_cnts:         []int{3, 3},
			min_last_col_node_cnt: 3,
			max_last_col_node_cnt: 3,
		},
	}

	is_ok := false
	bet_area := BET_AREA.ERROR

	for i := 0; i < len(check_items); i++ {
		check_item := check_items[i]

		col_nodes, err := extract_col_nodes(big_road_all, check_item.check_col_cnt)
		if err != nil {
			continue
		}
		col_nodes_len := len(col_nodes)
		last_col_node_cnt := col_nodes[col_nodes_len-1]

		//最后一列节点个数校验
		if last_col_node_cnt < check_item.min_last_col_node_cnt || last_col_node_cnt > check_item.max_last_col_node_cnt {
			continue
		}

		if check_item.style_type == HALF_STYLE { //半形态
			is_ok = true
			//校验列的节点数必须保持一致
			for j := col_nodes_len - 2; j >= 0; j-- {
				if check_item.col_node_cnts[j] != col_nodes[j] {
					is_ok = false
					break
				}
			}
			if !is_ok {
				continue
			}

			return true, &suggestion.BetAreaSuggestion{
				Style:    STYLE.MM_JUMP,
				Bet_area: big_road_all.Last_col().Result_area(),
				Comment:  "检测到_三跳_形态",
				Alart:    true,
			}
		}

		if check_item.style_type == FULL_STYLE { //全形态
			is_ok = true
			//校验列的节点数必须保持一致
			for j := col_nodes_len - 1; j >= 0; j-- {
				if check_item.col_node_cnts[j] != col_nodes[j] {
					is_ok = false
					break
				}
			}
			if !is_ok {
				continue
			}

			if big_road_all.Last_col().Result_area() == BET_AREA.BANKER {
				bet_area = BET_AREA.PLAYER
			} else {
				bet_area = BET_AREA.BANKER
			}

			return true, &suggestion.BetAreaSuggestion{
				Style:    STYLE.MM_JUMP,
				Bet_area: bet_area,
				Comment:  "检测到_三跳_形态",
				Alart:    true,
			}
		}
	}
	return false, nil
}

// 四跳形态检测
func _check_4_jump_style(nodes []*suggestion.FeedbackNode) (bool, *suggestion.BetAreaSuggestion) {
	const MIN_NODE_CNT = 7 //最小节点数
	const MIN_COL_CNT = 2  // 最小列数

	//最少节点数校验
	nodes_cnt := len(nodes)
	if nodes_cnt < MIN_NODE_CNT {
		return false, nil
	}

	//最少列数校验
	big_road_all := big_road.NewBigRoadWithNodes(nodes)
	cols_cnt := big_road_all.Col_cnt()
	if cols_cnt < MIN_COL_CNT {
		return false, nil
	}

	check_items := []CheckItem{
		{
			/*形态举例
			🔴🔵
			🔴🔵
			🔴🔵
			🔴
			*/
			style_type:            HALF_STYLE,
			check_col_cnt:         2,
			col_node_cnts:         []int{4, 3},
			min_last_col_node_cnt: 3,
			max_last_col_node_cnt: 3,
		},
		{
			/*形态举例
			🔴🔵🔴
			🔴🔵
			🔴🔵
			🔴🔵
			*/
			/*形态举例
			🔴🔵🔴
			🔴🔵🔴
			🔴🔵
			🔴🔵
			*/
			/*形态举例
			🔴🔵🔴
			🔴🔵🔴
			🔴🔵🔴
			🔴🔵
			*/
			style_type:            HALF_STYLE,
			check_col_cnt:         3,
			col_node_cnts:         []int{4, 4, -1},
			min_last_col_node_cnt: 1,
			max_last_col_node_cnt: 3,
		},
		{
			/*形态举例
			🔴🔵
			🔴🔵
			🔴🔵
			🔴🔵
			*/
			style_type:            FULL_STYLE,
			check_col_cnt:         2,
			col_node_cnts:         []int{4, 4},
			min_last_col_node_cnt: 4,
			max_last_col_node_cnt: 4,
		},
	}

	is_ok := false
	bet_area := BET_AREA.ERROR

	for i := 0; i < len(check_items); i++ {
		check_item := check_items[i]

		col_nodes, err := extract_col_nodes(big_road_all, check_item.check_col_cnt)
		if err != nil {
			continue
		}
		col_nodes_len := len(col_nodes)
		last_col_node_cnt := col_nodes[col_nodes_len-1]

		//最后一列节点个数校验
		if last_col_node_cnt < check_item.min_last_col_node_cnt || last_col_node_cnt > check_item.max_last_col_node_cnt {
			continue
		}

		if check_item.style_type == HALF_STYLE { //半形态
			is_ok = true
			//校验列的节点数必须保持一致
			for j := col_nodes_len - 2; j >= 0; j-- {
				if check_item.col_node_cnts[j] != col_nodes[j] {
					is_ok = false
					break
				}
			}
			if !is_ok {
				continue
			}

			return true, &suggestion.BetAreaSuggestion{
				Style:    STYLE.MM_JUMP,
				Bet_area: big_road_all.Last_col().Result_area(),
				Comment:  "检测到_四跳_形态",
				Alart:    true,
			}
		}

		if check_item.style_type == FULL_STYLE { //全形态
			is_ok = true
			//校验列的节点数必须保持一致
			for j := col_nodes_len - 1; j >= 0; j-- {
				if check_item.col_node_cnts[j] != col_nodes[j] {
					is_ok = false
					break
				}
			}
			if !is_ok {
				continue
			}

			if big_road_all.Last_col().Result_area() == BET_AREA.BANKER {
				bet_area = BET_AREA.PLAYER
			} else {
				bet_area = BET_AREA.BANKER
			}

			return true, &suggestion.BetAreaSuggestion{
				Style:    STYLE.MM_JUMP,
				Bet_area: bet_area,
				Comment:  "检测到_四跳_形态",
				Alart:    true,
			}
		}
	}
	return false, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
