/*
功能：排序
说明：
*/
package algorithm

import (
	"github.com/felixrobcoding/poker_oddscommon/ORDER_TYPE"
	"sort"
)

// Sort_by_suit 按花色排序,默认升序
func Sort_by_suit(cards []byte, Suit SuitFunc, Value ValueFunc, order_types ...ORDER_TYPE.TYPE) []byte {
	if Suit == nil || Value == nil {
		return []byte{}
	}

	tmp_cards := make([]byte, len(cards))
	copy(tmp_cards, cards)

	order_type := ORDER_TYPE.ASC
	if len(order_types) == 1 {
		order_type = order_types[0]
	}

	//升序,从小到大
	if order_type == ORDER_TYPE.ASC {
		sort.SliceStable(tmp_cards, func(i, j int) bool {
			suit_i := Suit(tmp_cards[i])
			suit_j := Suit(tmp_cards[j])
			if suit_i == suit_j { //值大的在前
				return Value(tmp_cards[j]) < Value(tmp_cards[i])
			}
			return suit_i < suit_j
		})
	} else { //降序,从大到小
		sort.SliceStable(tmp_cards, func(i, j int) bool {
			suit_i := Suit(tmp_cards[i])
			suit_j := Suit(tmp_cards[j])
			if suit_i == suit_j { //值大的在前
				return Value(tmp_cards[j]) < Value(tmp_cards[i])
			}
			return suit_j < suit_i
		})
	}
	return tmp_cards
}

// Sort_by_suit_cnt 按花色个数排序,默认升序
func Sort_by_suit_cnt(cards []byte, Suit SuitFunc, Value ValueFunc, order_types ...ORDER_TYPE.TYPE) []byte {
	if Suit == nil || Value == nil {
		return []byte{}
	}

	tmp_cards := make([]byte, len(cards))
	copy(tmp_cards, cards)

	order_type := ORDER_TYPE.ASC
	if len(order_types) == 1 {
		order_type = order_types[0]
	}

	//升序,从小到大
	if order_type == ORDER_TYPE.ASC {
		sort.SliceStable(tmp_cards, func(i, j int) bool {
			suit_i := Suit(tmp_cards[i])
			suit_j := Suit(tmp_cards[j])
			if suit_i == suit_j { //值小的在前
				return Value(tmp_cards[i]) < Value(tmp_cards[j])
			}
			suit_j_cnt := Find_suit_cnt(tmp_cards, suit_j, Suit)
			suit_i_cnt := Find_suit_cnt(tmp_cards, suit_i, Suit)
			if suit_i_cnt == suit_j_cnt {
				return suit_i < suit_j
			}
			return suit_i_cnt < suit_j_cnt
		})
	} else { //降序,从大到小
		sort.SliceStable(tmp_cards, func(i, j int) bool {
			suit_i := Suit(tmp_cards[i])
			suit_j := Suit(tmp_cards[j])
			if suit_i == suit_j { //值大的在前
				return Value(tmp_cards[j]) < Value(tmp_cards[i])
			}
			suit_j_cnt := Find_suit_cnt(tmp_cards, suit_j, Suit)
			suit_i_cnt := Find_suit_cnt(tmp_cards, suit_i, Suit)
			if suit_i_cnt == suit_j_cnt {
				return suit_j < suit_i
			}
			return suit_j_cnt < suit_i_cnt
		})
	}
	return tmp_cards
}

// Sort_by_value 按值大小排序,默认升序
func Sort_by_value(cards []byte, Suit SuitFunc, Value ValueFunc, order_types ...ORDER_TYPE.TYPE) []byte {
	if Suit == nil || Value == nil {
		return []byte{}
	}

	tmp_cards := make([]byte, len(cards))
	copy(tmp_cards, cards)

	order_type := ORDER_TYPE.ASC
	if len(order_types) == 1 {
		order_type = order_types[0]
	}

	//升序,从小到大
	if order_type == ORDER_TYPE.ASC {
		sort.SliceStable(tmp_cards, func(i, j int) bool {
			value_i := Value(tmp_cards[i])
			value_j := Value(tmp_cards[j])
			if value_i == value_j { //花色小的在前
				return Suit(tmp_cards[i]) < Suit(tmp_cards[j])
			}
			return Value(tmp_cards[i]) < Value(tmp_cards[j])
		})
	} else { //降序,从大到小
		sort.SliceStable(tmp_cards, func(i, j int) bool {
			value_i := Value(tmp_cards[i])
			value_j := Value(tmp_cards[j])
			if value_i == value_j { //花色大的在前
				return Suit(tmp_cards[j]) < Suit(tmp_cards[i])
			}
			return Value(tmp_cards[j]) < Value(tmp_cards[i])
		})
	}
	return tmp_cards
}

// Sort_by_value_cnt 按值个数排序,默认升序
func Sort_by_value_cnt(cards []byte, Suit SuitFunc, Value ValueFunc, order_types ...ORDER_TYPE.TYPE) []byte {
	if Suit == nil || Value == nil {
		return []byte{}
	}

	tmp_cards := make([]byte, len(cards))
	copy(tmp_cards, cards)

	order_type := ORDER_TYPE.ASC
	if len(order_types) == 1 {
		order_type = order_types[0]
	}

	//升序,从小到大
	if order_type == ORDER_TYPE.ASC {
		sort.SliceStable(tmp_cards, func(i, j int) bool {
			value_i := Value(tmp_cards[i])
			value_j := Value(tmp_cards[j])
			if value_i == value_j { //花色小的在前
				return Suit(tmp_cards[i]) < Suit(tmp_cards[j])
			}
			value_j_cnt := Find_value_cnt(tmp_cards, value_j, Value)
			value_i_cnt := Find_value_cnt(tmp_cards, value_i, Value)
			if value_i_cnt == value_j_cnt {
				return value_i < value_j
			}
			return value_i_cnt < value_j_cnt
		})
	} else { //降序,从大到小
		sort.SliceStable(tmp_cards, func(i, j int) bool {
			value_i := Value(tmp_cards[i])
			value_j := Value(tmp_cards[j])
			if value_i == value_j { //花色大的在前
				return Suit(tmp_cards[j]) < Suit(tmp_cards[i])
			}
			value_j_cnt := Find_value_cnt(tmp_cards, value_j, Value)
			value_i_cnt := Find_value_cnt(tmp_cards, value_i, Value)
			if value_i_cnt == value_j_cnt {
				return value_j < value_i
			}
			return value_j_cnt < value_i_cnt
		})
	}
	return tmp_cards
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
