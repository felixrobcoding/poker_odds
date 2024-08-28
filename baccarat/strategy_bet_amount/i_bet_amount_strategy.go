/*
功能：下注额策略-接口
说明：
*/
package strategy_bet_amount

import "Odds/common/BET_AMOUNT_STRATEGY"

type IBetAmountStrategy interface {
	Result_node_append(trend *ResultNode)               //追加结果
	Result_node_clear()                                 //清理结果
	Query_bet_amount() (int, error)                     //查询下注额
	Is_enough_money(payout_amount int) bool             //查询余额是否足够
	Query_option() (int, int, BET_AMOUNT_STRATEGY.TYPE) //查询配置
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
