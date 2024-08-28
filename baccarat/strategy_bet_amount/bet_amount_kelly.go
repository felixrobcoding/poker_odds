/*
功能：下注额策略-凯利策略
说明：
*/
package strategy_bet_amount

import (
	"Odds/common/BET_AMOUNT_STRATEGY"
	"errors"

	"github.com/poker-x-studio/x/xmath"
)

const (
	KELLY_PERCENTAGE = 0.05
)

type betAmountKelly struct {
	BetAmountStrategy
}

func newBetAmountKelly(init_chip int) IBetAmountStrategy {
	b := &betAmountKelly{}
	b.init(init_chip)
	return b
}

// 初始化
func (b *betAmountKelly) init(init_chip int) {
	b.t = BET_AMOUNT_STRATEGY.KELLY
	b.set_chip(init_chip)
}

// 查询下注额
// 每次投注总筹码的某一百分比
func (b *betAmountKelly) Query_bet_amount() (int, error) {
	bet := MIN_BET

	len := len(b.result_nodes)
	if len <= 0 {
		bet = int(float64(b.init_chip) * KELLY_PERCENTAGE)
	} else {
		last_node := b.result_nodes[len-1]
		bet = xmath.Min(MAX_BET, int(float64(last_node.Current_chip*KELLY_PERCENTAGE)))
	}

	if b.Is_enough_money(bet) {
		return bet, nil
	}
	return 0, errors.New(NOT_ENOUGH_MONEY)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
