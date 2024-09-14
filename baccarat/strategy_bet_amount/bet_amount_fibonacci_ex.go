/*
功能：下注额策略-斐波那契扩展策略
说明：
斐波那契扩展策略:下注额为之前所有输的筹码的总和
*/
package strategy_bet_amount

import (
	"Odds/common/BET_AMOUNT_STRATEGY"
	"errors"

	"github.com/poker-x-studio/x/xmath"
)

type betAmountFibonacciEx struct {
	BetAmountStrategy
}

func newBetAmountFibonacciEx(init_chip int) IBetAmountStrategy {
	b := &betAmountFibonacciEx{}
	b.init(init_chip)
	return b
}

// 初始化
func (b *betAmountFibonacciEx) init(init_chip int) {
	b.t = BET_AMOUNT_STRATEGY.FIBONACCI_EX
	b.set_chip(init_chip)
}

// 查询下注额
// 下注额为之前所有输的筹码的总和
func (b *betAmountFibonacciEx) Query_bet_amount() (int, error) {
	bet := MIN_BET

	len := len(b.feedback_nodes)
	if len < 1 {
		bet = MIN_BET
	} else {
		total_bet_amount := 0
		last_node := b.feedback_nodes[len-1]
		if last_node.Result_score > 0 { //赢了
			bet = MIN_BET
		} else { //输了
			for i := len - 1; i >= 0; i-- {
				if b.feedback_nodes[i].Result_score < 0 {
					total_bet_amount += b.feedback_nodes[i].Bet_amount
				} else {
					break
				}
			}

			bet = total_bet_amount
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
