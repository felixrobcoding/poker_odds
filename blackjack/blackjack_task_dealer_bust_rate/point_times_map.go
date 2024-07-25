/*
功能：blackjack任务-dealer爆牌率
说明：点数次数map
*/
package blackjack_task_dealer_bust_rate

import (
	"Odds/blackjack/define"
	"Odds/common"
	"fmt"
	"sort"
)

var (
	Point_times_map = make(map[int]int, 0) //key为点数,value为次数
)

// 初始化
func Init_point_times_map() {
	Point_times_map = make(map[int]int, 0)
}

// 输出
func Outputer_point_times_map() {
	sum_shoe_cnt := 0
	sum_shoe_cnt_bust := 0
	for k, v := range Point_times_map {
		sum_shoe_cnt += v
		if k >= define.POINT_BUST {
			sum_shoe_cnt_bust += v
		}
	}

	//key排序[点数]
	var keys []int
	for k, _ := range Point_times_map {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// for _, v := range keys {
	// 	times := Point_times_map[v]
	// 	ratio := float64(times) / float64(sum_shoe_cnt)
	// 	xlog_entry.Infof("sum_shoe_cnt:%d,point:%d,times:%d,ratio:%.2f%%,", sum_shoe_cnt, v, times, ratio*100)
	// }
	bust_ratio := float64(sum_shoe_cnt_bust) / float64(sum_shoe_cnt)
	//xlog_entry.Infof("sum_shoe_cnt:%d,sum_shoe_cnt_bust:%d,bust_ratio:%.2f%%,", sum_shoe_cnt, sum_shoe_cnt_bust, bust_ratio*100)

	//输出html表格
	{
		fmt.Printf("<h1>Blackjack,4副牌[4*52=208张],dealer明牌为%s,总运行次数10w次</h1>", common.Card_2_sign(show_card))

		fmt.Println("<table><tr><th>dealer停牌时点数</th><th>出现次数</th><th>百分比</th><th>爆牌率</th></tr>")
		for _, v := range keys {
			times := Point_times_map[v]
			ratio := float64(times) / float64(sum_shoe_cnt)
			if v == 17 {
				fmt.Printf("<tr><td>%d</td><td>%d</td><td>%.2f%%</td><td rowspan=\"5\"></td></tr>\r\n", v, times, ratio*100)
			} else if v == 22 {
				fmt.Printf("<tr><td>%d</td><td>%d</td><td>%.2f%%</td><td rowspan=\"5\">%.2f%%</td></tr>\r\n", v, times, ratio*100, bust_ratio*100)
			} else {
				fmt.Printf("<tr><td>%d</td><td>%d</td><td>%.2f%%</td></tr>\r\n", v, times, ratio*100)
			}
		}
		fmt.Println("</table>")
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
