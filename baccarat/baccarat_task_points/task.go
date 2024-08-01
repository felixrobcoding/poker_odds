/*
功能：百家乐任务-统计庄家/闲家点数分布
说明：
*/
package baccarat_task_points

import (
	"sync"

	"github.com/poker-x-studio/x/xdebug"
)

const (
	GO_ROUTINE_CNT = 10    //goroutine个数
	LOOP_TIMES     = 1000  //每个goroutine循环次数
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

	Init_point_times_map()

	for _, v := range shoe_stats {
		for i := 0; i < len(v.player_points); i++ {
			point := v.player_points[i]
			Player_point_times_map[point]++
		}
		for i := 0; i < len(v.dealer_points); i++ {
			point := v.dealer_points[i]
			Dealer_point_times_map[point]++
		}
	}

	//输出
	Outputer_point_times_map(Player_point_times_map, "闲家")
	Outputer_point_times_map(Dealer_point_times_map, "庄家")
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
