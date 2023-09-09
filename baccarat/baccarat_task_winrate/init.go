/*
功能：百家乐任务-统计庄家胜率
说明：
*/
package baccarat_task_winrate

import "github.com/poker-x-studio/x/xlog"

const (
	TASK_TAG = "baccarat_task_winrate" //任务名称
	DECKS    = 8                       //牌副数
)

var (
	xlog_entry = xlog.New_entry(TASK_TAG)
)

//-----------------------------------------------
//					the end
//-----------------------------------------------
