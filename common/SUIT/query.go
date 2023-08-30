/*
功能：花色描述
说明：
*/
package SUIT

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
