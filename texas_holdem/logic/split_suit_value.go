/*
功能：分离花色和值
说明：
*/
package logic

//牌花色
func Suit(card byte) byte {
	return (card & 0xF0) >> 4
}

//牌值
func Value(card byte) byte {
	return card & 0x0F
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
