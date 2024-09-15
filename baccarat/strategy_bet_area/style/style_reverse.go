/*
功能：相反形态
说明：列的第一颗之后，执行相反形态
*/
package style

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/define/STYLE"
	"Odds/baccarat/strategy_bet_area/big_road"
	"Odds/baccarat/strategy_bet_area/suggestion"
)

// 相反形态
func check_reverse_style(nodes []*suggestion.FeedbackNode) (bool, *suggestion.BetAreaSuggestion) {
	nodes_cnt := len(nodes)
	if nodes_cnt <= 0 {
		return true, suggestion.NewBetAreaSuggestion()
	}

	big_road_all := big_road.NewBigRoadWithNodes(nodes)
	last_col := big_road_all.Last_col()

	if last_col.Cnt() != 1 {
		return false, nil
	}

	//列的假想第二颗，执行相反下注区
	bet_area := BET_AREA.ERROR
	if last_col.Result_area() == BET_AREA.BANKER {
		bet_area = BET_AREA.PLAYER
	} else {
		bet_area = BET_AREA.BANKER
	}

	return true, &suggestion.BetAreaSuggestion{
		Style:    STYLE.REVERSE,
		Bet_area: bet_area,
		Comment:  "相反下注",
		Alart:    false,
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
