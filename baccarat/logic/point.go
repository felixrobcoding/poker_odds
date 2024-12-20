/*
功能：计算点数
说明：
*/
package logic

// 计算点数
func Point(card byte) int {
	value := common.Value(card)
	if value >= common.VALUE_T && value <= common.VALUE_K {
		return 0
	}
	return int(value)
}

// 计算点数
func Points(cards []byte) int {
	sum := 0
	for _, v := range cards {
		sum += Point(v)
	}
	return sum % 10
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
