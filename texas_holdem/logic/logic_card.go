/*
功能：逻辑牌
说明：
*/
package logic

import (
	"Odds/common"
	"Odds/texas_holdem/define"
)

// 牌转逻辑牌
//
//lcard:logic card
func card_2_lcard(card byte) byte {
	suit := Suit(card)
	value := Value(card)

	if value == common.VALUE_A {
		return suit<<4 | define.LOGIC_VALUE_A
	}
	return card
}

//lcards:logic cards
func cards_2_lcards(cards []byte) []byte {
	if len(cards) <= 0 {
		return nil
	}
	lcards := make([]byte, 0)
	for _, v := range cards {
		lcard := card_2_lcard(v)
		lcards = append(lcards, lcard)
	}
	return lcards
}

// 逻辑牌转牌
//
//lcard:logic card
func lcard_2_card(card byte) byte {
	suit := Suit(card)
	value := Value(card)

	if value == define.LOGIC_VALUE_A {
		return suit<<4 | common.VALUE_A
	}
	return card
}

//lcards:logic cards
func lcards_2_cards(lcards []byte) []byte {
	if len(lcards) <= 0 {
		return nil
	}
	cards := make([]byte, 0)
	for _, v := range lcards {
		card := lcard_2_card(v)
		cards = append(cards, card)
	}
	return cards
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
