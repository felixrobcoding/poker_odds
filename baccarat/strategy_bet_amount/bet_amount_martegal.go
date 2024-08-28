/*
功能：下注额策略-马丁格尔策略
说明：
*/
package strategy_bet_amount

import (
	"Odds/common/BET_AMOUNT_STRATEGY"
	"errors"

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

	len := len(b.result_nodes)
	if len <= 0 {
		bet = MIN_BET
	} else {
		last_node := b.result_nodes[len-1]
		last_sum_score := 0.0
		for _, v := range last_node.Current_scores {
			last_sum_score += v
		}

		if last_sum_score > 0 { //赢了
			bet = MIN_BET
		}
		if last_sum_score < 0 { //输了
			bet = xmath.Min(MAX_BET, 2*last_node.Min_bet())
		}
	}

	if b.Is_enough_money(bet) {
		return bet, nil
	}
	return 0, errors.New(NOT_ENOUGH_MONEY)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
