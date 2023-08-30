/*
功能：下注策略-基类
说明：
*/
package strategy_betting

import (
	"Odds/common/BETTING_TYPE"
)

const (
	MIN_BET          = 10   //最小下注
	MAX_BET          = 1000 //最大下注
	NOT_ENOUGH_MONEY = "余额不足"
)

type Betting struct {
	t            BETTING_TYPE.TYPE
	result_nodes []*ResultNode //结果走势
	init_chip    int           //起始筹码
}

func NewBetting(t BETTING_TYPE.TYPE, init_chip int) IBettingStrategy {
	types := []BETTING_TYPE.TYPE{
		BETTING_TYPE.ALL_IN,
		BETTING_TYPE.FIXED_AMOUNT,
		BETTING_TYPE.MARTEGAL,
		BETTING_TYPE.FIBONACCI,
		BETTING_TYPE.KELLY,
	}
	funcs := []func(int) IBettingStrategy{
		newBettingAllIn,
		newBettingFixedAmount,
		newBettingMartegal,
		newBettingFibonacci,
		newBettingKelly,
	}

	for k, v := range types {
		if v == t {
			return funcs[k](init_chip)
		}
	}
	panic("")
}

// 设置起始筹码
func (b *Betting) set_chip(init_chip int) {
	b.init_chip = init_chip
}

// 追加结果
func (b *Betting) Result_node_append(trend *ResultNode) {
	if b.result_nodes == nil {
		b.result_nodes = make([]*ResultNode, 0)
	}
	b.result_nodes = append(b.result_nodes, trend)
}

// 清理结果
func (b *Betting) Result_node_clear() {
	b.result_nodes = make([]*ResultNode, 0)
}

// 查询余额是否足够
func (b *Betting) Is_enough_money(payout_amount int) bool {
	if payout_amount < MIN_BET {
		return false
	}

	len := len(b.result_nodes)
	if len <= 0 {
		if b.init_chip >= payout_amount {
			return true
		}
		return false
	}

	last_node := b.result_nodes[len-1]
	if int(last_node.Current_chip) >= payout_amount {
		return true
	}
	return false
}

// 查询配置
func (b *Betting) Query_option() (int, int, BETTING_TYPE.TYPE) {
	return MIN_BET, MAX_BET, b.t
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
