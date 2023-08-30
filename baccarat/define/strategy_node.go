/*
功能：策略节点
说明：
*/
package define

import (
	"Odds/baccarat/define/BET_AREA"
)

// 策略节点
type StrategyNode struct {
	Current_chip         float64       //当前筹码
	Current_bet_area     BET_AREA.TYPE //当前下注区域
	Current_bet          int           //当前下注
	Current_win_bet_area BET_AREA.TYPE //当前获胜区域
	Current_score        float64       //当前得分
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
