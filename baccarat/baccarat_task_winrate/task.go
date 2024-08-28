/*
功能：百家乐任务-统计庄家胜率
说明：
*/
package baccarat_task_winrate

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/svg_utils"
	"Odds/common/BET_AMOUNT_STRATEGY"
	"fmt"
	"os"
	"sort"
	"sync"

	"github.com/poker-x-studio/x/ximage"

	"github.com/poker-x-studio/x/xdebug"
)

const (
	GO_ROUTINE_CNT = 1     //goroutine个数
	LOOP_TIMES     = 1     //每个goroutine循环次数
	is_output_jpeg = false //
)

var (
	wg         sync.WaitGroup
	mutex      sync.Mutex //
	shoe_stats []ShoeStat //每靴牌统计
)

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

		//校验Natural牌型
		err = flow_control.Check_natural()
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
	//获胜区域统计
	stat_win_bet_area_percentage()
	//统计大路
	stat_bigroad()

	//
	min_bet := 0
	max_bet := 0
	bet_amount_strategy := BET_AMOUNT_STRATEGY.ERROR

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
		bet_amount_strategy = v.bet_amount_strategy

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

	//测试输出
	// for k, v := range shoe_stats {
	// 	xlog_entry.Tracef("%d,shoe_stat:%s,", k, v.String())
	// }

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

	xlog_entry.Tracef("min_bet:%d,max_bet:%d,bet_amount_strategy:%s,player_min_profit:%.2f,player_max_profit:%.2f,", min_bet, max_bet, bet_amount_strategy.String(), player_profits[0], player_profits[len(player_profits)-1])
	xlog_entry.Tracef("sum_deal_times:%d,sum_hands:%d,sum_bets:%d,sum_profit:%.2f,hands_per_shoe:%.2f,profit_per_shoe:%.4f,profit_per_hand:%.4f,sum_profit/sum_bets:%.4f", sum_deal_times, sum_hands, sum_bets, sum_profit, hands_per_shoe, profit_per_shoe, profit_per_hand, porfit_bet_ratio)
	xlog_entry.Tracef("sum_player_lose_hands:%d,sum_player_push_hands:%d,sum_player_win_hands:%d,player_lose_hands_ratio:%.4f%%,player_push_hands_ratio:%.4f%%,player_win_hands_ratio:%.4f%%,", sum_player_lose_hands, sum_player_push_hands, sum_player_win_hands, player_lose_hands_ratio*100, player_push_hands_ratio*100, player_win_hands_ratio*100)
}

// 获胜下注区域统计
func stat_win_bet_area_percentage() {
	sum_hands := 0

	//获胜区域统计
	win_bet_area_map := make(map[BET_AREA.TYPE]int)
	for _, v := range shoe_stats {
		for _, v1 := range v.win_bet_areas {
			for _, v2 := range v1 {
				win_bet_area_map[v2]++
			}
			sum_hands++
		}
	}
	win_bet_area_stats := make([]*svg_utils.WinBetAreaStat, 0)
	win_bet_areas := BET_AREA.All_bet_area()
	for _, v := range win_bet_areas {
		cnt, ok := win_bet_area_map[v]
		if ok {
			stat := &svg_utils.WinBetAreaStat{
				Total_hands:      sum_hands,
				Win_bet_area:     v,
				Win_bet_area_cnt: cnt,
				Percentage:       float64(cnt) / float64(sum_hands),
			}
			win_bet_area_stats = append(win_bet_area_stats, stat)
		}
	}

	for _, v := range win_bet_area_stats {
		xlog_entry.Infof("%s", v.String())
	}

	if is_output_jpeg {
		svg_content := svg_utils.Instance_win_bet_area_stat_svg().Make_svg(win_bet_area_stats)
		fmt.Println(svg_content)

		jpeg_filepath, svg_filepath, err := ximage.Svg_2_jpeg(svg_content)
		if err != nil {
			return
		}
		os.Remove(svg_filepath)
		fmt.Println(jpeg_filepath)
	}
}

// 统计大路
func stat_bigroad() {
	type bigroad_item struct {
		hands_per_col int
		cols_cnt      int
		percentage    float64
	}
	sum_hands := 0
	sum_cols := 0
	col_stat_map := make(map[int]int, 0)
	for _, v := range shoe_stats {
		for _, v1 := range v.bigroad_stat.Col_stats {
			col_stat_map[v1.Hands_per_col] += v1.Cols_cnt
			sum_hands += v1.Hands_per_col * v1.Cols_cnt
			sum_cols += v1.Cols_cnt
		}
	}

	col_stat_percentage_map := make(map[int]bigroad_item, 0)
	for k, v := range col_stat_map {
		col_stat_percentage_map[k] = bigroad_item{k, v, float64(v) / (float64(sum_cols) * 1.0)}
	}

	col_stats := make([]bigroad_item, 0)
	for _, v := range col_stat_percentage_map {
		col_stats = append(col_stats, bigroad_item{v.hands_per_col, v.cols_cnt, v.percentage})
	}

	//排序
	sort.SliceStable(col_stats, func(i, j int) bool {
		return col_stats[i].hands_per_col < col_stats[j].hands_per_col
	})

	for _, v := range col_stats {
		xlog_entry.Infof("sum_hands:%d,sum_cols:%d,hands_per_col:%d,cols_cnt:%d,percentage::%.4f%%,", sum_hands, sum_cols, v.hands_per_col, v.cols_cnt, v.percentage*100)
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
