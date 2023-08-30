/*
功能：测试
说明：
*/
package logic

import (
	"fmt"
	"testing"
)

func Test_card_type(t *testing.T) {
	{
		// cards := []byte{0x11, 0x12, 0x13, 0x14, 0x15}
		// fmt.Printf("%+v\r\n", cards)

		// item := Analyse(cards)
		// if item != nil {
		// 	fmt.Printf("%s\r\n", item)
		// }
	}
	{
		cards := []byte{0x11, 0x1A, 0x1B, 0x1C, 0x1D}
		fmt.Printf("%+v\r\n", cards)

		item := Analyse(cards)
		if item != nil {
			fmt.Printf("%s\r\n", item)
		}
	}
}

func TestXxx2(t *testing.T) {

	cards := []byte{0x13, 0x1D, 0x26, 0x3A, 0x41}
	fmt.Printf("%+v\r\n", cards)

	item := NewAnalyseItem(cards)
	if item != nil {
		fmt.Printf("%+v\r\n", item)
	}

	cards_1 := []byte{0x11, 0x21, 0x31, 0x41}
	fmt.Printf("%+v\r\n", cards_1)

	item_1 := NewAnalyseItem(cards_1)
	if item_1 != nil {
		fmt.Printf("%+v\r\n", item_1)
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
