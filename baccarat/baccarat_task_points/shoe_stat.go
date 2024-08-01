/*
功能：百家乐任务-统计庄家/闲家点数分布
说明：
*/
package baccarat_task_points

import (
	"fmt"
)

// 每靴牌统计
type ShoeStat struct {
	shoe_index    int
	deal_times    int   //发牌次数
	player_points []int //闲家点数
	dealer_points []int //庄家点数
}

func (s *ShoeStat) String() string {
	str := fmt.Sprintf("[shoe_index:%d", s.shoe_index)
	str += fmt.Sprintf("]")
	return str
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
