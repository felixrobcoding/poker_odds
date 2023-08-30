/*
功能：牌型
说明：
*/
package CARD_TYPE

type TYPE int

const (
	EX_FLUSH_3           TYPE = 0x000001 //三张同花
	EX_FLUSH_4           TYPE = 0x000002 //四张同花
	EX_FLUSH_5           TYPE = 0x000004 //四张同花
	EX_MORE_ONE_PAIR     TYPE = 0x000008 //至少一对
	HIGH_CARD            TYPE = 0x001000 //高牌
	ONE_PAIR             TYPE = 0x002000 //一对
	TWO_PAIR             TYPE = 0x004000 //两对
	THREE_OF_A_KIND      TYPE = 0x008000 //三条
	STRAIGHT             TYPE = 0x010000 //顺子
	FLUSH                TYPE = 0x020000 //同花
	FULL_HOUSE           TYPE = 0x040000 //葫芦
	FOUR_OF_A_KIND       TYPE = 0x080000 //四条
	STRAIGHT_FLUSH       TYPE = 0x100000 //同花顺
	ROYAL_STRAIGHT_FLUSH TYPE = 0x200000 //皇家同花顺
	MIN                  TYPE = EX_FLUSH_3
	MAX                  TYPE = ROYAL_STRAIGHT_FLUSH
)

//-----------------------------------------------
//					the end
//-----------------------------------------------
