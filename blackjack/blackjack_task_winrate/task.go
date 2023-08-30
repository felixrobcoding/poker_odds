/*
功能：blackjack任务-胜率统计
说明：
*/
package blackjack_task_winrate

import (
	"Odds/blackjack/strategy/query_times"
	"Odds/common/BETTING_TYPE"
	"fmt"
	"os"
	"sort"
	"sync"

	"github.com/poker-x-studio/x/ximage"

	"github.com/poker-x-studio/x/xdebug"
)

const (
	GO_ROUTINE_CNT          = 50    //goroutine个数
	LOOP_TIMES              = 10000 //每个goroutine循环次数
	is_output_strategy_jpeg = false //
)

var wg sync.WaitGroup
var mutex sync.Mutex      //
var shoe_stats []ShoeStat //每靴牌统计

// 开启
func Start() {
	xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_ENTER)
	defer xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_EXIT)

	var lifeTime xdebug.LifeTime
	lifeTime.Start()
	defer func() {
		lifeTime.End()
	}()

	shoe_stats = make([]ShoeStat, 0)
	wg.Add(GO_ROUTINE_CNT)

	for i := 0; i < GO_ROUTINE_CNT; i++ {
		go run()
	}

	wg.Wait()

	//统计
	stat()
}

// 运行
func run() {
	for i := 0; i < LOOP_TIMES; i++ {
		run_unit()
	}
	wg.Done()
}

// 运行单元
func run_unit() {
	xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_ENTER)
	defer xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_EXIT)

	//创建
	flow_control := NewFlowControl()
	//洗牌
	flow_control.Shuffle()

	//循环
	for {
		//发牌
		err := flow_control.Round_begin_to_deal()
		if err != nil {
			flow_control.Game_over() //游戏结束
			break
		}

		//校验blackjack牌型
		err = flow_control.Check_blackjack()
		if err != nil {
			flow_control.Compare()   //比牌
			flow_control.Round_end() //本轮结束
			continue
		}

		//闲家操作
		err = flow_control.Player_turn()
		if err != nil {
			flow_control.Compare()   //比牌
			flow_control.Round_end() //本轮结束
			continue
		}

		//庄家操作
		_ = flow_control.Dealer_turn()
		flow_control.Compare()   //比牌
		flow_control.Round_end() //本轮结束
	}

	mutex.Lock()
	shoe_stats = append(shoe_stats, *flow_control.Extract_shoe_stat())
	mutex.Unlock()
}

// 统计
func stat() {
	//策略表统计
	stat_strategy()

	//
	min_bet := 0
	max_bet := 0
	betting_t := BETTING_TYPE.ERROR

	sum_deal_times := 0
	sum_hands := 0
	sum_bets := 0
	sum_player_lose_hands := 0 //闲家输手数之和
	sum_player_push_hands := 0 //闲家和手数之和
	sum_player_win_hands := 0  //闲家赢手数之和
	sum_profit := 0.0

	player_profits := make([]float64, 0)
	for _, v := range shoe_stats {
		min_bet = v.min_bet
		max_bet = v.max_bet
		betting_t = v.betting_t

		sum_deal_times += v.deal_times
		sum_bets += v.player_total_bets
		sum_player_lose_hands += v.player_lose_hands
		sum_player_push_hands += v.player_push_hands
		sum_player_win_hands += v.player_win_hands

		profit := v.profit()
		sum_profit += profit
		player_profits = append(player_profits, profit)
	}
	sum_hands = sum_player_lose_hands + sum_player_push_hands + sum_player_win_hands

	sort.SliceStable(player_profits, func(i, j int) bool {
		return player_profits[i] < player_profits[j]
	})

	for k, v := range shoe_stats {
		xlog_entry.Infof("%d,shoe_stat:%s,", k, v.String())
	}

	//靴牌数
	shoe_cnt := len(shoe_stats)
	//每靴牌手数
	hands_per_shoe := float64(sum_hands) / float64(shoe_cnt)
	//每靴牌利润
	profit_per_shoe := sum_profit / float64(shoe_cnt)

	//每手利润
	profit_per_hand := sum_profit / float64(sum_hands)
	//利润投注比
	porfit_bet_ratio := sum_profit / float64(sum_bets)

	player_lose_hands_ratio := float64(sum_player_lose_hands) / float64(sum_hands)
	player_push_hands_ratio := float64(sum_player_push_hands) / float64(sum_hands)
	player_win_hands_ratio := float64(sum_player_win_hands) / float64(sum_hands)

	xlog_entry.Infof("min_bet:%d,max_bet:%d,betting_strategy:%s,player_min_profit:%.2f,player_max_profit:%.2f,", min_bet, max_bet, betting_t.String(), player_profits[0], player_profits[len(player_profits)-1])
	xlog_entry.Infof("sum_deal_times:%d,sum_hands:%d,sum_bets:%d,sum_profit:%.2f,hands_per_shoe:%.2f,profit_per_shoe:%.4f,profit_per_hand:%.4f,sum_profit/sum_bets:%.4f%", sum_deal_times, sum_hands, sum_bets, sum_profit, hands_per_shoe, profit_per_shoe, profit_per_hand, porfit_bet_ratio)
	xlog_entry.Infof("sum_player_lose_hands:%d,sum_player_push_hands:%d,sum_player_win_hands:%d,player_lose_hands_ratio:%.4f%%,player_push_hands_ratio:%.4f%%,player_win_hands_ratio:%.4f%%,", sum_player_lose_hands, sum_player_push_hands, sum_player_win_hands, player_lose_hands_ratio*100, player_push_hands_ratio*100, player_win_hands_ratio*100)
}

// 策略表统计
func stat_strategy() {
	if !is_output_strategy_jpeg {
		return
	}

	//查询次数
	svg_content := query_times.Instance().Output_svg_query_times()
	jpeg_filepath, svg_filepath, err := ximage.Svg_2_jpeg(svg_content)
	if err != nil {
		return
	}
	os.Remove(svg_filepath)
	fmt.Println(jpeg_filepath)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
