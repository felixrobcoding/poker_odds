/*
功能：下注区域建议
说明：
*/
package suggestion

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/define/STYLE"
	"fmt"
)

const (
	DEFAULT_BET_TIMES = 1.0
)

// 下注区域建议
type BetAreaSuggestion struct {
	Style    STYLE.TYPE    //形态
	Bet_area BET_AREA.TYPE //下注区域
}

// 默认的下注区域建议
func NewBetAreaSuggestion() *BetAreaSuggestion {
	return &BetAreaSuggestion{
		Style:    STYLE.FOLLOW,
		Bet_area: BET_AREA.BANKER,
	}
}

// 转字符串描述
func (s *BetAreaSuggestion) String() string {
	return fmt.Sprintf("[style:%s,bet_area:%s]", s.Style.String(), s.Bet_area.String())
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
