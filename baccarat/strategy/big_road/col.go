/*
功能：大路
说明：
*/
package big_road

import (
	"Odds/baccarat/define/BET_AREA"
)

// 列,列中的所有元素都相同
type Col struct {
	bet_areas []BET_AREA.TYPE
}

// 插入列中
func (c *Col) push(bet_area BET_AREA.TYPE) {
	c.bet_areas = append(c.bet_areas, bet_area)
}

func (c *Col) Bet_area() BET_AREA.TYPE {
	return c.bet_areas[0]
}

// 列中的个数
func (c *Col) Cnt() int {
	return len(c.bet_areas)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
