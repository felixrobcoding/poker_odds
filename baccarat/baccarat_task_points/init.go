/*
功能：百家乐任务-统计庄家/闲家点数分布
说明：
*/
package baccarat_task_points

import "github.com/poker-x-studio/x/xlog"

const (
	TASK_TAG = "baccarat_task_points" //任务名称
	DECKS    = 8                      //牌副数
)

var (
	xlog_entry = xlog.New_entry(TASK_TAG)
)

//-----------------------------------------------
//					the end
//-----------------------------------------------
