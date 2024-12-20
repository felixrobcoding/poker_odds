/*
功能：下注额策略-全下
说明：
*/
package strategy_bet_amount

import (
	"errors"
	"github.com/felixrobcoding/poker_oddscommon/BET_AMOUNT_STRATEGY"

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
	bet := MIN_BET

	len := len(b.feedback_nodes)
	if len <= 0 {
		bet = b.init_chip
	} else {
		last_node := b.feedback_nodes[len-1]
		bet = int(last_node.Current_chip)
	}

	bet = xmath.Max(MIN_BET, bet)
	bet = xmath.Min(MAX_BET, bet)

	if b.Is_enough_money(bet) {
		return bet, nil
	}
	return 0, errors.New(NOT_ENOUGH_MONEY)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
