/*
功能：百家乐任务-统计庄家/闲家点数分布
说明：流程控制
*/
package baccarat_task_points

import (
	"Odds/baccarat/define"
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/define/CARD_TYPE"
	"Odds/baccarat/logic"
	"Odds/baccarat/user_info"
	"Odds/common"
	"Odds/common/USER_TYPE"
	"Odds/common/algorithm"
	"errors"
	"fmt"
)

const (
	PLAYER_INIT_CHIP = 10000
)

var shoe_index int //靴牌索引

type FlowControl struct {
	shoe_index      int                 //靴牌索引
	shoe_cards      []byte              //牌靴里的牌
	deal_times      int                 //发牌次数
	player          *user_info.UserInfo //闲家
	dealer          *user_info.UserInfo //庄家
	messages        []string            //复盘信息
	player_points   []int               //闲家点数
	dealer_points   []int               //庄家点数
	player_pair_cnt int                 //闲对
	dealer_pair_cnt int                 //庄对
}

func NewFlowControl() *FlowControl {
	flow_control := &FlowControl{}
	flow_control.init()
	return flow_control
}

func (f *FlowControl) init() {
	f.player = user_info.NewUserInfo(USER_TYPE.PLAYER, PLAYER_INIT_CHIP)
	f.dealer = user_info.NewUserInfo(USER_TYPE.BANKER, 0)

	f.player_points = make([]int, 0)
	f.dealer_points = make([]int, 0)
}

// 洗牌
func (f *FlowControl) Shuffle() {
	f.shoe_index = shoe_index
	shoe_index++
	f.shoe_cards = make([]byte, 0)
	f.shoe_cards = algorithm.Shuffle_cards(DECKS)
	f.deal_times = 0

	//测试
	//f.shoe_cards = []byte{0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03}
	//f.shoe_cards = []byte{0x01, 0x11, 0x21, 0x31, 0x01, 0x11, 0x21, 0x31, 0x01, 0x11, 0x21, 0x31, 0x01, 0x11, 0x21, 0x31}
}

// 发牌
func (f *FlowControl) Round_begin_to_deal() error {
	f.messages = make([]string, 0)

	shoe_card_cnt := len(f.shoe_cards)
	if !f.Is_valid_shoe_cards() {
		msg := fmt.Sprintf("本局结束,shoe牌不足,deal_times:%d,shoe_card_cnt:%d,", f.deal_times, shoe_card_cnt)
		f.push_message(msg)
		return errors.New(msg)
	}

	player_cards := make([]byte, 0)
	dealer_cards := make([]byte, 0)

	//发第一张牌
	player_cards = append(player_cards, f.Deal_1_card())
	dealer_cards = append(dealer_cards, f.Deal_1_card())

	//发第二张牌
	player_cards = append(player_cards, f.Deal_1_card())
	dealer_cards = append(dealer_cards, f.Deal_1_card())

	//闲对
	if logic.Card_type(player_cards) == CARD_TYPE.PAIR {
		f.player_pair_cnt++
	}
	//庄对
	if logic.Card_type(dealer_cards) == CARD_TYPE.PAIR {
		f.dealer_pair_cnt++
	}

	f.deal_times++

	f.player.Deal(player_cards, BET_AREA.PLAYER, 0) //闲家押注
	f.dealer.Deal(dealer_cards, BET_AREA.BANKER, 0)

	player_point := logic.Points(f.player.Current_hand().Cards())
	dealer_point := logic.Points(f.dealer.Current_hand().Cards())
	shoe_card_cnt = len(f.shoe_cards)
	msg := fmt.Sprintf("发牌,shoe_card_cnt:%d,player_cards:%s[点数:%d],dealer_cards:%s[点数:%d],", shoe_card_cnt, common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point)
	f.push_message(msg)

	return nil
}

// 剩下的shoe牌是否有效
func (f *FlowControl) Is_valid_shoe_cards() bool {
	shoe_card_cnt := len(f.shoe_cards)
	if shoe_card_cnt <= define.REMAIN_CARD_CNT {
		return false
	}
	return true
}

