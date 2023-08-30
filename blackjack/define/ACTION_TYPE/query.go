/*
功能：map查询
说明：
*/
package ACTION_TYPE

//文字描述
func (t TYPE) String() string {
	value, ok := type_map[t]
	if ok {
		return value.txt_chs
	}
	return ""
}

func (t TYPE) String_eng() string {
	value, ok := type_map[t]
	if ok {
		return value.txt_eng
	}
	return ""
}

func (t TYPE) String_short() string {
	value, ok := type_map[t]
	if ok {
		return value.txt_short
	}
	return ""
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
