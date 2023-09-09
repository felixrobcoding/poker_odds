/*
功能：查找
说明：
*/
package algorithm

// Find_cnt 查找个数
func Find_cnt[T byte | int](all []T, element T) int {
	cnt := 0
	for _, v := range all {
		if v == element {
			cnt++
		}
	}
	return cnt
}

// Find_value_cnt 查找牌值个数
func Find_value_cnt(cards []byte, value byte, Value ValueFunc) int {
	cnt := 0
	for _, v := range cards {
		if Value(v) == value {
			cnt++
		}
	}
	return cnt
}

// Find_suit_cnt 查找花色个数
func Find_suit_cnt(cards []byte, suit byte, Suit SuitFunc) int {
	cnt := 0
	for _, v := range cards {
		if Suit(v) == suit {
			cnt++
		}
	}
	return cnt
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
