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
		cards := []byte{0x11, 0x12, 0x13, 0x14, 0x15}
		fmt.Printf("%+v\r\n", cards)

		pt := Points(cards)
		fmt.Printf("%v\r\n", pt)
	}

}

//-----------------------------------------------
//					the end
//-----------------------------------------------
