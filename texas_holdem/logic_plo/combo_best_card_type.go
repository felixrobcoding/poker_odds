/*
功能：组合最佳牌型
说明：
*/
package logic_plo

import (
	"github.com/felixrobcoding/poker_oddscommon/COMPARE_TYPE"
	"github.com/felixrobcoding/poker_oddscommon/algorithm"
	"github.com/felixrobcoding/poker_oddstexas_holdem/define"
	"github.com/felixrobcoding/poker_oddstexas_holdem/logic"
)

// Combo_best_card_type 组合最佳牌型
// 手里2张牌,公共3张牌组合最佳牌型
func Combo_best_card_type(hole_cards []byte, board_cards []byte) *logic.AnalyseItem {
	if len(hole_cards) < define.HOLE_CNT_IN_COMBO {
		return nil
	}
	if len(board_cards) < define.BOARD_CNT_IN_COMBO {
		return nil
	}

	//组合排列
	card_ctrls := algorithm.Combo_card_type(hole_cards, define.HOLE_CNT_IN_COMBO, board_cards, define.BOARD_CNT_IN_COMBO)
	if card_ctrls == nil {
		return nil
	}

	//选择最佳组合排列
	analyse_items := make([]*logic.AnalyseItem, 0)
	for i := 0; i < len(card_ctrls); i++ {
		item := logic.NewAnalyseItem(card_ctrls[i].Cards)
		analyse_items = append(analyse_items, item)
	}

	max_item := analyse_items[0]
	for i := 1; i < len(analyse_items); i++ {
		if ct := logic.Compare(max_item, analyse_items[i]); ct == COMPARE_TYPE.BIGGER {
			max_item = analyse_items[i]
		}
	}

	return max_item
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
