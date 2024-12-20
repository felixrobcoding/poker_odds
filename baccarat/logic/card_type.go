/*
功能：牌型
说明：
*/
package logic

import (
	"github.com/felixrobcoding/poker_oddsbaccarat/define"
	"github.com/felixrobcoding/poker_oddsbaccarat/define/CARD_TYPE"
)

// Card_type 牌型
func Card_type(cards []byte) CARD_TYPE.TYPE {
	len := len(cards)
	if len < 2 {
		return CARD_TYPE.ERROR
	}

	pt := Points(cards)
	if len == 2 { //天牌
		if pt == define.POINT_8 || pt == define.POINT_9 {
			return CARD_TYPE.NATURAL
		}
	}

	//对子
	value_0 := common.Value(cards[0])
	value_1 := common.Value(cards[1])
	if value_0 == value_1 {
		return CARD_TYPE.PAIR
	}
	return CARD_TYPE.ERROR
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
