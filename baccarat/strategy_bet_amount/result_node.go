/*
功能：下注额策略-结果节点
说明：
*/
package strategy_bet_amount

import (
	"sort"
)

type ResultNode struct {
	Current_chip   float64   //当前筹码
	Current_bets   []int     //当前下注[一轮中有分牌情况,所以存在多个]
	Current_scores []float64 //当前得分
}

// 最小下注
func (r ResultNode) Min_bet() int {
	len := len(r.Current_bets)
	if len == 0 {
		return 0
	} else if len == 1 {
		return r.Current_bets[0]
	} else {
		sort.SliceStable(r.Current_bets, func(i, j int) bool {
			return r.Current_bets[i] < r.Current_bets[j]
		})
		return r.Current_bets[0]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
