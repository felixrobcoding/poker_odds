/*
功能：输出-x,y轴标题
说明：
*/
package outputer

import (
	"github.com/felixrobcoding/poker_oddsblackjack/define"
	"github.com/felixrobcoding/poker_oddsblackjack/logic"
)

const (
	X_AXIS_CNT          = 10
	Y_AXIS_CNT_HARD     = 10
	Y_AXIS_CNT_HARD_ALL = 17
	Y_AXIS_CNT_SOFT     = 7
	Y_AXIS_CNT_SOFT_ALL = 9
	Y_AXIS_CNT_SPLITS   = 10
)
const (
	HARD_HAND_HEADING   = "Hard"   //标题
	SOFT_HAND_HEADING   = "Soft"   //标题
	SPLITS_HAND_HEADING = "Splits" //标题
)

// 庄家牌转x轴标题,dealer's card
func Dealer_card_2_x_axis_heading(dealer_card byte) string {
	point := logic.Point(dealer_card)
	if point == define.POINT_10 {
		return X_axis_headings()[8]
	}
	if point == define.POINT_A_1 {
		return X_axis_headings()[9]
	}
	return X_axis_headings()[point-2]
}

// x轴标题
func X_axis_headings() []string {
	var x_axis_headings = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "A"} //dealer's card
	if len(x_axis_headings) != X_AXIS_CNT {
		panic("")
	}
	return x_axis_headings
}

// y轴标题-hard
func Y_axis_headings_hard() []string {
	var y_axis_headings_hard = []string{"5-8", "9", "10", "11", "12", "13", "14", "15", "16", "17-21"}
	if len(y_axis_headings_hard) != Y_AXIS_CNT_HARD {
		panic("")
	}
	return y_axis_headings_hard
}

func Y_axis_headings_hard_all() []string {
	var y_axis_headings_hard = []string{"5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21"}
	if len(y_axis_headings_hard) != Y_AXIS_CNT_HARD_ALL {
		panic("")
	}
	return y_axis_headings_hard
}

// y轴标题-soft
func Y_axis_headings_soft() []string {
	var y_axis_headings_soft = []string{"13", "14", "15", "16", "17", "18", "19-21"}
	if len(y_axis_headings_soft) != Y_AXIS_CNT_SOFT {
		panic("")
	}
	return y_axis_headings_soft
}

func Y_axis_headings_soft_all() []string {
	var y_axis_headings_soft = []string{"13", "14", "15", "16", "17", "18", "19", "20", "21"}
	if len(y_axis_headings_soft) != Y_AXIS_CNT_SOFT_ALL {
		panic("")
	}
	return y_axis_headings_soft
}

// y轴标题-分牌
func Y_axis_headings_splits() []string {
	var y_axis_headings_splits = []string{"2,2", "3,3", "4,4", "5,5", "6,6", "7,7", "8,8", "9,9", "10,10", "A,A"}
	if len(y_axis_headings_splits) != Y_AXIS_CNT_SPLITS {
		panic("")
	}
	return y_axis_headings_splits
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
