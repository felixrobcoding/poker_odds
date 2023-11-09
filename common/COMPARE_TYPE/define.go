/*
功能：比较类型
说明：
*/
package COMPARE_TYPE

type TYPE int

const (
	ERROR   TYPE = 0
	SMALLER TYPE = 1 //更小
	EQUAL   TYPE = 2 //相等
	BIGGER  TYPE = 3 //更大
	MIN     TYPE = SMALLER
	MAX     TYPE = BIGGER
)

//是否有效
func Is_valid(t TYPE) bool {
	if t >= MIN && t <= MAX {
		return true
	}
	return false
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
