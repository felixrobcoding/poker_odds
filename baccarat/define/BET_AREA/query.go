/*
功能：map查询
说明：
*/
package BET_AREA

import "sort"

// 文字描述
func (t TYPE) String() string {
	value, ok := type_map[t]
	if ok {
		return value.txt_chs
	}
	return ""
}

// 赔率
func (t TYPE) Odds() float64 {
	value, ok := type_map[t]
	if ok {
		return value.odds
	}
	return 1
}

// 下注区域列表
func All_bet_area() []TYPE {
	types := make([]TYPE, 0)
	for _, v := range type_map {
		types = append(types, v.t)
	}

	//排序
	sort.SliceStable(types, func(i, j int) bool {
		return types[i] < types[j]
	})
	return types
}

func Is_contains(all []TYPE, element TYPE) bool {
	if all == nil {
		return false
	}
	for _, v := range all {
		if v == element {
			return true
		}
	}
	return false
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
