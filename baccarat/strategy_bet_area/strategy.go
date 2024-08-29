/*
功能：下注区域策略-
说明：
策略链中只保存庄和闲
*/
package strategy_bet_area

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/strategy_bet_area/big_road"
	"Odds/baccarat/strategy_bet_area/style"
	"Odds/baccarat/strategy_bet_area/suggestion"
)

const (
	LONG_NODE_CNT = 3 //多少个节点判断为长龙
)

type Strategy struct {
	nodes []*suggestion.ResultNode //策略节点
}

func NewStrategy() *Strategy {
	s := &Strategy{}

	options := style.NewStyleOption(
		style.WithLongNodeCnt(LONG_NODE_CNT))
	style.Style_set_option(options)
	return s
}

// 追加策略节点
func (s *Strategy) Result_node_append(node *suggestion.ResultNode) {
	if node == nil {
		return
	}
	if s.nodes == nil {
		s.Result_node_clear()
	}
	//策略链中不插入tie
	if node.Current_win_bet_area == BET_AREA.TIE {
		return
	}
	s.nodes = append(s.nodes, node)
}

// 查询下注区域
func (s *Strategy) Query_bet_area() *suggestion.BetAreaSuggestion {
	suggestion := style.Style_query(s.nodes)
	//fmt.Println(suggestion.String())
	return suggestion
}

// 查询大路
func (s *Strategy) Query_big_road() *big_road.BigRoad {
	return big_road.NewBigRoadWithNodes(s.nodes)
}

// 清空
func (s *Strategy) Result_node_clear() {
	s.nodes = make([]*suggestion.ResultNode, 0)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
