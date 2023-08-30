/*
功能：玩法策略-查询
说明：
*/
package strategy

import (
	"Odds/blackjack/define"
	"Odds/blackjack/define/ACTION_TYPE"
	"Odds/blackjack/define/HAND_TYPE"
	"Odds/blackjack/logic"
	"Odds/blackjack/strategy/node"
	"Odds/blackjack/strategy/outputer"
	"Odds/blackjack/strategy/query_times"
	"Odds/blackjack/strategy/standard_strategy"
	"Odds/common"
	"Odds/common/algorithm"
	"fmt"

	"github.com/poker-x-studio/x/xdebug"
)

// 闲家查询action
func Player_query_action(player_cards []byte, dealer_card byte) (action ACTION_TYPE.TYPE, point int, dealer_value string) {
	if player_cards == nil {
		panic("")
	}

	//
	hand_type := logic.Analyse_hand_type(player_cards)

	//点数
	points, _ := logic.Points(player_cards)
	if (hand_type == HAND_TYPE.HARD) || (hand_type == HAND_TYPE.SOFT) {
		point = logic.Player_pick_best_point_to_query(hand_type, points)
	} else if hand_type == HAND_TYPE.SPLITS {
		value_cnt := algorithm.Find_value_cnt(player_cards, common.VALUE_A, logic.Value)
		if value_cnt > 0 {
			point = define.POINT_A_11
		} else {
			point = logic.Point(player_cards[0])
		}
	}

	//
	dealer_value = outputer.Dealer_card_2_x_axis_heading(dealer_card)
	//key
	key := node.Make_key(hand_type, point, dealer_value)

	strategy_map := standard_strategy.Get_strategy_map()
	node, ok := strategy_map[key]
	if ok {
		//查询次数增加
		query_times.Instance().Increase(key)

		action = node.Action

		if xdebug.Is_debug() {
			fmt.Printf("player_cards:%s[点数:%d],dealer_card:0x%02X[%s],hand_type:%s,key:%s,action:%s\r\n", common.String(player_cards), point, dealer_card, dealer_value, hand_type, key, action)
		}
		return
	}
	action = ACTION_TYPE.STAND
	if xdebug.Is_debug() {
		fmt.Printf("player_cards:%s[点数:%d],dealer_card:0x%02X[%s],hand_type:%s,key:%s,action:%s\r\n", common.String(player_cards), point, dealer_card, dealer_value, hand_type, key, action)
	}

	return
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
