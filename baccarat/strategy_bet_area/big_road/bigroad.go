/*
功能：大路
说明：
*/
package big_road

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/strategy_bet_area/suggestion"

	"github.com/poker-x-studio/x/xdebug"
)

// 大路
type BigRoad struct {
	cols []Col //列
}

func NewBigRoad(nodes []Node) *BigRoad {
	road := &BigRoad{}
	road.init()
	road.push(nodes)
	return road
}

// NewBigRoadWithNodes 策略节点链构造大路
func NewBigRoadWithNodes(suggestion_nodes []*suggestion.FeedbackNode) *BigRoad {
	nodes := make([]Node, 0)
	for i := 0; i < len(suggestion_nodes); i++ {
		node := Node{
			bet_area:     suggestion_nodes[i].Current_bet_area,
			bet_amount:   suggestion_nodes[i].Current_bet_amount,
			result_area:  suggestion_nodes[i].Result_area,
			result_score: suggestion_nodes[i].Result_score,
		}
		nodes = append(nodes, node)
	}
	return NewBigRoad(nodes)
}

func NewBigRoadWithCols(cols []*Col) *BigRoad {
	nodes := make([]Node, 0)
	for i := 0; i < len(cols); i++ {
		col := cols[i]

		for j := 0; j < col.Cnt(); j++ {
			node := Node{
				bet_area:     col.Get_node(j).bet_area,
				bet_amount:   col.Get_node(j).bet_amount,
				result_area:  col.Get_node(j).result_area,
				result_score: col.Get_node(j).result_score,
			}
			nodes = append(nodes, node)
		}
	}
	return NewBigRoad(nodes)
}

func (b *BigRoad) init() {
	b.cols = make([]Col, 0)
}

// push 插入多个元素
func (b *BigRoad) push(nodes []Node) {
	for _, v := range nodes {
		b.push_element(v)
	}
}

// push_element 插入元素
func (b *BigRoad) push_element(node Node) {
	if b.cols == nil {
		b.cols = make([]Col, 0)
	}

	if b.Col_cnt() == 0 {
		col := Col{}
		col.push(node)

		b.cols = append(b.cols, col)
		return
	}

	//最后一列
	col := b.Last_col()
	if col.Result_area() == node.result_area {
		col.push(node)
	} else { //新创建1列
		col := Col{}
		col.push(node)

		b.cols = append(b.cols, col)
	}
}

// Col_cnt 列数
func (b *BigRoad) Col_cnt() int {
	return len(b.cols)
}

// Col_max_node_cnt 最长列节点个数
func (b *BigRoad) Col_max_node_cnt() int {
	max_node_cnt := 0

	for col_index := 0; col_index < b.Col_cnt(); col_index++ {
		if b.Get_col(col_index).Cnt() > max_node_cnt {
			max_node_cnt = b.Get_col(col_index).Cnt()
		}
	}

	return max_node_cnt
}

// Total_cnt 总个数
func (b *BigRoad) Total_cnt() int {
	total_cnt := 0

	for col_index := 0; col_index < b.Col_cnt(); col_index++ {
		total_cnt += b.Get_col(col_index).Cnt()
	}
	return total_cnt
}

// Extract_stat_for_svg 统计信息
func (b *BigRoad) Extract_stat_for_svg() (banker_cnt int, player_cnt int, win_cnt int, lose_cnt int, total_bet_amont int, total_result_score float64) {

	for col_index := 0; col_index < b.Col_cnt(); col_index++ {
		col := b.Get_col(col_index)

		for row_index := 0; row_index < col.Cnt(); row_index++ {
			node := col.Get_node(row_index)

			if node.result_area == BET_AREA.BANKER {
				banker_cnt++
			}
			if node.result_area == BET_AREA.PLAYER {
				player_cnt++
			}

			if node.bet_area == node.result_area {
				win_cnt++
			} else {
				lose_cnt++
			}

			total_bet_amont += node.bet_amount
			total_result_score += node.result_score
		}
	}
	return
}

// Last_col 最后的一列
func (b *BigRoad) Last_col() *Col {
	if b.cols == nil || b.Col_cnt() == 0 {
		panic(xdebug.Funcname())
	}
	return &b.cols[len(b.cols)-1]
}

// Get_col 获取列
func (b *BigRoad) Get_col(col_index int) *Col {
	if col_index < 0 || col_index >= b.Col_cnt() {
		panic(xdebug.Funcname())
	}
	return &b.cols[col_index]
}

// Extract_bigroad_stat 提取列统计
func (b *BigRoad) Extract_bigroad_stat() *BigRoadStat {
	tmp_map := make(map[int]int, 0)
	for _, v := range b.cols {
		tmp_map[v.Cnt()]++
	}
	stat := &BigRoadStat{}
	for k, v := range tmp_map {
		col_stat := ColStat{
			Hands_per_col: k,
			Cols_cnt:      v,
		}
		stat.Col_stats = append(stat.Col_stats, col_stat)
	}
	stat.sort()
	return stat
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
