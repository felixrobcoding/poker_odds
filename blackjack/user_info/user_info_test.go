/*
功能：测试
说明：
*/
package user_info

import (
	"fmt"
	"testing"

	"github.com/poker-x-studio/x/xutils"
)

func TestXxx(t *testing.T) {
	numbers := []int{1, 2, 3}
	fmt.Printf("%+v\r\n", numbers)

	insert_1_numbers := xutils.Slice_insert(numbers, 0, 4)
	fmt.Printf("%+v\r\n", insert_1_numbers)

	insert_2_numbers := xutils.Slice_insert(numbers, 1, 4)
	fmt.Printf("%+v\r\n", insert_2_numbers)

	insert_3_numbers := xutils.Slice_insert(numbers, 2, 4)
	fmt.Printf("%+v\r\n", insert_3_numbers)

	insert_4_numbers := xutils.Slice_insert(numbers, 3, 4)
	fmt.Printf("%+v\r\n", insert_4_numbers)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
