/*
功能：牌型
说明：
Mixed pair【混合对子】 = same number/face card value, different suit and colour.
Coloured pair【同色对子】 = same number/face card value, same colour, different suit.
Perfect pair【完美对子】 = same number/face card value, same colour, same suit.
*/
package CARD_TYPE

type TYPE int

const (
	ERROR      TYPE = 0
	POINT      TYPE = 1 //点牌[计算点数的牌]
	SURRENDER  TYPE = 2 //投降
	BUST       TYPE = 3 //爆点
	BLACK_JACK TYPE = 4
	MIN        TYPE = POINT
	MAX        TYPE = BLACK_JACK
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
