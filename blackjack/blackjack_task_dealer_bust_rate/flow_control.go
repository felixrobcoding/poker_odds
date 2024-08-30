/*
功能：blackjack任务-dealer爆牌率
说明：流程控制
*/
package blackjack_task_dealer_bust_rate

import (
	"Odds/blackjack"
	"Odds/blackjack/define"
	"Odds/blackjack/define/ACTION_TYPE"
	"Odds/blackjack/define/CARD_TYPE"
	"Odds/blackjack/logic"
	"Odds/blackjack/user_info"
	"Odds/common"
	"Odds/common/USER_TYPE"
	"Odds/common/algorithm"
	"errors"
	"fmt"
)

const ()

var shoe_index int //靴牌索引

type FlowControl struct {
	shoe_index int                 //靴牌索引
	shoe_cards []byte              //牌靴里的牌
	deal_times int                 //发牌次数
	show_card  byte                //明牌
	player     *user_info.UserInfo //闲家
	dealer     *user_info.UserInfo //庄家
	messages   []string            //复盘信息
}

func NewFlowControl() *FlowControl {
	flow_control := &FlowControl{}
	flow_control.init()
	return flow_control
}

func (f *FlowControl) init() {
	f.player = user_info.NewUserInfo(USER_TYPE.PLAYER, 0)
	f.dealer = user_info.NewUserInfo(USER_TYPE.BANKER, 0)
}

// Shuffle 洗牌
func (f *FlowControl) Shuffle() {
	f.shoe_index = shoe_index
	shoe_index++
	f.shoe_cards = make([]byte, 0)
	f.shoe_cards = algorithm.Shuffle_cards(blackjack.DECKS)
	f.deal_times = 0

	//测试
	//f.shoe_cards = []byte{0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03}
	//f.shoe_cards = []byte{0x01, 0x11, 0x21, 0x31, 0x01, 0x11, 0x21, 0x31, 0x01, 0x11, 0x21, 0x31, 0x01, 0x11, 0x21, 0x31}
}

// Round_begin_to_deal 发牌
func (f *FlowControl) Round_begin_to_deal(show_card byte) error {
	f.messages = make([]string, 0)

	f.show_card = show_card

	shoe_card_cnt := len(f.shoe_cards)
	if !f.Is_valid_shoe_cards() {
		msg := fmt.Sprintf("本局结束,shoe牌个数不足,deal_times:%d,shoe_card_cnt:%d,", f.deal_times, shoe_card_cnt)
		f.push_message(msg)
		return errors.New(msg)
	}

	//从明牌开始发牌
	index := -1
	for index = 0; index < shoe_card_cnt; index++ {
		if f.shoe_cards[index] == show_card {
			break
		}
	}
	if index >= shoe_card_cnt {
		panic("")
	}
	f.shoe_cards = f.shoe_cards[index+1:]

	dealer_cards := make([]byte, 0)
	//发第一张牌
	dealer_cards = append(dealer_cards, show_card)
	//发第二张牌
	dealer_cards = append(dealer_cards, f.Deal_1_card())

	f.dealer.Deal(dealer_cards, 0)

	f.deal_times++

	_, dealer_points_txt := logic.Points(f.dealer.Current_hand().Cards())
	shoe_card_cnt = len(f.shoe_cards)
	msg := fmt.Sprintf("发牌,shoe_card_cnt:%d,dealer_cards:%s[点数:%s],", shoe_card_cnt, common.Cards_2_string(dealer_cards), dealer_points_txt)
	f.push_message(msg)

	return nil
}

// Is_valid_shoe_cards 剩下的shoe牌是否有效
func (f *FlowControl) Is_valid_shoe_cards() bool {
	shoe_card_cnt := len(f.shoe_cards)
	if shoe_card_cnt <= define.REMAIN_CARD_CNT {
		return false
	}
	return true
}

// Check_blackjack 校验blackjack牌型
func (f *FlowControl) Check_blackjack() error {
	player_card_type := f.player.Current_hand().Card_type()
	dealer_card_type := f.dealer.Current_hand().Card_type()

	if (player_card_type == CARD_TYPE.BLACK_JACK) && (dealer_card_type == CARD_TYPE.BLACK_JACK) {
		msg := fmt.Sprintf("本局结束,庄家闲家都Blackjack")
		f.push_message(msg)
		return errors.New(msg)
	}

	if dealer_card_type == CARD_TYPE.BLACK_JACK {
		msg := fmt.Sprintf("本局结束,庄家Blackjack")
		f.push_message(msg)
		return errors.New(msg)
	}

	if player_card_type == CARD_TYPE.BLACK_JACK {
		msg := fmt.Sprintf("本局结束,闲家Blackjack")
		f.push_message(msg)
		return errors.New(msg)
	}

	return nil
}

// Deal_1_card 发一张牌
func (f *FlowControl) Deal_1_card() byte {
	if (f.shoe_cards == nil) || (len(f.shoe_cards) < 1) {
		panic("")
	}
	card := f.shoe_cards[0]
	f.shoe_cards = f.shoe_cards[1:]
	return card
}

// Dealer_turn 庄家操作
func (f *FlowControl) Dealer_turn() error {

	for {
		//点数
		points, _ := logic.Points(f.dealer.Current_hand().Cards())
		point := logic.Dealer_pick_best_point(points)
		if point <= define.POINT_16 { //要牌
			card := f.Deal_1_card()
			f.dealer.Current_hand().Append_card(card)

			_, points_txt := logic.Points(f.dealer.Current_hand().Cards())
			msg := fmt.Sprintf("庄家 %s,dealer_cards:%s[点数:%s],", ACTION_TYPE.HIT.String(), common.Cards_2_string(f.dealer.Current_hand().Cards()), points_txt)
			f.push_message(msg)

		} else { //停牌

			_, points_txt := logic.Points(f.dealer.Current_hand().Cards())
			msg := fmt.Sprintf("庄家 %s,dealer_cards:%s[点数:%s],", ACTION_TYPE.STAND.String(), common.Cards_2_string(f.dealer.Current_hand().Cards()), points_txt)
			f.push_message(msg)

			break
		}
	}

	points, points_txt := logic.Points(f.dealer.Current_hand().Cards())
	if points[0] >= define.POINT_BUST {
		msg := fmt.Sprintf("本局结束,庄家爆点,点数:%s", points_txt)
		f.push_message(msg)
		return errors.New(msg)
	}
	return nil
}

// Deal_times 发牌次数
func (f *FlowControl) Deal_times() int {
	return f.deal_times
}

// push_message 复盘信息
func (f *FlowControl) push_message(txt string) {
	msg := fmt.Sprintf("shoe_index:%d,轮数:%d,%d,%s", f.shoe_index, f.Deal_times(), len(f.messages), txt)
	f.messages = append(f.messages, msg)
}

// Extract_shoe_stat 提取每靴牌的统计
func (f *FlowControl) Extract_shoe_stat() *ShoeStat {
	points, _ := logic.Points(f.dealer.Current_hand().Cards())

	shoe_stat := &ShoeStat{
		shoe_index:       f.shoe_index,
		dealer_show_card: f.show_card,
		dealer_cards:     f.dealer.Current_hand().Cards(),
		dealer_card_type: f.dealer.Current_hand().Card_type(),
		dealer_point:     logic.Dealer_pick_best_point(points),
	}

	return shoe_stat
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
