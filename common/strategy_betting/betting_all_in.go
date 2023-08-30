/*
功能：下注策略-全下
说明：
*/
package strategy_betting

import (
	"Odds/common/BETTING_TYPE"
	"errors"

	"github.com/poker-x-studio/x/xmath"
)

type bettingAllIn struct {
	Betting
}

func newBettingAllIn(init_chip int) IBettingStrategy {
	b := &bettingAllIn{}
	b.init(init_chip)
	return b
}

// 初始化
func (b *bettingAllIn) init(init_chip int) {
	b.t = BETTING_TYPE.ALL_IN
	b.set_chip(init_chip)
}

// 查询下注额
func (b *bettingAllIn) Query_bet() (int, error) {
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

// 查询配置
func (b *bettingAllIn) Query_option() (int, int, BETTING_TYPE.TYPE) {
	return MIN_BET, MAX_BET, b.t
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
