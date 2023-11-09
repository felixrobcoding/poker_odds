/*
功能：测试
说明：
*/
package algorithm

import (
	"fmt"
	"math/big"
	"testing"
)

func TestXxx1(t *testing.T) {
	all_cards := []byte{0x00, 0x01, 0x02, 0x03, 0x04}
	const target_cnt int = 4
	card_ctrls := Combo_cards(all_cards, target_cnt)
	for i := 0; i < len(card_ctrls); i++ {
		fmt.Println(card_ctrls[i].String())
	}
	fmt.Println("end")
}

func TestXxx2(t *testing.T) {
	hole_cards := []byte{0x07, 0x08, 0x09, 0x0A, 0x0B}
	const target_hole_cnt int = 2
	board_cards := []byte{0x00, 0x01, 0x02, 0x03, 0x04}
	const target_board_cnt int = 3

	card_ctrls := Combo_card_type(hole_cards, target_hole_cnt, board_cards, target_board_cnt)
	for i := 0; i < len(card_ctrls); i++ {
		fmt.Println(card_ctrls[i].String())
	}

	//组合数
	player_cnt := 2
	card_cnt_per_player := 2
	combo_cnt := Combo_cnt(52-card_cnt_per_player*player_cnt, 5)
	fmt.Println(combo_cnt)
	fmt.Println("end")
}

func Test_factorial(t *testing.T) {
	f := Factorial(100)
	fmt.Println(f)
}

func Test_big(t *testing.T) {
	a1 := big.NewInt(1)
	a2 := big.NewInt(3)
	a1.Add(a1, a2)
	fmt.Println(a1.Int64())
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
