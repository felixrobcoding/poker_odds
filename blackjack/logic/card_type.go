/*
功能：判断牌型
说明：
*/
package logic

import (
	"Odds/blackjack/define"
	"Odds/blackjack/define/CARD_TYPE"
	"Odds/blackjack/define/HAND_TYPE"
	"Odds/common"
	"Odds/common/algorithm"
)

// 分析hand type
func Analyse_hand_type(player_cards []byte) HAND_TYPE.TYPE {

	//分牌
	if len(player_cards) == 2 {
		if cnt := algorithm.Find_value_cnt(player_cards, common.Value(player_cards[0]), common.Value); cnt == 2 {
			return HAND_TYPE.SPLITS
		}
	}

	A_cnt := algorithm.Find_value_cnt(player_cards, common.VALUE_A, common.Value)
	if A_cnt > 0 { //可能是soft
		points, _ := Points(player_cards)
		if points[1] < define.POINT_BUST { //A存在可变点数
			return HAND_TYPE.SOFT
		}
	}

	return HAND_TYPE.HARD
}

// 判断牌型
func Analyse_card_type(cards []byte) CARD_TYPE.TYPE {
	if is_blackjack(cards) {
		return CARD_TYPE.BLACK_JACK
	}
	if is_bust(cards) {
		return CARD_TYPE.BUST
	}
	return CARD_TYPE.POINT
}

func is_blackjack(cards []byte) bool {
	if len(cards) != 2 {
		return false
	}
	A_cnt := algorithm.Find_value_cnt(cards, common.VALUE_A, common.Value)
	if A_cnt != 1 {
		return false
	}

	value_tens := []byte{
		common.VALUE_K,
		common.VALUE_Q,
		common.VALUE_J,
		common.VALUE_T,
	}
	for _, v := range value_tens {
		if cnt := algorithm.Find_value_cnt(cards, v, common.Value); cnt > 0 {
			return true
		}
	}
	return false
}

func is_bust(cards []byte) bool {
	points, _ := Points(cards)
	if points[0] >= define.POINT_BUST { //爆点
		return true
	}
	return false
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
