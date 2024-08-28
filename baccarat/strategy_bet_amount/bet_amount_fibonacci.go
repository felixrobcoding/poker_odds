/*
功能：下注额策略-斐波那契策略
说明：
*/
package strategy_bet_amount

import (
	"Odds/common/BET_AMOUNT_STRATEGY"
	"errors"

	"github.com/poker-x-studio/x/xmath"
)

type betAmountFibonacci struct {
	BetAmountStrategy
}

func newBetAmountFibonacci(init_chip int) IBetAmountStrategy {
	b := &betAmountFibonacci{}
	b.init(init_chip)
	return b
}

// 初始化
func (b *betAmountFibonacci) init(init_chip int) {
	b.t = BET_AMOUNT_STRATEGY.FIBONACCI
	b.set_chip(init_chip)
}

// 查询下注额
// 斐波那契数就是由之前的两数相加而得出
func (b *betAmountFibonacci) Query_bet_amount() (int, error) {
	bet := MIN_BET

	len := len(b.result_nodes)
	if len <= 1 {
		bet = MIN_BET
	} else if len <= 2 {
		bet = MIN_BET * 2
	} else {
		last_node := b.result_nodes[len-1]
		last_2_node := b.result_nodes[len-2]
		bet = xmath.Min(MAX_BET, last_node.Min_bet()+last_2_node.Min_bet())
	}

	if b.Is_enough_money(bet) {
		return bet, nil
	}
	return 0, errors.New(NOT_ENOUGH_MONEY)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
