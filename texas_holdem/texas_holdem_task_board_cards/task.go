/*
功能：德州扑克任务-统计公共牌
说明：适用于任何五张牌的概率
*/
package texas_holdem_task_board_cards

import (
	"Odds/common/algorithm"
	"fmt"
	"os"

	"github.com/poker-x-studio/x/ximage"

	"Odds/texas_holdem/define/CARD_TYPE"
	"Odds/texas_holdem/logic"
	"Odds/texas_holdem/svg_utils"

	"sync"

	"github.com/poker-x-studio/x/xdebug"
)

const (
	GO_ROUTINE_CNT = 1 //goroutine个数
	LOOP_TIMES     = 1 //每个goroutine循环次数
	MIN_CARD_CNT   = 2
	MAX_CARD_CNT   = 5
)

var (
	wg            sync.WaitGroup
	mutex         sync.Mutex
	analyse_items []logic.AnalyseItem
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

	analyse_items = make([]logic.AnalyseItem, 0)
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

	for card_cnt := MIN_CARD_CNT; card_cnt <= MAX_CARD_CNT; card_cnt++ {

		shoe_cards := algorithm.Shuffle_cards(DECKS)
		board_cards := shoe_cards[:card_cnt]
		analyse_item := logic.Analyse(board_cards)

		mutex.Lock()
		analyse_items = append(analyse_items, *analyse_item)
		mutex.Unlock()
	}
}

// 统计
func stat() {
	arr_cnt := MAX_CARD_CNT - MIN_CARD_CNT + 1
	stat_items := make([][]logic.AnalyseItem, arr_cnt)
	for i := 0; i < arr_cnt; i++ {
		stat_items[i] = make([]logic.AnalyseItem, 0)
	}

	//分类
	for _, v := range analyse_items {
		stat_items[v.Card_cnt()-MIN_CARD_CNT] = append(stat_items[v.Card_cnt()-MIN_CARD_CNT], v)
	}

	//是否需要输出svg
	is_output_svg := true
	svg_card_type_stats := make([]*svg_utils.CardTypeStat, 0)

	//统计
	for card_cnt := MIN_CARD_CNT; card_cnt <= MAX_CARD_CNT; card_cnt++ {

		total_item_cnt := len(stat_items[card_cnt-MIN_CARD_CNT])
		card_cnt := stat_items[card_cnt-MIN_CARD_CNT][0].Card_cnt()
		card_types := CARD_TYPE.Card_types()
		card_type_cnts := make([]int, len(card_types))

		for _, v := range stat_items[card_cnt-MIN_CARD_CNT] {
			for k1, v1 := range card_types {

				for _, v2 := range v.Card_types {
					if v2 == v1 {
						card_type_cnts[k1]++
					}
				}

			}
			//输出所有牌型信息
			xlog_entry.Tracef("%s", v.String(true))
		}

		//日志输出
		for k1, v1 := range card_types {
			ratio := float64(card_type_cnts[k1]) / float64(total_item_cnt*1.0)
			xlog_entry.Infof("%d,total_item_cnt:%d,牌张数:%d, 牌型:%s,个数:%d,百分比:%.4f%%", k1, total_item_cnt, card_cnt, v1, card_type_cnts[k1], ratio*100.0)

			if is_output_svg {
				svg_card_type_stat := &svg_utils.CardTypeStat{
					Total_run_times: total_item_cnt,
					Deal_card_cnt:   card_cnt,
					Type:            v1,
					Type_cnt:        card_type_cnts[k1],
					Percentage:      ratio,
				}
				svg_card_type_stats = append(svg_card_type_stats, svg_card_type_stat)
			}
		}
	}

	if is_output_svg {
		svg_content := svg_utils.Instance_card_type_stat_svg().Make_svg(svg_utils.TXT_HEADER_TEXAS_HOLDEM, svg_card_type_stats)
		fmt.Println(svg_content)

		png_filepath, svg_filepath, err := ximage.Svg_2_png(svg_content)
		if err != nil {
			return
		}
		os.Remove(svg_filepath)
		fmt.Println(png_filepath)
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
