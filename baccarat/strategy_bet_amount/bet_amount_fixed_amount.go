/*
功能：下注额策略-固定额度
说明：
*/
package strategy_bet_amount

import (
	"Odds/common/BET_AMOUNT_STRATEGY"
	"errors"
)

type betAmountFixedAmount struct {
	BetAmountStrategy
}

func newBetAmountFixedAmount(init_chip int) IBetAmountStrategy {
	b := &betAmountFixedAmount{}
	b.init(init_chip)
	return b
}

// 初始化
func (b *betAmountFixedAmount) init(init_chip int) {
	b.t = BET_AMOUNT_STRATEGY.FIXED_AMOUNT
	b.set_chip(init_chip)
}

// 查询下注额
func (b *betAmountFixedAmount) Query_bet_amount() (int, error) {
	bet := MIN_BET

	if b.Is_enough_money(bet) {
		return bet, nil
	}
	return 0, errors.New(NOT_ENOUGH_MONEY)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
