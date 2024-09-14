/*
功能：下注额策略-斐波那契策略
说明：
*/
package strategy_bet_amount

import (
	"Odds/common/BET_AMOUNT_STRATEGY"
	"errors"

	"github.com/poker-x-studio/x/xmath"
)

type betAmountFibonacci struct {
	BetAmountStrategy
}

func newBetAmountFibonacci(init_chip int) IBetAmountStrategy {
	b := &betAmountFibonacci{}
	b.init(init_chip)
	return b
}

// 初始化
func (b *betAmountFibonacci) init(init_chip int) {
	b.t = BET_AMOUNT_STRATEGY.FIBONACCI
	b.set_chip(init_chip)
}

// 斐波那契函数
func fibonacci(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 查询下注额
// 斐波那契数就是由之前的两数相加而得出
func (b *betAmountFibonacci) Query_bet_amount() (int, error) {
	bet := MIN_BET

	len := len(b.feedback_nodes)
	if len < 1 {
		bet = MIN_BET
	} else {
		last_node := b.feedback_nodes[len-1]
		if last_node.Result_score > 0 { //赢了
			bet = MIN_BET
		} else { //输了
			lose_cnt := 0
			for i := len - 1; i >= 0; i-- {
				if b.feedback_nodes[i].Result_score < 0 {
					lose_cnt++
				} else {
					break
				}
			}

			bet = MIN_BET * fibonacci(lose_cnt)
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
