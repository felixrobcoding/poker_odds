/*
功能：下注额策略-全下
说明：
*/
package strategy_bet_amount

import (
	"Odds/common/BET_AMOUNT_STRATEGY"
	"errors"

	"github.com/poker-x-studio/x/xmath"
)

type betAmountAllIn struct {
	BetAmountStrategy
}

func newBetAmountAllIn(init_chip int) IBetAmountStrategy {
	b := &betAmountAllIn{}
	b.init(init_chip)
	return b
}

// 初始化
func (b *betAmountAllIn) init(init_chip int) {
	b.t = BET_AMOUNT_STRATEGY.ALL_IN
	b.set_chip(init_chip)
}

// 查询下注额
func (b *betAmountAllIn) Query_bet_amount() (int, error) {
	bet := 0

	len := len(b.result_nodes)
	if len <= 0 {
		bet = xmath.Min(MAX_BET, b.init_chip)
	} else {
		last_node := b.result_nodes[len-1]
		bet = xmath.Min(MAX_BET, int(last_node.Current_chip))
	}

	if b.Is_enough_money(bet) {
		return bet, nil
	}
	return 0, errors.New(NOT_ENOUGH_MONEY)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
