/*
功能：大路
说明：
*/
package big_road

import (
	"Odds/baccarat/define"
	"Odds/baccarat/define/BET_AREA"

	"github.com/poker-x-studio/x/xdebug"
)

// 大路
type BigRoad struct {
	cols []Col //列
}

func NewBigRoad(bet_areas []BET_AREA.TYPE) *BigRoad {
	road := &BigRoad{}
	road.init()
	road.push(bet_areas)
	return road
}

// NewBigRoadWithNodes 策略节点链构造大路
func NewBigRoadWithNodes(nodes []*define.StrategyNode) *BigRoad {
	win_bet_areas := make([]BET_AREA.TYPE, 0)
	for i := 0; i < len(nodes); i++ {
		win_bet_areas = append(win_bet_areas, nodes[i].Current_win_bet_area)
	}
	return NewBigRoad(win_bet_areas)
}

func (b *BigRoad) init() {
	b.cols = make([]Col, 0)
}

// push 插入多个元素
func (b *BigRoad) push(bet_areas []BET_AREA.TYPE) {
	for _, v := range bet_areas {
		b.push_element(v)
	}
}

// push_element 插入元素
func (b *BigRoad) push_element(bet_area BET_AREA.TYPE) {
	if b.cols == nil {
		b.cols = make([]Col, 0)
	}

	if b.Col_cnt() == 0 {
		col := Col{}
		col.push(bet_area)

		b.cols = append(b.cols, col)
		return
	}

	col := b.Last_col()
	if col.Bet_area() == bet_area {
		col.push(bet_area)
	} else { //新创建1列
		col := Col{}
		col.push(bet_area)

		b.cols = append(b.cols, col)
	}
}

// Col_cnt 列数
func (b *BigRoad) Col_cnt() int {
	return len(b.cols)
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
