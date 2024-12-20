/*
功能：下注区域策略-
说明：
策略链中只保存庄和闲
*/
package strategy_bet_area

import (
	"github.com/felixrobcoding/poker_oddsbaccarat/define/BET_AREA"
	"github.com/felixrobcoding/poker_oddsbaccarat/strategy_bet_area/big_road"
	"github.com/felixrobcoding/poker_oddsbaccarat/strategy_bet_area/style"
	"github.com/felixrobcoding/poker_oddsbaccarat/strategy_bet_area/suggestion"
)

type Strategy struct {
	nodes []*suggestion.FeedbackNode //反馈节点
}

func NewStrategy() *Strategy {
	s := &Strategy{}
	return s
}

// 追加反馈节点
func (s *Strategy) Feedback_node_append(node *suggestion.FeedbackNode) {
	if node == nil {
		return
	}
	if s.nodes == nil {
		s.Feedback_node_clear()
	}
	//策略链中不插入tie
	if node.Result_area == BET_AREA.TIE {
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
func (s *Strategy) Feedback_node_clear() {
	s.nodes = make([]*suggestion.FeedbackNode, 0)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
