/*
功能：map查询
说明：
*/
package GAME_RESULT

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
