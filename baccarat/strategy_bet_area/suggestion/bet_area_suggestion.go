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
	Comment  string        //描述
	Alart    bool          //警报,检测到特殊形态后报警
}

// 默认的下注区域建议
func NewBetAreaSuggestion() *BetAreaSuggestion {
	return &BetAreaSuggestion{
		Style:    STYLE.FOLLOW,
		Bet_area: BET_AREA.BANKER,
		Comment:  "默认跟随下注",
		Alart:    false,
	}
}

// 转字符串描述
func (s *BetAreaSuggestion) String() string {
	return fmt.Sprintf("[style:%s,bet_area:%s,comment:%s,alart:%t]", s.Style.String(), s.Bet_area.String(), s.Comment, s.Alart)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
