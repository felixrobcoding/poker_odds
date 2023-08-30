/*
功能：花色
说明：
*/
package SUIT

type TYPE int

const (
	ERROR   TYPE = 0
	DIAMOND TYPE = 1 //方块
	CLUB    TYPE = 2 //梅花
	HEART   TYPE = 3 //红心
	SPADE   TYPE = 4 //黑桃
	MIN     TYPE = DIAMOND
	MAX     TYPE = SPADE
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
