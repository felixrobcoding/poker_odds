/*
功能：大路-节点
说明：
*/
package big_road

import (
	"github.com/felixrobcoding/poker_oddsbaccarat/define/BET_AREA"
)

// 大路中的节点
type Node struct {
	bet_area     BET_AREA.TYPE //下注区域
	bet_amount   int           //下注额
	result_area  BET_AREA.TYPE //结果区域
	result_score float64       //结果
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
