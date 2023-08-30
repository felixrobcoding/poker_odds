/*
功能：map查询
说明：
*/
package CARD_TYPE

//赔率
func (t TYPE) Odds() float64 {
	value, ok := type_map[t]
	if ok {
		return value.odds
	}
	return 1.0
}

//文字描述
func (t TYPE) String() string {
	value, ok := type_map[t]
	if ok {
		return value.txt_chs
	}
	return ""
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
