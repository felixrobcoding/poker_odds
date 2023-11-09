/*
功能：测试
说明：
*/
package logic

import (
	"fmt"
	"testing"
)

func Test_sign_1(t *testing.T) {
	{
		cards := []byte{0x01, 0x12, 0x23, 0x34, 0x15, 0x4E}
		fmt.Printf("%s\r\n", Signs(cards))
	}
	{
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
