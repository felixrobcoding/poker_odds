/*
功能：下注策略-马丁格尔策略
说明：
*/
package strategy_betting

import (
	"Odds/common/BETTING_TYPE"
	"errors"

	"github.com/poker-x-studio/x/xmath"
)

type bettingMartegal struct {
	Betting
}

func newBettingMartegal(init_chip int) IBettingStrategy {
	b := &bettingMartegal{}
	b.init(init_chip)
	return b
}

// 初始化
func (b *bettingMartegal) init(init_chip int) {
	b.t = BETTING_TYPE.MARTEGAL
	b.set_chip(init_chip)
}

// 查询下注额
// 输了就加倍
func (b *bettingMartegal) Query_bet() (int, error) {
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
