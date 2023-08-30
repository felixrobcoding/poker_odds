/*
功能：下注策略-固定额度
说明：
*/
package strategy_betting

import (
	"Odds/common/BETTING_TYPE"
	"errors"
)

type bettingFixedAmount struct {
	Betting
}

func newBettingFixedAmount(init_chip int) IBettingStrategy {
	b := &bettingFixedAmount{}
	b.init(init_chip)
	return b
}

// 初始化
func (b *bettingFixedAmount) init(init_chip int) {
	b.t = BETTING_TYPE.FIXED_AMOUNT
	b.set_chip(init_chip)
}

// 查询下注额
func (b *bettingFixedAmount) Query_bet() (int, error) {
	bet := MIN_BET

	if b.Is_enough_money(bet) {
		return bet, nil
	}
	return 0, errors.New(NOT_ENOUGH_MONEY)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
