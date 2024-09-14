/*
功能：下注额策略-马丁格尔N策略
说明：
马丁格尔N策略:连续输N把内,输了加倍,之后恢复原始注码,增加止损的概念
*/
package strategy_bet_amount

import (
	"Odds/common/BET_AMOUNT_STRATEGY"
	"errors"

	"github.com/poker-x-studio/x/xmath"
)

const (
	CONTINUE_LOSE_CNT = 4
)

type betAmountMartegalN struct {
	BetAmountStrategy
}

func newBetAmountMartegalN(init_chip int) IBetAmountStrategy {
	b := &betAmountMartegalN{}
	b.init(init_chip)
	return b
}

// 初始化
func (b *betAmountMartegalN) init(init_chip int) {
	b.t = BET_AMOUNT_STRATEGY.MARTEGAL_N
	b.set_chip(init_chip)
}

// 查询下注额
// 连续输N把内,输了加倍,之后恢复原始注码,增加止损的概念
func (b *betAmountMartegalN) Query_bet_amount() (int, error) {
	bet := MIN_BET

	len := len(b.feedback_nodes)
	if len <= 0 {
		bet = MIN_BET
	} else {
		last_node := b.feedback_nodes[len-1]
		if last_node.Result_score > 0 { //赢了
			bet = MIN_BET
		} else { //输了
			//连续输的节点数
			lose_node_cnt := 0
			for i := len - 1; i >= 0; i-- {
				if b.feedback_nodes[i].Result_score < 0 {
					lose_node_cnt++
				}
				if b.feedback_nodes[i].Result_score > 0 {
					break
				}
			}

			if lose_node_cnt <= CONTINUE_LOSE_CNT { //输少于等于N颗，倍投
				bet = 2 * last_node.Bet_amount
			} else { //连续输N颗，放弃倍投
				bet = MIN_BET
			}
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
