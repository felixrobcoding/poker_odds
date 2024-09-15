/*
功能：百家乐任务-统计庄家胜率
说明：
*/
package baccarat_task_winrate

import "sort"

type bigroad_col_stat struct {
	hands_per_col int
	cols_cnt      int
	percentage    float64
}

// 统计大路-每列多少节点数
func col_stat(shoe_stats []ShoeStat) {
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

	col_stat_percentage_map := make(map[int]bigroad_col_stat, 0)
	for k, v := range col_stat_map {
		col_stat_percentage_map[k] = bigroad_col_stat{k, v, float64(v) / (float64(sum_cols) * 1.0)}
	}

	col_stats := make([]bigroad_col_stat, 0)
	for _, v := range col_stat_percentage_map {
		col_stats = append(col_stats, bigroad_col_stat{v.hands_per_col, v.cols_cnt, v.percentage})
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
