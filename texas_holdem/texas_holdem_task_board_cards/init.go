/*
功能：德州扑克任务-统计公共牌
说明：适用于任何五张牌的概率
*/
package texas_holdem_task_board_cards

import "github.com/poker-x-studio/x/xlog"

const (
	TASK_TAG = "holdem_task_board_cards" //任务名称
	DECKS    = 1                         //牌副数
)

var (
	xlog_entry = xlog.New_entry(TASK_TAG)
	is_debug   = true
)

//-----------------------------------------------
//					the end
//-----------------------------------------------
