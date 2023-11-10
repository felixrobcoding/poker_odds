/*
功能：测试
说明：
*/
package logic

import (
	"Odds/common"
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	{
		cards := []byte{0x01, 0x31, 0x21, 0x11, 0x1B}
		points, points_txt := Points(cards)
		fmt.Printf("points:%s,points_txt:%s\r\n", common.Ints_2_string(points), points_txt)
	}

	{
		cards := []byte{0x02, 0x35, 0x23, 0x17, 0x1B}
		points, points_txt := Points(cards)
		fmt.Printf("points:%s,points_txt:%s\r\n", common.Ints_2_string(points), points_txt)
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
