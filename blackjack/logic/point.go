/*
功能：逻辑-计算点数
说明：
*/
package logic

import (
	"Odds/blackjack/define"
	"Odds/blackjack/define/HAND_TYPE"
	"Odds/common"
	"fmt"

	"github.com/poker-x-studio/x/xutils"
)

// 点数
func Point(card byte) int {
	v := common.Value(card)

	value_tens := []byte{
		common.VALUE_K,
		common.VALUE_Q,
		common.VALUE_J,
		common.VALUE_T,
	}
	if xutils.Is_contains(value_tens, v) {
		return define.POINT_10
	}

	if v == common.VALUE_A {
		return define.POINT_A_1
	}

	return int(v)
}

// 点数
// 从小到大排序
func Points(cards []byte) (points []int, points_txt string) {
	a_cnt := 0 //A 个数
	sum_point := 0
	for _, v := range cards {
		point := Point(v)
		if point == define.POINT_A_1 {
			a_cnt++
		}
		sum_point += point
	}

	if a_cnt <= 0 { //不包含A
		points = []int{sum_point}
		points_txt = points_2_string(points)
		return
	}

	//包含A
	points = make([]int, a_cnt+1)
	for i := 0; i <= a_cnt; i++ {
		points[i] = sum_point + i*10
	}
	points_txt = points_2_string(points)
	return
}

// 闲家在查询时候,选择最恰当的点数
func Player_pick_best_point_to_query(hand_type HAND_TYPE.TYPE, points []int) int {
	if points == nil {
		panic("")
	}
	if len(points) == 1 {
		return points[0]
	}

	if hand_type == HAND_TYPE.HARD {
		for _, v := range points {
			if (v >= define.POINT_17) && (v < define.POINT_BUST) {
				return v
			}
		}
	} else if hand_type == HAND_TYPE.SOFT {
		for _, v := range points {
			if (v >= define.POINT_13) && (v <= define.POINT_21) {
				return v
			}
		}
	}
	return points[0]
}

// 闲家在比牌时候,选择最恰当的点数
func Player_pick_best_point_to_compare(points []int) int {
	if points == nil {
		panic("")
	}
	if len(points) == 1 {
		return points[0]
	}

	for _, v := range points {
		if (v >= define.POINT_17) && (v < define.POINT_BUST) {
			return v
		}
	}
	return points[0]
}

// 庄家选择最恰当的点数
func Dealer_pick_best_point(points []int) int {
	if points == nil {
		panic("")
	}
	if len(points) == 1 {
		return points[0]
	}

	for _, v := range points {
		if (v >= define.POINT_17) && (v < define.POINT_BUST) {
			return v
		}
	}
	return points[0]
}

func points_2_string(points []int) string {
	if points == nil {
		return ""
	}

	str := ""
	for i := 0; i < len(points); i++ {
		if i == 0 {
			str += fmt.Sprintf("(")
		}

		if i == len(points)-1 {
			str += fmt.Sprintf("%02d", points[i])
		} else {
			str += fmt.Sprintf("%02d,", points[i])
		}

		if i == len(points)-1 {
			str += fmt.Sprintf(")")
		}
	}
	return str
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
