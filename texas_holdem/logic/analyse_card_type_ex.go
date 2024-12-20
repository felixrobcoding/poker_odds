/*
功能：分析扩展牌型
说明：
*/
package logic

import (
	"github.com/felixrobcoding/poker_oddscommon/ORDER_TYPE"
	"github.com/felixrobcoding/poker_oddscommon/algorithm"
)

// is_ex_flush_3 是否3张同色
func is_ex_flush_3(item *AnalyseItem) (bool, []byte) {
	if item == nil {
		return false, []byte{}
	}

	if item.Card_cnt() < 3 {
		return false, []byte{}
	}

	for _, v := range item.suit_cnts {
		if v == 3 {
			//花色个数降序排序
			sort_cards := algorithm.Sort_by_suit_cnt(item.cards, common.Suit, common.Value, ORDER_TYPE.DESC)
			return true, sort_cards
		}
	}
	return false, []byte{}
}

// is_ex_flush_4 是否4张同色
func is_ex_flush_4(item *AnalyseItem) (bool, []byte) {
	if item == nil {
		return false, []byte{}
	}

	if item.Card_cnt() < 4 {
		return false, []byte{}
	}

	for _, v := range item.suit_cnts {
		if v == 4 {
			//花色个数降序排序
			sort_cards := algorithm.Sort_by_suit_cnt(item.cards, common.Suit, common.Value, ORDER_TYPE.DESC)
			return true, sort_cards
		}
	}
	return false, []byte{}
}

// is_ex_flush_5 是否5张同色
func is_ex_flush_5(item *AnalyseItem) (bool, []byte) {
	if item == nil {
		return false, []byte{}
	}

	if item.Card_cnt() < 5 {
		return false, []byte{}
	}

	for _, v := range item.suit_cnts {
		if v == 5 {
			//花色个数降序排序
			sort_cards := algorithm.Sort_by_suit_cnt(item.cards, common.Suit, common.Value, ORDER_TYPE.DESC)
			return true, sort_cards
		}
	}
	return false, []byte{}
}

// is_ex_more_one_pair 至少一对
func is_ex_more_one_pair(item *AnalyseItem) (bool, []byte) {
	if item == nil {
		return false, []byte{}
	}

	if item.Card_cnt() < 2 {
		return false, []byte{}
	}

	for _, v := range item.value_cnts {
		if v == 2 || v == 3 || v == 4 {
			//值个数降序排序
			sort_cards := algorithm.Sort_by_value_cnt(item.cards, common.Suit, common.Value, ORDER_TYPE.DESC)
			return true, sort_cards
		}
	}
	return false, []byte{}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
