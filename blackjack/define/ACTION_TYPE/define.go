/*
功能：动作类型
说明：
*/
package ACTION_TYPE

type TYPE int

const (
	ERROR       TYPE = 0
	HIT         TYPE = 0x0001 //要牌
	DOUBLE_DOWN TYPE = 0x0002 //加倍叫停
	STAND       TYPE = 0x0004 //停牌
	SPLIT       TYPE = 0x0008 //分牌
	INSURANCE   TYPE = 0x0010 //保险
	SURRENDER   TYPE = 0x0020 //投降
	MIN         TYPE = HIT
	MAX         TYPE = SURRENDER
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
