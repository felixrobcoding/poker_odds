/*
功能：测试
说明：
*/
package svg_utils

import (
	"fmt"
	"github.com/felixrobcoding/poker_oddstexas_holdem/define/CARD_TYPE"
	"testing"
)

func Test_card_type_stat1(t *testing.T) {
	stats := []*CardTypeStat{
		{Total_run_times: 10000, Deal_card_cnt: 4, Type: CARD_TYPE.ONE_PAIR, Type_cnt: 10, Percentage: 0.1},
		{Total_run_times: 10000, Deal_card_cnt: 5, Type: CARD_TYPE.FLUSH, Type_cnt: 90, Percentage: 0.9},
	}

	txt := Instance_card_type_stat_svg().Make_svg(TXT_HEADER_TEXAS_HOLDEM, stats)
	fmt.Println(txt)
}

func Test_card_type_stat2(t *testing.T) {
	stats := []*CardTypeStat{
		{Total_run_times: 10000, Deal_card_cnt: 4, Type: CARD_TYPE.ONE_PAIR, Type_cnt: 10, Percentage: 0.1},
		{Total_run_times: 10000, Deal_card_cnt: 5, Type: CARD_TYPE.FLUSH, Type_cnt: 90, Percentage: 0.9},
	}

	txt := Instance_card_type_stat_svg().Make_svg(TXT_HEADER_UTLIMATE, stats)
	fmt.Println(txt)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
