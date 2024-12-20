/*
功能：blackjack任务-dealer爆牌率
说明：
*/
package blackjack_task_dealer_bust_rate

import (
	"github.com/felixrobcoding/poker_oddsblackjack/define/CARD_TYPE"
	"sync"
)

const (
	GO_ROUTINE_CNT = 10    //goroutine个数
	LOOP_TIMES     = 10000 //每个goroutine循环次数
	is_output_svg  = true  //
)

var (
	wg         sync.WaitGroup
	mutex      sync.Mutex //
	shoe_stats []ShoeStat //每靴牌统计
	show_card  byte       //dealer明牌
	mutex_ex   sync.Mutex //
)

// 开启
func Start() {
	//xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_ENTER)
	//defer xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_EXIT)

	for show_card = 0x01; show_card <= 0x0D; show_card++ {
		mutex_ex.Lock()
		Start_ex()
		mutex_ex.Unlock()
	}
}

func Start_ex() {
	//xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_ENTER)
	//defer xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_EXIT)

	// var lifeTime xdebug.LifeTime
	// lifeTime.Start()
	// defer func() {
	// 	lifeTime.End()
	// }()

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
	//xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_ENTER)
	//defer xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_EXIT)

	//创建
	flow_control := NewFlowControl()

	{
		//洗牌
		flow_control.Shuffle()

		//dealer发牌[指定一张牌]
		flow_control.Round_begin_to_deal(show_card)

		//dealer持续发牌到停牌
		flow_control.Dealer_turn()
	}

	mutex.Lock()
	shoe_stats = append(shoe_stats, *flow_control.Extract_shoe_stat())
	mutex.Unlock()
}

// 统计
func stat() {
	//
	sum_deal_times := 0
	sum_dealer_bust := 0
	//show_card := byte(0)

	Init_point_times_map()
	for _, v := range shoe_stats {
		sum_deal_times++
		//xlog_entry.Infof("%d,shoe_stat:%s,", k, v.String())

		//点数次数统计
		Point_times_map[v.dealer_point]++

		//爆点次数
		//show_card = v.dealer_show_card
		if v.dealer_card_type == CARD_TYPE.BUST {
			sum_dealer_bust++
		}
	}

	//靴牌数
	//shoe_cnt := len(shoe_stats)

	//爆牌率
	//dealer_bust_ratio := float64(sum_dealer_bust) / float64(shoe_cnt)

	//xlog_entry.Infof("shoe_cnt:%d,show_card:%s,dealer_bust_ratio:%.2f%%,", shoe_cnt, common.Card_2_sign(show_card), dealer_bust_ratio*100)

	//输出
	Outputer_point_times_map()
}

// svg输出
func stat_point_times() {
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
