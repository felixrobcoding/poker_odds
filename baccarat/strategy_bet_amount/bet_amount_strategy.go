/*
功能：下注额策略-基类
说明：
*/
package strategy_bet_amount

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/common/BET_AMOUNT_STRATEGY"
)

const (
	MIN_BET          = 10     //最小下注
	MAX_BET          = 1000   //最大下注
	NOT_ENOUGH_MONEY = "余额不足" //
)

type BetAmountStrategy struct {
	t              BET_AMOUNT_STRATEGY.TYPE
	feedback_nodes []*FeedbackNode //反馈节点
	init_chip      int             //起始筹码
}

func NewBetAmountStrategy(t BET_AMOUNT_STRATEGY.TYPE, init_chip int) IBetAmountStrategy {
	types := []BET_AMOUNT_STRATEGY.TYPE{
		BET_AMOUNT_STRATEGY.ALL_IN,
		BET_AMOUNT_STRATEGY.FIXED_AMOUNT,
		BET_AMOUNT_STRATEGY.MARTEGAL,
		BET_AMOUNT_STRATEGY.FIBONACCI,
		BET_AMOUNT_STRATEGY.KELLY,
		BET_AMOUNT_STRATEGY.MARTEGAL_N,
		BET_AMOUNT_STRATEGY.FIBONACCI_EX,
	}
	funcs := []func(int) IBetAmountStrategy{
		newBetAmountAllIn,
		newBetAmountFixedAmount,
		newBetAmountMartegal,
		newBetAmountFibonacci,
		newBetAmountKelly,
		newBetAmountMartegalN,
		newBetAmountFibonacciEx,
	}

	for k, v := range types {
		if v == t {
			return funcs[k](init_chip)
		}
	}
	panic("")
}

// 设置起始筹码
func (b *BetAmountStrategy) set_chip(init_chip int) {
	b.init_chip = init_chip
}

// 追加反馈节点
func (b *BetAmountStrategy) Feedback_node_append(node *FeedbackNode) {
	if node == nil {
		return
	}

	//策略链中不插入tie
	if node.Result_area == BET_AREA.TIE {
		return
	}

	if b.feedback_nodes == nil {
		b.Feedback_node_clear()
	}
	b.feedback_nodes = append(b.feedback_nodes, node)
}

// 反馈节点清除
func (b *BetAmountStrategy) Feedback_node_clear() {
	b.feedback_nodes = make([]*FeedbackNode, 0)
}

// 查询余额是否足够
func (b *BetAmountStrategy) Is_enough_money(amount int) bool {
	if amount < MIN_BET {
		return false
	}

	len := len(b.feedback_nodes)
	if len <= 0 {
		if b.init_chip >= amount {
			return true
		}
		return false
	}

	last_node := b.feedback_nodes[len-1]
	if int(last_node.Current_chip) >= amount {
		return true
	}
	return false
}

// 查询配置
func (b *BetAmountStrategy) Query_option() (int, int, BET_AMOUNT_STRATEGY.TYPE) {
	return MIN_BET, MAX_BET, b.t
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
