package svg_utils

import (
	"Odds/baccarat/define/BET_AREA"
	"fmt"
	"testing"
)

func Test_win_bet_area_stat(t *testing.T) {
	stats := []*WinBetAreaStat{
		{Total_hands: 10000, Win_bet_area: BET_AREA.BANKER, Win_bet_area_cnt: 10, Percentage: 0.1},
		{Total_hands: 10000, Win_bet_area: BET_AREA.PLAYER, Win_bet_area_cnt: 10, Percentage: 0.1},
		{Total_hands: 10000, Win_bet_area: BET_AREA.TIE, Win_bet_area_cnt: 10, Percentage: 0.1},
		{Total_hands: 10000, Win_bet_area: BET_AREA.BANKER_PAIR, Win_bet_area_cnt: 10, Percentage: 0.1},
		{Total_hands: 10000, Win_bet_area: BET_AREA.PLAYER_PAIR, Win_bet_area_cnt: 90, Percentage: 0.9},
	}

	txt := Instance_win_bet_area_stat_svg().Make_svg(stats)
	fmt.Println(txt)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
