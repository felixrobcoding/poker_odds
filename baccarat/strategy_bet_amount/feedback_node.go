/*
功能：下注额策略-反馈节点
说明：通过反馈节点 矫正 下注额
*/
package strategy_bet_amount

import (
	"github.com/felixrobcoding/poker_oddsbaccarat/define/BET_AREA"
)

type FeedbackNode struct {
	Current_chip float64       //当前筹码
	Bet_amount   int           //下注额
	Result_area  BET_AREA.TYPE //结果区域
	Result_score float64       //结果得分
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
