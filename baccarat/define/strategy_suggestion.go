/*
功能：策略建议
说明：
*/
package define

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/define/STYLE"
	"fmt"
)

const (
	DEFAULT_BET_TIMES = 1.0
)

// 策略建议
type StrategySuggestion struct {
	Style     STYLE.TYPE    //形态
	Bet_area  BET_AREA.TYPE //下注区域
	Bet_times float64       //下注倍数
}

// 默认的策略建议
func NewStrategySuggestion() *StrategySuggestion {
	return &StrategySuggestion{
		Style:     STYLE.FOLLOW,
		Bet_area:  BET_AREA.BANKER,
		Bet_times: DEFAULT_BET_TIMES,
	}
}

// 转字符串描述
func (s *StrategySuggestion) String() string {
	return fmt.Sprintf("[style:%s,bet_area:%s,bet_times:%.1f]", s.Style.String(), s.Bet_area.String(), s.Bet_times)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
