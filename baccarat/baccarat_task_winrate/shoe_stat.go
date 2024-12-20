/*
功能：百家乐任务-统计庄家胜率
说明：
*/
package baccarat_task_winrate

import (
	"fmt"
	"github.com/felixrobcoding/poker_oddsbaccarat/define/BET_AREA"
	"github.com/felixrobcoding/poker_oddsbaccarat/strategy_bet_area/big_road"
	"github.com/felixrobcoding/poker_oddscommon/BET_AMOUNT_STRATEGY"
)

// 每靴牌统计
type ShoeStat struct {
	shoe_index          int
	option_min_bet      int                      //
	option_max_bet      int                      //
	bet_amount_strategy BET_AMOUNT_STRATEGY.TYPE //下注额策略
	deal_times          int                      //发牌次数
	max_bet_amount      int                      //实际最大下注额
	player_init_chip    float64                  //闲家起始筹码
	player_chip         float64                  //闲家筹码
	player_total_bets   int                      //闲家总投注
	player_lose_hands   int                      //闲家输的手数
	player_push_hands   int                      //闲家和的手数
	player_win_hands    int                      //闲家赢的手数
	win_bet_areas       [][]BET_AREA.TYPE        //获胜区域
	bigroad_stat        *big_road.BigRoadStat    //大路统计
}

func (s *ShoeStat) String() string {
	str := fmt.Sprintf("[shoe_index:%d,option_min_bet:%d,option_max_bet:%d,bet_amount_strategy:%s,", s.shoe_index, s.option_min_bet, s.option_max_bet, s.bet_amount_strategy.String())
	str += fmt.Sprintf("bigroad_stat:%s,", s.bigroad_stat.String())
	str += fmt.Sprintf("deal_times:%d,max_bet_amount:%d,player_init_chip:%.2f,player_chip:%.2f,player_total_bets:%d,profit:%.2f,", s.deal_times, s.max_bet_amount, s.player_init_chip, s.player_chip, s.player_total_bets, s.profit())
	str += fmt.Sprintf("total_hands:%d,player_lose_hands:%d,player_push_hands:%d,player_win_hands:%d]", s.total_hands(), s.player_lose_hands, s.player_push_hands, s.player_win_hands)
	return str
}

func (s *ShoeStat) profit() float64 {
	return s.player_chip - s.player_init_chip
}

func (s *ShoeStat) total_hands() int {
	return s.player_lose_hands + s.player_push_hands + s.player_win_hands
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
