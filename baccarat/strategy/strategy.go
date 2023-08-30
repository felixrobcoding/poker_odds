/*
功能：玩法策略-
说明：
策略链中只保存庄和闲
*/
package strategy

import (
	"Odds/baccarat/define"
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/strategy/big_road"
	"Odds/baccarat/strategy/style"
)

const (
	SINGLE_JUMP_BET_TIMES = 10.0
	DOUBLE_JUMP_BET_TIMES = 10.0
	LONG_NODE_CNT         = 3
	LONG_BET_TIMES        = 10.0
)

type Strategy struct {
	init_chip int                    //起始筹码
	nodes     []*define.StrategyNode //策略节点
}

func NewStrategy(init_chip int) *Strategy {
	s := &Strategy{}
	s.set_chip(init_chip)

	options := style.NewStyleOption(
		style.WithSingleJumpBetTimes(SINGLE_JUMP_BET_TIMES),
		style.WithDoubleJumpBetTimes(DOUBLE_JUMP_BET_TIMES),
		style.WithLongNodeCnt(LONG_NODE_CNT),
		style.WithLongBetTimes(LONG_BET_TIMES))
	style.Style_set_option(options)
	return s
}

// 设置起始筹码
func (s *Strategy) set_chip(init_chip int) {
	s.init_chip = init_chip
}

// 追加策略节点
func (s *Strategy) Strategy_node_append(node *define.StrategyNode) {
	if node == nil {
		return
	}
	if s.nodes == nil {
		s.Strategy_node_clear()
	}
	//策略链中不插入tie
	if node.Current_win_bet_area == BET_AREA.TIE {
		return
	}
	s.nodes = append(s.nodes, node)
}

// 查询策略建议
func (s *Strategy) Query_strategy_suggestion() *define.StrategySuggestion {
	suggestion := style.Style_query(s.nodes)
	//fmt.Println(suggestion.String())
	return suggestion
}

// 查询大路
func (s *Strategy) Query_big_road() *big_road.BigRoad {
	return big_road.NewBigRoadWithNodes(s.nodes)
}

// 清空
func (s *Strategy) Strategy_node_clear() {
	s.nodes = make([]*define.StrategyNode, 0)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
