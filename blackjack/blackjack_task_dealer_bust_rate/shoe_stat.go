/*
功能：blackjack任务-dealer爆牌率
说明：每靴牌统计
*/
package blackjack_task_dealer_bust_rate

import (
	"Odds/blackjack/define/CARD_TYPE"
	"Odds/common"
	"fmt"
)

// 每靴牌统计
type ShoeStat struct {
	shoe_index       int
	dealer_show_card byte           //dealer的明牌
	dealer_cards     []byte         //dealer的牌
	dealer_card_type CARD_TYPE.TYPE //dealer的牌型
	dealer_point     int            //dealer的点数
}

func (s *ShoeStat) String() string {
	str := fmt.Sprintf("[shoe_index:%d,", s.shoe_index)
	str += fmt.Sprintf("dealer_show_card:%s,", common.Card_2_sign(s.dealer_show_card))
	str += fmt.Sprintf("dealer_cards:%s,", common.Cards_2_string(s.dealer_cards))
	str += fmt.Sprintf("dealer_card_type:%s,dealer_point:%d]", s.dealer_card_type.String(), s.dealer_point)
	return str
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
