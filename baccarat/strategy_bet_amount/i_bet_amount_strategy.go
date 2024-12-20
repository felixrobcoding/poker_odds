/*
功能：下注额策略-接口
说明：
*/
package strategy_bet_amount

import "github.com/felixrobcoding/poker_oddscommon/BET_AMOUNT_STRATEGY"

type IBetAmountStrategy interface {
	Feedback_node_append(node *FeedbackNode)            //追加反馈节点
	Feedback_node_clear()                               //反馈节点清除
	Query_bet_amount() (int, error)                     //查询下注额
	Is_enough_money(amount int) bool                    //查询余额是否足够
	Query_option() (int, int, BET_AMOUNT_STRATEGY.TYPE) //查询配置
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
