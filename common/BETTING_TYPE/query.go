/*
功能：map查询
说明：
*/
package BETTING_TYPE

//文字描述
func (t TYPE) String() string {
	value, ok := type_map[t]
	if ok {
		return value.txt_eng
	}
	return ""
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
