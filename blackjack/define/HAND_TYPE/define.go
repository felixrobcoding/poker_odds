/*
功能：hand type
说明：
*/
package HAND_TYPE

type TYPE int

const (
	ERROR  TYPE = 0
	HARD   TYPE = 1 //不包含A,再來一张牌也不会淘汰
	SOFT   TYPE = 2 //包含A
	SPLITS TYPE = 3 //对子,可分牌
	MIN    TYPE = HARD
	MAX    TYPE = SPLITS
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
