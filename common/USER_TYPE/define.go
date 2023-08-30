/*
功能：用户类型
说明：
*/
package USER_TYPE

type TYPE int

const (
	ERROR  TYPE = 0
	PLAYER TYPE = 1 //闲家
	BANKER TYPE = 2 //庄家
	MIN    TYPE = PLAYER
	MAX    TYPE = BANKER
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
