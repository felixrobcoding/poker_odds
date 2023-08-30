/*
功能：压注区域
说明：
*/
package BET_AREA

type TYPE int

const (
	ERROR       TYPE = 0
	BANKER      TYPE = 0x0001 //庄
	PLAYER      TYPE = 0x0002 //闲
	TIE         TYPE = 0x0004 //和
	BANKER_PAIR TYPE = 0x0008 //庄对
	PLAYER_PAIR TYPE = 0x0010 //闲对
	MIN         TYPE = BANKER
	MAX         TYPE = PLAYER_PAIR
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
