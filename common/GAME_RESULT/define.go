/*
功能：游戏结果
说明：
*/
package GAME_RESULT

type TYPE int

const (
	ERROR TYPE = 0
	LOSE  TYPE = 1 //输
	PUSH  TYPE = 2 //和
	WIN   TYPE = 3 //赢
	MIN   TYPE = LOSE
	MAX   TYPE = WIN
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
