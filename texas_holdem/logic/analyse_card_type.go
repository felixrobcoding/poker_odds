/*
功能：分析牌型
说明：
*/
package logic

import (
	"Odds/common/ORDER_TYPE"
	"Odds/common/algorithm"
	"Odds/texas_holdem/define"
)

type HandlerAnalyseCardType func(item *AnalyseItem) (bool, []byte)

// sort_high_card 高牌排序
func sort_high_card(item *AnalyseItem) []byte {
	if item == nil {
		return []byte{}
	}
	//转逻辑值
	lcards := cards_2_lcards(item.cards)
	//值降序排序
	sort_cards := algorithm.Sort_by_value(lcards, Suit, Value, ORDER_TYPE.DESC)
	return sort_cards
}

// is_one_pair 是否一对
func is_one_pair(item *AnalyseItem) (bool, []byte) {
	if item == nil {
		return false, []byte{}
	}

	if item.same_value.pair_cnt == 1 && item.same_value.three_cnt == 0 {
		//值个数降序排序
		sort_cards := algorithm.Sort_by_value_cnt(item.cards, Suit, Value, ORDER_TYPE.DESC)
		return true, sort_cards
	}
	return false, []byte{}
}

// is_two_pair 是否两对
func is_two_pair(item *AnalyseItem) (bool, []byte) {
	if item == nil {
		return false, []byte{}
	}

	if item.same_value.pair_cnt == 2 {
		//值个数降序排序
		sort_cards := algorithm.Sort_by_value_cnt(item.cards, Suit, Value, ORDER_TYPE.DESC)
		return true, sort_cards

	}
	return false, []byte{}
}

// is_three_of_a_kind 是否三条
func is_three_of_a_kind(item *AnalyseItem) (bool, []byte) {
	if item == nil {
		return false, []byte{}
	}

	if item.same_value.three_cnt == 1 && item.same_value.pair_cnt == 0 {
		//值个数降序排序
		sort_cards := algorithm.Sort_by_value_cnt(item.cards, Suit, Value, ORDER_TYPE.DESC)
		return true, sort_cards
	}
	return false, []byte{}
}

// is_straight 是否顺子
func is_straight(item *AnalyseItem) (bool, []byte) {
	if item == nil {
		return false, []byte{}
	}

	if item.Card_cnt() != 5 {
		return false, []byte{}
	}

	//10,J,Q,K,A
	straight_cnt := 0
	if item.value_cnts[0] > 0 { //包含A
		straight_cnt++
		for i := len(item.value_cnts) - 2; i > 0; i-- {
			if item.value_cnts[i] > 0 {
				straight_cnt++
			} else {
				break
			}
		}
		if straight_cnt == 5 {
			//转逻辑值
			lcards := cards_2_lcards(item.cards)
			//值降序排序
			sort_cards := algorithm.Sort_by_value(lcards, Suit, Value, ORDER_TYPE.DESC)
			return true, sort_cards
		}
	}

	//其他
	start_index := -1
	straight_cnt = 0
	for k, v := range item.value_cnts {
		if start_index == -1 && v > 0 {
			start_index = k
			straight_cnt = 1
			continue
		}

		//结束循环
		if start_index >= 0 && v == 0 {
			break
		}

		if start_index >= 0 && v > 0 {
			straight_cnt++
		}
	}

	//判断是否是顺子
	if straight_cnt == 5 {
		//值降序排序
		sort_cards := algorithm.Sort_by_value(item.cards, Suit, Value, ORDER_TYPE.DESC)
		return true, sort_cards
	}
	return false, []byte{}
}

// is_flush 是否同花
func is_flush(item *AnalyseItem) (bool, []byte) {
	if item == nil {
		return false, []byte{}
	}

	if item.Card_cnt() < 3 {
		return false, []byte{}
	}

	suit_cnt := 0
	for _, v := range item.suit_cnts {
		if v > 0 {
			suit_cnt++
		}
	}
	if suit_cnt == 1 {
		//转逻辑值
		lcards := cards_2_lcards(item.cards)
		//值降序排序
		sort_cards := algorithm.Sort_by_value(lcards, Suit, Value, ORDER_TYPE.DESC)
		return true, sort_cards
	}
	return false, []byte{}
}

// is_full_house 是否葫芦
func is_full_house(item *AnalyseItem) (bool, []byte) {
	if item == nil {
		return false, []byte{}
	}

	if item.same_value.three_cnt == 1 && item.same_value.pair_cnt == 1 {
		//值个数降序排序
		sort_cards := algorithm.Sort_by_value_cnt(item.cards, Suit, Value, ORDER_TYPE.DESC)
		return true, sort_cards
	}
	return false, []byte{}
}

// is_four_of_a_kind 是否四条
func is_four_of_a_kind(item *AnalyseItem) (bool, []byte) {
	if item == nil {
		return false, []byte{}
	}

	if item.same_value.four_cnt == 1 && item.same_value.single_cnt == 1 {
		//值个数降序排序
		sort_cards := algorithm.Sort_by_value_cnt(item.cards, Suit, Value, ORDER_TYPE.DESC)
		return true, sort_cards
	}
	return false, []byte{}
}

// is_straight_flush 是否同花顺
func is_straight_flush(item *AnalyseItem) (bool, []byte) {
	if is, _ := is_flush(item); !is {
		return false, []byte{}
	}
	if is, sort_cards := is_straight(item); is {
		return true, sort_cards
	}

	return false, []byte{}
}

// is_royal_straight_flush 是否皇家同花顺
func is_royal_straight_flush(item *AnalyseItem) (bool, []byte) {
	if is, sort_cards := is_straight_flush(item); is {
		if algorithm.Find_value_cnt(sort_cards, define.LOGIC_VALUE_A, Value) > 0 { //A需要转化为逻辑牌
			return true, sort_cards
		}
	}
	return false, []byte{}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
