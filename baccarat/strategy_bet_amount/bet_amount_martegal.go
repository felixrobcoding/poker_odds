/*
功能：下注额策略-马丁格尔策略
说明：
*/
package strategy_bet_amount

import (
	"errors"
	"github.com/felixrobcoding/poker_oddscommon/BET_AMOUNT_STRATEGY"

	"github.com/poker-x-studio/x/xmath"
)

type betAmountMartegal struct {
	BetAmountStrategy
}

func newBetAmountMartegal(init_chip int) IBetAmountStrategy {
	b := &betAmountMartegal{}
	b.init(init_chip)
	return b
}

// 初始化
func (b *betAmountMartegal) init(init_chip int) {
	b.t = BET_AMOUNT_STRATEGY.MARTEGAL
	b.set_chip(init_chip)
}

// 查询下注额
// 输了就加倍
func (b *betAmountMartegal) Query_bet_amount() (int, error) {
	bet := MIN_BET

	len := len(b.feedback_nodes)
	if len <= 0 {
		bet = MIN_BET
	} else {
		last_node := b.feedback_nodes[len-1]

		if last_node.Result_score > 0 { //赢了
			bet = MIN_BET
		}
		if last_node.Result_score < 0 { //输了
			bet = 2 * last_node.Bet_amount
		}
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
