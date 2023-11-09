/*
功能：牌符号
说明：
*/
package logic

import (
	"Odds/common"
	"Odds/texas_holdem/define"
)

const (
	ERROR_SIGN = "?"
)

func Sign(card byte) string {
	suits := []string{"♦", "♣", "♥", "♠"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

	suit := Suit(card)
	value := Value(card)

	if int(suit) < 0 || int(suit) >= len(suits) {
		return ERROR_SIGN
	}

	if value == common.VALUE_A || value == define.LOGIC_VALUE_A {
		value = define.LOGIC_VALUE_A
	}
	if int(value) < 2 || int(value) > define.LOGIC_VALUE_A {
		return ERROR_SIGN
	}

	return values[value-2] + suits[suit]
}

func Signs(cards []byte) string {
	str := ""
	for i := 0; i < len(cards); i++ {
		str += Sign(cards[i])
		if i < len(cards)-1 {
			str += " "
		}
	}
	return str
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
