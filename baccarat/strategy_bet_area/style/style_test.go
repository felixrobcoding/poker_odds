/*
功能：测试
说明：
*/
package style

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/strategy_bet_area/suggestion"
	"fmt"
	"testing"
)

func Test_big_road(t *testing.T) {
	nodes := []*suggestion.FeedbackNode{
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.BANKER},
	}
	win_bet_areas := make([]BET_AREA.TYPE, 0)
	for i := 0; i < len(nodes); i++ {
		win_bet_areas = append(win_bet_areas, nodes[i].Result_area)
	}

	//big_road := big_road.NewBigRoad(win_bet_areas)
	//fmt.Println(big_road.Col_cnt())
}

func Test_long(t *testing.T) {
	nodes := []*suggestion.FeedbackNode{
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
	}
	is, suggestion := check_long_style(nodes)
	if is {
		fmt.Println(suggestion.String())
	}
}

func Test_single_jump(t *testing.T) {
	nodes := []*suggestion.FeedbackNode{
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.BANKER},
	}
	is, suggestion := check_single_jump_style(nodes)
	if is {
		fmt.Println(suggestion.String())
	}
}

func Test_double_jump1(t *testing.T) {
	nodes := []*suggestion.FeedbackNode{
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.PLAYER},
	}
	is, suggestion := check_double_jump_style(nodes)
	if is {
		fmt.Println(suggestion.String())
	}
}

func Test_double_jump2(t *testing.T) {
	nodes1 := []*suggestion.FeedbackNode{
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.BANKER},
	}
	nodes2 := []*suggestion.FeedbackNode{
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.PLAYER},
	}

	is, suggestion := check_double_jump_style(nodes1)
	if is {
		fmt.Println(suggestion.String())
	}
	is, suggestion = check_double_jump_style(nodes2)
	if is {
		fmt.Println(suggestion.String())
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
