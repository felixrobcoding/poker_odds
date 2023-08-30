/*
功能：牌型
说明：
*/
package CARD_TYPE

type TYPE int

const (
	ERROR   TYPE = 0
	PAIR    TYPE = 0x0001 //对子
	NATURAL TYPE = 0x0002 //天牌
	MIN     TYPE = PAIR
	MAX     TYPE = NATURAL
)

// 是否有效
func Is_valid(t TYPE) bool {
	if t >= MIN && t <= MAX {
		return true
	}
	return false
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
