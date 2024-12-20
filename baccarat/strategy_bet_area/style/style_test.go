/*
功能：测试
说明：
*/
package style

import (
	"fmt"
	"github.com/felixrobcoding/poker_oddsbaccarat/define/BET_AREA"
	"github.com/felixrobcoding/poker_oddsbaccarat/strategy_bet_area/suggestion"
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
	is, suggestion := _check_single_jump_style(nodes)
	if is {
		fmt.Println(suggestion.String())
	}
}

func Test_double_jump_half_col_1(t *testing.T) {
	nodes := []*suggestion.FeedbackNode{
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.PLAYER},
	}
	is, suggestion := _check_double_jump_style(nodes)
	if is {
		fmt.Println(suggestion.String())
	}
}

func Test_double_jump_half_col_2(t *testing.T) {
	nodes := []*suggestion.FeedbackNode{
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.BANKER},
	}

	is, suggestion := _check_double_jump_style(nodes)
	if is {
		fmt.Println(suggestion.String())
	}
}

func Test_double_jump_full_col_1(t *testing.T) {
	nodes := []*suggestion.FeedbackNode{
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.PLAYER},
	}
	is, suggestion := _check_double_jump_style(nodes)
	if is {
		fmt.Println(suggestion.String())
	}
}
func Test_3_jump_half_col_1(t *testing.T) {
	nodes := []*suggestion.FeedbackNode{
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.PLAYER},
	}
	is, suggestion := _check_3_jump_style(nodes)
	if is {
		fmt.Println(suggestion.String())
	}
}

func Test_3_jump_half_col_2(t *testing.T) {
	nodes := []*suggestion.FeedbackNode{
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.BANKER},
	}
	is, suggestion := _check_3_jump_style(nodes)
	if is {
		fmt.Println(suggestion.String())
	}
}

func Test_3_jump_full_col_1(t *testing.T) {
	nodes := []*suggestion.FeedbackNode{
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.BANKER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.PLAYER},
		{Result_area: BET_AREA.PLAYER},
	}
	is, suggestion := _check_3_jump_style(nodes)
	if is {
		fmt.Println(suggestion.String())
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
