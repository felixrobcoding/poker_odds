/*
功能：blackjack任务-胜率统计
说明：每靴牌统计
*/
package blackjack_task_winrate

import (
	"Odds/common/BETTING_TYPE"
	"fmt"
)

// 每靴牌统计
type ShoeStat struct {
	shoe_index        int
	min_bet           int               //
	max_bet           int               //
	betting_t         BETTING_TYPE.TYPE //下注策略
	deal_times        int               //发牌次数
	player_init_chip  float64           //闲家起始筹码
	player_chip       float64           //闲家筹码
	player_total_bets int               //闲家总投注
	player_lose_hands int               //闲家输的手数
	player_push_hands int               //闲家和的手数
	player_win_hands  int               //闲家赢的手数
}

func (s *ShoeStat) String() string {
	str := fmt.Sprintf("[shoe_index:%d,min_bet:%d,max_bet:%d,betting_strategy:%s,", s.shoe_index, s.min_bet, s.max_bet, s.betting_t.String())
	str += fmt.Sprintf("deal_times:%d,player_init_chip:%.2f,player_chip:%.2f,player_total_bets:%d,profit:%.2f,", s.deal_times, s.player_init_chip, s.player_chip, s.player_total_bets, s.profit())
	str += fmt.Sprintf("total_hands:%d,player_lose_hands:%d,player_push_hands:%d,player_win_hands:%d]", s.total_hands(), s.player_lose_hands, s.player_push_hands, s.player_win_hands)
	return str
}

// profit 盈利
func (s *ShoeStat) profit() float64 {
	return s.player_chip - s.player_init_chip
}

// total_hands 总手数
func (s *ShoeStat) total_hands() int {
	return s.player_lose_hands + s.player_push_hands + s.player_win_hands
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
