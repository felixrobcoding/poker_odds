/*
功能：便利函数
说明：
*/
package common

import (
	"fmt"
	"reflect"
)

// 模板函数
func String[T byte | int](cards []T) string {
	if cards == nil {
		return ""
	}

	str := fmt.Sprintf("个数:%d,", len(cards))
	for i := 0; i < len(cards); i++ {

		reflect_type_name := reflect.ValueOf(cards[i]).Type().Name()

		if reflect_type_name == "uint8" {
			str += fmt.Sprintf(" 0x%02X", cards[i])
		} else if reflect_type_name == "int" {
			str += fmt.Sprintf(" %02d", cards[i])
		}
	}
	return str
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
