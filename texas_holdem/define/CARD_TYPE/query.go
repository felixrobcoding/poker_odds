/*
功能：map查询
说明：
*/
package CARD_TYPE

import "github.com/poker-x-studio/x/xutils"

//文字描述
func (t TYPE) String() string {
	value, ok := type_map[t]
	if ok {
		return value.txt_chs
	}
	return ""
}

func String(types []TYPE) string {
	str := ""
	for _, v := range types {
		str += v.String() + xutils.SPACE
	}
	return str
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
