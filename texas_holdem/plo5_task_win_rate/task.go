/*
功能：5张奥马哈任务-计算胜率
说明：
*/
package plo5_task_win_rate

import (
	"Odds/common/algorithm"
	"bufio"
	"fmt"
	"os"

	"Odds/texas_holdem/define"
	"Odds/texas_holdem/logic"
	"Odds/texas_holdem/logic_plo"

	"sync"

	"github.com/poker-x-studio/x/xdebug"
)

const (
	GO_ROUTINE_CNT = 100 //goroutine个数
	LOOP_TIMES     = 100 //每个goroutine循环次数
)

var (
	wg                 sync.WaitGroup
	stat_plo5_win_rate StatPlo5WinRate
)

// 开启
func Start() {
	xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_ENTER)
	defer xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_EXIT)

	is_input := true
	if is_input { // 接收用户输入
		var err error
		stat_plo5_win_rate.play1_hole_cards, stat_plo5_win_rate.play2_hole_cards, err = user_input()
		if err != nil {
			xlog_entry.Errorf("Start(),exit")
			panic("error")
		}
	} else { //直接赋值调试
		stat_plo5_win_rate.play1_hole_cards = []byte{0x01, 0x02, 0x03, 0x04, 0x05}
		stat_plo5_win_rate.play2_hole_cards = []byte{0x11, 0x12, 0x13, 0x14, 0x15}
	}

	var lifeTime xdebug.LifeTime
	lifeTime.Start()
	defer func() {
		lifeTime.End()
	}()

	wg.Add(GO_ROUTINE_CNT)

	for i := 0; i < GO_ROUTINE_CNT; i++ {
		go run()
	}

	wg.Wait()

	//统计
	stat()
}

// 接收用户输入
func user_input() (player1_hole_cards []byte, player2_hole_cards []byte, err error) {
	player1_hole_cards = make([]byte, 5)
	player2_hole_cards = make([]byte, 5)
	err = nil

	stdin := bufio.NewReader(os.Stdin)
	fmt.Println("输入玩家1的5张手牌[16进制],使用空格分开[比如,0x01 0x12 0x23 0x34 0x05]:")
	fmt.Fscan(stdin, &player1_hole_cards[0], &player1_hole_cards[1], &player1_hole_cards[2], &player1_hole_cards[3], &player1_hole_cards[4])

	stdin = bufio.NewReader(os.Stdin)
	fmt.Println("输入玩家2的5张手牌[16进制],使用空格分开[比如,0x01 0x12 0x23 0x34 0x05]:")
	fmt.Fscan(stdin, &player2_hole_cards[0], &player2_hole_cards[1], &player2_hole_cards[2], &player2_hole_cards[3], &player2_hole_cards[4])

	return
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
	// xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_ENTER)
	// defer xlog_entry.Tracef("%s,%s", xdebug.Funcname(), xdebug.FUNC_EXIT)

	//洗牌
	shoe_cards := algorithm.Shuffle_cards(DECKS)

	//玩家1手牌
	removed_cnt, remain_shoe_cards := algorithm.Removes(shoe_cards, stat_plo5_win_rate.play1_hole_cards)
	if removed_cnt != len(stat_plo5_win_rate.play1_hole_cards) {
		xlog_entry.Errorf("run_unit(),removed_cnt:%d,", removed_cnt)
		panic("error")
	}

	//玩家2手牌
	removed_cnt, remain_shoe_cards = algorithm.Removes(remain_shoe_cards, stat_plo5_win_rate.play2_hole_cards)
	if removed_cnt != len(stat_plo5_win_rate.play2_hole_cards) {
		xlog_entry.Errorf("run_unit(),removed_cnt:%d,", removed_cnt)
		panic("error")
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

	// str_play1_hole_cards := common.Cards_2_sign(stat_plo5_win_rate.play1_hole_cards)
	// str_play2_hole_cards := common.Cards_2_sign(stat_plo5_win_rate.play2_hole_cards)
	// str_board_cards := common.Cards_2_sign(board_cards)
	// xlog_entry.Tracef("str_play1_hole_cards:%s", str_play1_hole_cards)
	// xlog_entry.Tracef("str_play2_hole_cards:%s", str_play2_hole_cards)
	// xlog_entry.Tracef("str_board_cards:%s", str_board_cards)

	// str_play1_compare_cards := common.Cards_2_sign(play1_item.Compare_cards())
	// str_play2_compare_cards := common.Cards_2_sign(play2_item.Compare_cards())
	// xlog_entry.Tracef("str_play1_compare_cards:%s", str_play1_compare_cards)
	// xlog_entry.Tracef("str_play2_compare_cards:%s", str_play2_compare_cards)

	// xlog_entry.Tracef("play1_item:%s", play1_item.String(false))
	// xlog_entry.Tracef("play2_item:%s", play2_item.String(false))
}

// 统计
func stat() {
	xlog_entry.Infof("stat_plo5_win_rate:%s", stat_plo5_win_rate.String())
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
