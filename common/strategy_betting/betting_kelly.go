/*
功能：下注策略-凯利策略
说明：
*/
package strategy_betting

import (
	"Odds/common/BETTING_TYPE"
	"errors"

	"github.com/poker-x-studio/x/xmath"
)

const (
	KELLY_PERCENTAGE = 0.05
)

type bettingKelly struct {
	Betting
}

func newBettingKelly(init_chip int) IBettingStrategy {
	b := &bettingKelly{}
	b.init(init_chip)
	return b
}

// 初始化
func (b *bettingKelly) init(init_chip int) {
	b.t = BETTING_TYPE.KELLY
	b.set_chip(init_chip)
}

// 查询下注额
// 每次投注总筹码的某一百分比
func (b *bettingKelly) Query_bet() (int, error) {
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
