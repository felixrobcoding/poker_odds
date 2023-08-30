/*
功能：map查询
说明：
*/
package HAND_TYPE

//文字描述
func (t TYPE) String() string {
	value, ok := type_map[t]
	if ok {
		return value.txt_eng
		//return value.txt_eng + value.txt_chs
	}
	return ""
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
