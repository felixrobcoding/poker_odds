/*
功能：5张奥马哈任务-计算胜率
说明：
*/
package plo5_task_win_rate

import (
	"Odds/common"
	"Odds/common/algorithm"

	"Odds/texas_holdem/define"
	"Odds/texas_holdem/logic"
	"Odds/texas_holdem/logic_plo"

	"sync"

	"github.com/poker-x-studio/x/xdebug"
)

const (
	GO_ROUTINE_CNT = 10  //goroutine个数
	LOOP_TIMES     = 600 //每个goroutine循环次数
)

var (
	wg                 sync.WaitGroup
	stat_plo5_win_rate StatPlo5WinRate
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

	//设置手牌
	stat_plo5_win_rate.play1_hole_cards = []byte{0x3B, 0x24, 0x33, 0x35, 0x08}
	stat_plo5_win_rate.play2_hole_cards = []byte{0x13, 0x1A, 0x2C, 0x2A, 0x27}

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

	//洗牌
	shoe_cards := algorithm.Shuffle_cards(DECKS)

	//玩家1手牌
	removed_cnt, remain_shoe_cards := algorithm.Removes(shoe_cards, stat_plo5_win_rate.play1_hole_cards)
	if removed_cnt != len(stat_plo5_win_rate.play1_hole_cards) {
		xlog_entry.Errorf("run_unit(),removed_cnt:%d,", removed_cnt)
		return
	}

	//玩家2手牌
	removed_cnt, remain_shoe_cards = algorithm.Removes(remain_shoe_cards, stat_plo5_win_rate.play2_hole_cards)
	if removed_cnt != len(stat_plo5_win_rate.play2_hole_cards) {
		xlog_entry.Errorf("run_unit(),removed_cnt:%d,", removed_cnt)
		return
	}

	//公共牌
	start_index := 0
	end_index := start_index + define.BOARD_CNT
	board_cards := remain_shoe_cards[start_index:end_index]

	//组合最佳牌型
	play1_item := logic_plo.Combo_best_card_type(stat_plo5_win_rate.play1_hole_cards, board_cards)
	play2_item := logic_plo.Combo_best_card_type(stat_plo5_win_rate.play2_hole_cards, board_cards)

	//比较
	ct := logic.Compare(play1_item, play2_item)
	stat_plo5_win_rate.Set_result(ct)

	str_play1_hole_cards := common.Cards_2_sign(stat_plo5_win_rate.play1_hole_cards)
	str_play2_hole_cards := common.Cards_2_sign(stat_plo5_win_rate.play2_hole_cards)
	str_board_cards := common.Cards_2_sign(board_cards)
	xlog_entry.Tracef("str_play1_hole_cards:%s", str_play1_hole_cards)
	xlog_entry.Tracef("str_play2_hole_cards:%s", str_play2_hole_cards)
	xlog_entry.Tracef("str_board_cards:%s", str_board_cards)

	str_play1_compare_cards := common.Cards_2_sign(play1_item.Compare_cards())
	str_play2_compare_cards := common.Cards_2_sign(play2_item.Compare_cards())
	xlog_entry.Tracef("str_play1_compare_cards:%s", str_play1_compare_cards)
	xlog_entry.Tracef("str_play2_compare_cards:%s", str_play2_compare_cards)

	xlog_entry.Tracef("play1_item:%s", play1_item.String(false))
	xlog_entry.Tracef("play2_item:%s", play2_item.String(false))
}

// 统计
func stat() {
	xlog_entry.Infof("stat_plo5_win_rate:%s", stat_plo5_win_rate.String())
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
