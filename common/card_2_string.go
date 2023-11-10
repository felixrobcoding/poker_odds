/*
功能：转字符串
说明：
*/
package common

import (
	"fmt"
	"reflect"
)

const (
	ERROR_SIGN = "?"
)

// 牌数组转字符串
func Cards_2_string(cards []byte) string {
	return slice_2_string(cards)
}

// 整型数组转字符串
func Ints_2_string(ints []int) string {
	return slice_2_string(ints)
}

func Card_2_sign(card byte) string {
	suits := []string{"♦", "♣", "♥", "♠"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

	suit := Suit(card)
	value := Value(card)

	if int(suit) < 0 || int(suit) >= len(suits) {
		return ERROR_SIGN
	}

	if value == VALUE_A || value == LOGIC_VALUE_A {
		value = LOGIC_VALUE_A
	}
	if int(value) < 2 || int(value) > LOGIC_VALUE_A {
		return ERROR_SIGN
	}

	return values[value-2] + suits[suit]
}

func Cards_2_sign(cards []byte) string {
	len := len(cards)
	str := ""
	for i := 0; i < len; i++ {
		str += Card_2_sign(cards[i])
		if i < len-1 {
			str += " "
		}
	}
	return str
}

// 模板函数-数组转字符串
func slice_2_string[T byte | int](all_elements []T) string {
	if all_elements == nil {
		return ""
	}

	len := len(all_elements)
	str := fmt.Sprintf("个数:%d,", len)
	for i := 0; i < len; i++ {
		reflect_type_name := reflect.ValueOf(all_elements[i]).Type().Name()

		if reflect_type_name == "uint8" {
			str += fmt.Sprintf("0x%02X", all_elements[i])
		} else if reflect_type_name == "int" {
			str += fmt.Sprintf("%02d", all_elements[i])
		}

		if i < len-1 {
			str += ","
		}
	}
	return str
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
