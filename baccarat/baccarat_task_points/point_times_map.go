/*
功能：百家乐任务-统计庄家/闲家点数分布
说明：点数次数map
*/
package baccarat_task_points

import (
	"fmt"
	"sort"
)

var (
	Player_point_times_map = make(map[int]int, 0) //key为点数,value为次数
	Dealer_point_times_map = make(map[int]int, 0) //key为点数,value为次数
)

// 初始化
func Init_point_times_map() {
	Player_point_times_map = make(map[int]int, 0)
	Dealer_point_times_map = make(map[int]int, 0)
}

// 输出
func Outputer_point_times_map(Point_times_map map[int]int, title string) {
	sum_cnt := 0
	for _, v := range Point_times_map {
		sum_cnt += v
	}

	//key排序[点数]
	var keys []int
	for k, _ := range Point_times_map {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	//输出html表格
	{
		fmt.Printf("<h1>Baccarat,8副牌[8*52=416张],点数次数统计,总运行次数1w次</h1>")

		fmt.Printf("<table><tr><th>%s点数</th><th>出现次数</th><th>百分比</th><th>总手数</th></tr>", title)
		for k, v := range keys {
			times := Point_times_map[v]
			ratio := float64(times) / float64(sum_cnt)

			if k == 0 {
				fmt.Printf("<tr><td>%d</td><td>%d</td><td>%.2f%%</td><td rowspan=\"10\">%d</td></tr></tr>\r\n", v, times, ratio*100, sum_cnt)
			} else {
				fmt.Printf("<tr><td>%d</td><td>%d</td><td>%.2f%%</td></tr></tr>\r\n", v, times, ratio*100)
			}
		}
		fmt.Println("</table>")
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
