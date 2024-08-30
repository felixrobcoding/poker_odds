/*
功能：反馈节点
说明：
*/
package suggestion

import (
	"Odds/baccarat/define/BET_AREA"
)

// 反馈节点
type FeedbackNode struct {
	Current_chip       float64       //当前-筹码
	Current_bet_area   BET_AREA.TYPE //当前-下注区域
	Current_bet_amount int           //当前-下注额
	Result_area        BET_AREA.TYPE //结果-获胜区域
	Result_score       float64       //结果-得分
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
