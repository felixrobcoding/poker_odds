/*
功能：大路统计
说明：
*/
package big_road

import (
	"fmt"
	"sort"
)

// 列统计
type ColStat struct {
	Hands_per_col int //每列手数
	Cols_cnt      int //列数
}

func (c *ColStat) String() string {
	return fmt.Sprintf("[Hands_per_col:%d,Cols_cnt:%d]", c.Hands_per_col, c.Cols_cnt)
}

// 大路统计
type BigRoadStat struct {
	Col_stats []ColStat //列统计
}

// 排序
func (b *BigRoadStat) sort() {
	sort.SliceStable(b.Col_stats, func(i, j int) bool {
		return b.Col_stats[i].Hands_per_col < b.Col_stats[j].Hands_per_col
	})
}

func (b *BigRoadStat) String() string {
	b.sort()
	txt := ""
	for _, v := range b.Col_stats {
		txt += fmt.Sprintf("[ColStat:%s]", v.String())
	}
	return txt
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