// 校验natural牌型
func (f *FlowControl) Check_natural() error {
	player_card_type := f.player.Current_hand().Card_type()
	dealer_card_type := f.dealer.Current_hand().Card_type()

	if (player_card_type == CARD_TYPE.NATURAL) && (dealer_card_type == CARD_TYPE.NATURAL) {
		msg := fmt.Sprintf("本局结束,庄家闲家都是Natural")
		f.push_message(msg)
		return errors.New(msg)
	}

	if dealer_card_type == CARD_TYPE.NATURAL {
		msg := fmt.Sprintf("本局结束,庄家Natural")
		f.push_message(msg)
		return errors.New(msg)
	}

	if player_card_type == CARD_TYPE.NATURAL {
		msg := fmt.Sprintf("本局结束,闲家Natural")
		f.push_message(msg)
		return errors.New(msg)
	}
	return nil
}

// 发一张牌
func (f *FlowControl) Deal_1_card() byte {
	if (f.shoe_cards == nil) || (len(f.shoe_cards) < 1) {
		panic("")
	}
	card := f.shoe_cards[0]
	f.shoe_cards = f.shoe_cards[1:]
	return card
}

// 闲家操作
func (f *FlowControl) Player_turn() error {

	if logic.Draw_card_for_player(f.player.Current_hand().Cards(), f.dealer.Current_hand().Cards()) { //闲家补牌
		draw_card := f.Deal_1_card()
		f.player.Current_hand().Append_card(draw_card)

		player_cards := f.player.Current_hand().Cards()
		player_point := logic.Points(player_cards)
		dealer_cards := f.dealer.Current_hand().Cards()
		dealer_point := logic.Points(dealer_cards)

		msg := fmt.Sprintf("闲家补牌:0x%02X,player_cards:%s[点数:%d],dealer_cards:%s[点数:%d],", draw_card, common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point)
		f.push_message(msg)

	} else { //闲家不补牌
		player_cards := f.player.Current_hand().Cards()
		player_point := logic.Points(player_cards)
		dealer_cards := f.dealer.Current_hand().Cards()
		dealer_point := logic.Points(dealer_cards)

		msg := fmt.Sprintf("闲家不补牌,player_cards:%s[点数:%d],dealer_cards:%s[点数:%d],", common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point)
		f.push_message(msg)
	}
	return nil
}

// 庄家操作
func (f *FlowControl) Dealer_turn() error {

	if logic.Draw_card_for_dealer(f.player.Current_hand().Cards(), f.dealer.Current_hand().Cards()) { //庄家补牌
		draw_card := f.Deal_1_card()
		f.dealer.Current_hand().Append_card(draw_card)

		player_cards := f.player.Current_hand().Cards()
		player_point := logic.Points(player_cards)
		dealer_cards := f.dealer.Current_hand().Cards()
		dealer_point := logic.Points(dealer_cards)

		msg := fmt.Sprintf("庄家补牌:0x%2X,player_cards:%s[点数:%d],dealer_cards:%s[点数:%d],", draw_card, common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point)
		f.push_message(msg)

	} else { //庄家不补牌
		player_cards := f.player.Current_hand().Cards()
		player_point := logic.Points(player_cards)
		dealer_cards := f.dealer.Current_hand().Cards()
		dealer_point := logic.Points(dealer_cards)

		msg := fmt.Sprintf("庄家不补牌,player_cards:%s[点数:%d],dealer_cards:%s[点数:%d],", common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point)
		f.push_message(msg)
	}
	return nil
}

