/*
功能：组合
说明：
*/
package algorithm

import (
	"Odds/common"
	"math/big"

	"github.com/poker-x-studio/x/xlog"
)

// Combo_card_type 按所需张数进行组合,返回组合牌型的列表
// hole_cards:手里的牌
// target_hole_cnt:需要手里的牌的张数
// board_cards:公共牌
// target_board_cnt:需要公共牌的张数
func Combo_card_type(hole_cards []byte, target_hole_cnt int, board_cards []byte, target_board_cnt int) []common.CardCtrl {
	if len(hole_cards) < target_hole_cnt {
		return nil
	}
	if len(board_cards) < target_board_cnt {
		return nil
	}
	all_combo_card_ctrls := make([]common.CardCtrl, 0)

	hole_combo := Combo_cards(hole_cards, target_hole_cnt)
	board_combo := Combo_cards(board_cards, target_board_cnt)

	//再次组合
	for i := 0; i < len(hole_combo); i++ {
		for j := 0; j < len(board_combo); j++ {

			new_card_ctrl := common.NewCardCtrlWithCards(hole_combo[i].Cards, -1)
			new_card_ctrl.Append_cards(board_combo[j].Cards)

			all_combo_card_ctrls = append(all_combo_card_ctrls, *new_card_ctrl)
		}
	}
	return all_combo_card_ctrls
}

// Combo_cards 组合牌
func Combo_cards(all_cards []byte, target_cnt int) []common.CardCtrl {
	if len(all_cards) < target_cnt {
		return nil
	}

	//组合个数
	combo_cnt := Combo_cnt(len(all_cards), target_cnt)

	total_card_ctrls := make([]common.CardCtrl, 0)

	for i := 0; i < len(all_cards); i++ {
		//临时
		tmp_card_ctrls := make([]common.CardCtrl, 0)

		new_card_ctrl := common.NewCardCtrlWithCard(all_cards[i], target_cnt)
		//递归组合
		_recursion_append_one_card(all_cards[i:], target_cnt, 1, new_card_ctrl, &tmp_card_ctrls)

		total_card_ctrls = append(total_card_ctrls, tmp_card_ctrls...)
	}

	total_len := int64(len(total_card_ctrls))
	if combo_cnt != total_len {
		xlog.New_entry("TASK_TAG").Errorf("Combo_cards(),combo_cnt:%d,total_len:%d,", combo_cnt, total_len)
		return nil
	}
	return total_card_ctrls
}

// _recursion_append_one_card 递归追加一张牌
func _recursion_append_one_card(all_cards []byte, target_cnt int, start_index int, cur_card_ctrl *common.CardCtrl, tmp_card_ctrls *[]common.CardCtrl) {
	if cur_card_ctrl.Is_full() { //牌控件充满
		*tmp_card_ctrls = append(*tmp_card_ctrls, *cur_card_ctrl)
		return
	}

	//开始裂变
	for i := start_index; i < len(all_cards); i++ {
		new_card_ctrl := common.NewCardCtrlWithCards(cur_card_ctrl.Cards, target_cnt)
		new_card_ctrl.Append_card(all_cards[i])

		_recursion_append_one_card(all_cards, target_cnt, i+1, new_card_ctrl, tmp_card_ctrls)
	}
}

// Permutation_cnt 排列数,数学方法计算排列数(从n中取k个数),排列有顺序
// P(n,k) = n!/(n-k)!
func Permutation_cnt(n int, k int) int64 {
	if n < k {
		return 0
	}
	tmp := new(big.Int)
	div := tmp.Div(Factorial(n), Factorial(n-k))
	return div.Int64()
}

// Combo_cnt 组合数,数学方法计算组合数(从n中取k个数),组合没有顺序
// C(n,k) = n!/((n-k)!*k!)
func Combo_cnt(n int, k int) int64 {
	if n < k {
		return 0
	}
	tmp := new(big.Int)
	mul := tmp.Mul(Factorial(n-k), Factorial(k))
	div := tmp.Div(Factorial(n), mul)
	return div.Int64()
}

// Factorial 阶乘
// 5!=5*4*3*2*1
func Factorial(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		tmp := big.NewInt(int64(i))
		result.Mul(result, tmp)
	}
	return result
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
