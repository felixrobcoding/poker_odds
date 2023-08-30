/*
功能：删除
说明：
*/
package algorithm

import "sort"

//删除
func Remove[T byte | int](all []T, target_element T) (removed_cnt int, results []T) {
	results = make([]T, 0)
	if cnt := Find_cnt(all, target_element); cnt <= 0 {
		removed_cnt = 0
		return
	}

	for _, v := range all {
		if v == target_element {
			removed_cnt++
		} else {
			results = append(results, v)
		}
	}
	return removed_cnt, results
}

//删除
func Removes[T byte | int](all []T, target_elements []T) (removed_cnt int, results []T) {
	results = make([]T, 0)
	copy(results, all)

	for _, v := range target_elements {
		tmp_removed_cnt, tmp_results := Remove(results, v)
		if removed_cnt <= 0 {
			return 0, []T{}
		}
		removed_cnt += tmp_removed_cnt
		results = tmp_results
	}
	return
}

//删除牌值
func Remove_value(all []byte, target_value byte, Value ValueFunc) (removed_cnt int, result_cards []byte) {
	result_cards = make([]byte, 0)
	if cnt := Find_value_cnt(all, target_value, Value); cnt <= 0 {
		removed_cnt = 0
		return
	}

	for _, v := range all {
		if Value(v) == target_value {
			removed_cnt++
		} else {
			result_cards = append(result_cards, v)
		}
	}
	return removed_cnt, result_cards
}

func Remove_values(all []byte, target_values []byte, Value ValueFunc) (removed_cnt int, result_cards []byte) {
	result_cards = make([]byte, 0)
	copy(result_cards, all)

	for _, v := range target_values {
		tmp_removed_cnt, tmp_result_cards := Remove_value(result_cards, v, Value)
		if removed_cnt <= 0 {
			return 0, []byte{}
		}
		removed_cnt += tmp_removed_cnt
		result_cards = tmp_result_cards
	}
	return
}

//去重
func Remove_duplication[T byte | int](all []T) []T {
	len := len(all)
	if len == 0 {
		return all
	}

	tmp_all := make([]T, 0)
	copy(tmp_all, all)

	tmp_map := make(map[T]interface{}, 0)
	for _, v := range tmp_all {
		tmp_map[v] = 0
	}

	index := 0
	for k, _ := range tmp_map {
		tmp_all[index] = k
		index++
	}

	//排序,从小到大
	sort.SliceStable(tmp_all, func(i, j int) bool {
		return tmp_all[i] < tmp_all[j]
	})

	return tmp_all
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