// 比牌
func (f *FlowControl) Compare() {
	dealer_cards := f.dealer.Current_hand().Cards()
	dealer_point := logic.Points(dealer_cards)
	dealer_card_type := f.dealer.Current_hand().Card_type()

	player_cards := f.player.Current_hand().Cards()
	player_point := logic.Points(player_cards)
	player_card_type := f.player.Current_hand().Card_type()

	win_bet_areas := make([]BET_AREA.TYPE, 0)
	if dealer_card_type == CARD_TYPE.PAIR { //庄对
		win_bet_areas = append(win_bet_areas, BET_AREA.BANKER_PAIR)
	}
	if player_card_type == CARD_TYPE.PAIR { //闲对
		win_bet_areas = append(win_bet_areas, BET_AREA.PLAYER_PAIR)
	}

	win_bet_area := BET_AREA.ERROR
	dealer_profit := 0.0
	player_profit := 0.0

	if dealer_point > player_point { //庄赢
		win_bet_area = BET_AREA.BANKER
		win_bet_areas = append(win_bet_areas, win_bet_area)

		bet_area, bet_amount, _ := f.player.Current_hand().Get_bet()
		if bet_area == win_bet_area {
			player_profit = 1 * float64(bet_amount) * win_bet_area.Odds()
			dealer_profit = -1 * float64(bet_amount) * win_bet_area.Odds()
		} else {
			player_profit = -1 * float64(bet_amount)
			dealer_profit = 1 * float64(bet_amount)
		}
		f.dealer.Update_score(dealer_profit, win_bet_areas)
		f.player.Update_score(player_profit, win_bet_areas)

		msg := fmt.Sprintf("庄赢,player_cards:%s[点数%d],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,闲输赢:%.2f,庄输赢:%.2f,", common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), f.player.Get_profit(), f.dealer.Get_profit())
		f.push_message(msg)

	} else if dealer_point == player_point { //tie
		win_bet_area = BET_AREA.TIE
		win_bet_areas = append(win_bet_areas, win_bet_area)

		bet_area, bet_amount, _ := f.player.Current_hand().Get_bet()
		if bet_area == win_bet_area {
			player_profit = 1 * float64(bet_amount) * win_bet_area.Odds()
			dealer_profit = -1 * float64(bet_amount) * win_bet_area.Odds()
		} else {
			player_profit = 0
			dealer_profit = 0
		}
		f.dealer.Update_score(dealer_profit, win_bet_areas)
		f.player.Update_score(player_profit, win_bet_areas)

		msg := fmt.Sprintf("Tie,player_cards:%s[点数%d],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,闲输赢:%.2f,庄输赢:%.2f,", common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), f.player.Get_profit(), f.dealer.Get_profit())
		f.push_message(msg)

	} else { //闲赢
		win_bet_area = BET_AREA.PLAYER
		win_bet_areas = append(win_bet_areas, win_bet_area)

		bet_area, bet_amount, _ := f.player.Current_hand().Get_bet()
		if bet_area == win_bet_area {
			player_profit = 1 * float64(bet_amount) * win_bet_area.Odds()
			dealer_profit = -1 * float64(bet_amount) * win_bet_area.Odds()
		} else {
			player_profit = -1 * float64(bet_amount)
			dealer_profit = 1 * float64(bet_amount)
		}
		f.dealer.Update_score(dealer_profit, win_bet_areas)
		f.player.Update_score(player_profit, win_bet_areas)

		msg := fmt.Sprintf("闲赢,player_cards:%s[点数%d],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,闲输赢:%.2f,庄输赢:%.2f,", common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), f.player.Get_profit(), f.dealer.Get_profit())
		f.push_message(msg)
	}
	f.player_points = append(f.player_points, player_point)
	f.dealer_points = append(f.dealer_points, dealer_point)
}

// 本轮结束
func (f *FlowControl) Round_end() {
	msg := fmt.Sprintf("=====本轮结束=====")
	f.push_message(msg)

	//写入日志
	for _, v := range f.messages {
		xlog_entry.Debugf("%s", v)
	}
}

// 游戏结束
func (f *FlowControl) Game_over() {
	player_stat := f.player.Extract_user_stat()
	dealer_stat := f.dealer.Extract_user_stat()

	xlog_entry.Debugf("=====游戏结束,总轮数:%d,shoe_card_cnt:%d=====", f.Deal_times(), len(f.shoe_cards))
	xlog_entry.Debugf("=====游戏结束,player_stat:%s=====", player_stat.String())
	xlog_entry.Debugf("=====游戏结束,dealer_stat:%s=====", dealer_stat.String())
}

// 发牌次数
func (f *FlowControl) Deal_times() int {
	return f.deal_times
}

// 复盘信息
func (f *FlowControl) push_message(txt string) {
	msg := fmt.Sprintf("shoe_index:%d,轮数:%d,%d,%s", f.shoe_index, f.Deal_times(), len(f.messages), txt)
	f.messages = append(f.messages, msg)
}

// 提取每靴牌的统计
func (f *FlowControl) Extract_shoe_stat() *ShoeStat {
	if len(f.player_points) != len(f.dealer_points) {
		panic("error")
	}

	//用户统计
	shoe_stat := &ShoeStat{
		shoe_index:      f.shoe_index,
		deal_times:      f.Deal_times(),
		player_points:   f.player_points,
		dealer_points:   f.dealer_points,
		player_pair_cnt: f.player_pair_cnt,
		dealer_pair_cnt: f.dealer_pair_cnt,
	}

	return shoe_stat
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
