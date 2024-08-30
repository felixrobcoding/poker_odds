/*
功能：blackjack任务-胜率统计
说明：流程控制
*/
package blackjack_task_winrate

import (
	"Odds/baccarat/strategy_bet_amount"
	"Odds/blackjack"
	"Odds/blackjack/define"
	"Odds/blackjack/define/ACTION_TYPE"
	"Odds/blackjack/define/CARD_TYPE"
	"Odds/blackjack/logic"
	"Odds/blackjack/strategy"
	"Odds/blackjack/user_info"
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
	shoe_index          int                                    //靴牌索引
	shoe_cards          []byte                                 //牌靴里的牌
	deal_times          int                                    //发牌次数
	player              *user_info.UserInfo                    //闲家
	dealer              *user_info.UserInfo                    //庄家
	messages            []string                               //复盘信息
	bet_amount_strategy strategy_bet_amount.IBetAmountStrategy //下注额策略
}

func NewFlowControl() *FlowControl {
	flow_control := &FlowControl{}
	flow_control.init()
	return flow_control
}

func (f *FlowControl) init() {
	f.player = user_info.NewUserInfo(USER_TYPE.PLAYER, PLAYER_INIT_CHIP)
	f.dealer = user_info.NewUserInfo(USER_TYPE.BANKER, 0)

	//f.bet_amount_strategy = strategy_bet_amount.NewBetAmountStrategy(BET_AMOUNT_STRATEGY.ALL_IN, PLAYER_INIT_CHIP)
	//f.bet_amount_strategy = strategy_bet_amount.NewBetAmountStrategy(BET_AMOUNT_STRATEGY.FIXED_AMOUNT, PLAYER_INIT_CHIP)
	//f.bet_amount_strategy = strategy_bet_amount.NewBetAmountStrategy(BET_AMOUNT_STRATEGY.MARTEGAL, PLAYER_INIT_CHIP)
	//f.bet_amount_strategy = strategy_bet_amount.NewBetAmountStrategy(BET_AMOUNT_STRATEGY.FIBONACCI, PLAYER_INIT_CHIP)
	//f.bet_amount_strategy = strategy_bet_amount.NewBetAmountStrategy(BET_AMOUNT_STRATEGY.KELLY, PLAYER_INIT_CHIP)
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
func (f *FlowControl) Round_begin_to_deal() error {
	f.messages = make([]string, 0)

	shoe_card_cnt := len(f.shoe_cards)
	if !f.Is_valid_shoe_cards() {
		msg := fmt.Sprintf("本局结束,shoe牌个数不足,deal_times:%d,shoe_card_cnt:%d,", f.deal_times, shoe_card_cnt)
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

	bet, err := f.bet_amount_strategy.Query_bet_amount() //下注额策略
	if err != nil {
		msg := fmt.Sprintf("本局结束,deal_times:%d,shoe_card_cnt:%d,err:%s,", f.deal_times, shoe_card_cnt, err.Error())
		f.push_message(msg)
		return errors.New(msg)
	}
	f.deal_times++

	f.player.Deal(player_cards, bet) //闲家押注
	f.dealer.Deal(dealer_cards, 0)

	_, player_points_txt := logic.Points(f.player.Current_hand().Cards())
	_, dealer_points_txt := logic.Points(f.dealer.Current_hand().Cards())
	shoe_card_cnt = len(f.shoe_cards)
	msg := fmt.Sprintf("发牌,shoe_card_cnt:%d,player_cards:%s[点数:%s],dealer_cards:%s[点数:%s],", shoe_card_cnt, common.Cards_2_string(player_cards), player_points_txt, common.Cards_2_string(dealer_cards), dealer_points_txt)
	f.push_message(msg)

	msg = fmt.Sprintf("闲家押注:%d,", f.player.Current_hand().Get_bet_amount())
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

// Player_turn 闲家操作
func (f *FlowControl) Player_turn() error {
	// 闲家叫牌,庄家一张明牌
	DEALER_CARD := f.dealer.Current_hand().Cards()[0]

	for {
		if f.player.Current_hand().Is_split() && f.player.Current_hand().Card_cnt() == 1 {
			//发一张牌
			card := f.Deal_1_card()
			f.player.Current_hand().Append_card(card)

			_, points_txt := logic.Points(f.player.Current_hand().Cards())
			msg := fmt.Sprintf("闲家 分牌之后,发第二张牌card:0x%02X,player_cards:%s[闲:%s],", card, common.Cards_2_string(f.player.Current_hand().Cards()), points_txt)
			f.push_message(msg)
		}

		action, point, dealer_value := strategy.Player_query_action(f.player.Current_hand().Cards(), DEALER_CARD)

		msg := fmt.Sprintf("查询动作,结果为:闲家 %s,player_cards:%s[闲key:%d 庄明牌:%s],", action.String(), common.Cards_2_string(f.player.Current_hand().Cards()), point, dealer_value)
		f.push_message(msg)

		if action == ACTION_TYPE.HIT { //要牌
			card := f.Deal_1_card()
			f.player.Current_hand().Append_card(card)

			_, points_txt := logic.Points(f.player.Current_hand().Cards())
			msg := fmt.Sprintf("闲家 %s,player_cards:%s[闲:%s 庄明牌:%s],", action.String(), common.Cards_2_string(f.player.Current_hand().Cards()), points_txt, dealer_value)
			f.push_message(msg)

		} else if action == ACTION_TYPE.DOUBLE_DOWN { //加倍停牌
			card := f.Deal_1_card()
			f.player.Current_hand().Append_card(card)
			f.player.Current_hand().Double_down()

			_, points_txt := logic.Points(f.player.Current_hand().Cards())
			msg := fmt.Sprintf("闲家 %s,player_cards:%s[闲:%s 庄明牌:%s],", action.String(), common.Cards_2_string(f.player.Current_hand().Cards()), points_txt, dealer_value)
			f.push_message(msg)

			break
		} else if action == ACTION_TYPE.STAND { //停牌

			_, points_txt := logic.Points(f.player.Current_hand().Cards())
			msg := fmt.Sprintf("闲家 %s,player_cards:%s[闲:%s 庄明牌:%s],", action.String(), common.Cards_2_string(f.player.Current_hand().Cards()), points_txt, dealer_value)
			f.push_message(msg)

			break
		} else if action == ACTION_TYPE.SPLIT { //分牌

			_, points_txt := logic.Points(f.player.Current_hand().Cards())
			msg := fmt.Sprintf("闲家 %s,player_cards:%s[闲:%s 庄明牌:%s],", action.String(), common.Cards_2_string(f.player.Current_hand().Cards()), points_txt, dealer_value)
			f.push_message(msg)

			right_card := f.Deal_1_card()
			right_cards, left_cards := f.player.Split_card(right_card)

			_, right_points_txt := logic.Points(right_cards)
			_, left_points_txt := logic.Points(left_cards)

			msg = fmt.Sprintf("闲家 %s,分牌后,right_cards:%s[闲:%s 庄明牌:%s],left_cards:%s[闲:%s 庄明牌:%s],", action.String(), common.Cards_2_string(right_cards), right_points_txt, dealer_value, common.Cards_2_string(left_cards), left_points_txt, dealer_value)
			f.push_message(msg)

			continue

		} else if action == ACTION_TYPE.SURRENDER { //投降

			_, points_txt := logic.Points(f.player.Current_hand().Cards())
			msg := fmt.Sprintf("闲家 %s,player_cards:%s[闲:%s 庄明牌:%s],", action.String(), common.Cards_2_string(f.player.Current_hand().Cards()), points_txt, dealer_value)
			f.push_message(msg)

			f.player.Current_hand().Surrender()

			break
		} else {
			panic("")
		}
	}

	points, points_txt := logic.Points(f.player.Current_hand().Cards())
	if points[0] >= define.POINT_BUST {
		msg := fmt.Sprintf("本手结束,闲家爆点,点数:%s,", points_txt)
		f.push_message(msg)
	}

	//如果还有下一手
	if err := f.player.Next_hand(); err == nil {
		f.Player_turn()
	}

	//全部爆点
	is, _, surrender_cnt, bust_cnt := f.player.Is_all_bust_or_surrender()
	if is {
		msg := ""
		if surrender_cnt == 0 {
			msg = fmt.Sprintf("本局结束,闲家爆点,")
		}
		if bust_cnt == 0 {
			msg = fmt.Sprintf("本局结束,闲家投降,")
		}
		f.push_message(msg)
		return errors.New(msg)
	}

	return nil
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

// Compare 比牌
func (f *FlowControl) Compare() {
	dealer_cards := f.dealer.Current_hand().Cards()
	dealer_points, _ := logic.Points(dealer_cards)
	dealer_point := logic.Dealer_pick_best_point(dealer_points)

	//1 blackjack
	f.player.Reset_hand_index()
	dealer_card_type := f.dealer.Current_hand().Card_type()
	if dealer_card_type == CARD_TYPE.BLACK_JACK { //庄家blackjack
		for i := 0; i < f.player.Hand_cnt(); i++ {
			player_card_type := f.player.Current_hand().Card_type()

			player_cards := f.player.Current_hand().Cards()
			_, player_points_txt := logic.Points(player_cards)

			if player_card_type == CARD_TYPE.BLACK_JACK { //闲家blackjack
				f.dealer.Update_score(0)
				f.player.Update_score(0)

				msg := fmt.Sprintf("庄家blackjack,闲家Blackjack,player_cards:%s[点数%s],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,本局闲输赢:%.2f,本局庄输赢:%.2f,", common.Cards_2_string(player_cards), player_points_txt, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), f.player.Get_profit(), f.dealer.Get_profit())
				f.push_message(msg)
			} else { //闲家非blackjack
				f.dealer.Update_score(1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.BLACK_JACK.Odds())
				f.player.Update_score(-1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.BLACK_JACK.Odds())

				msg := fmt.Sprintf("庄家blackjack,闲家非Blackjack,player_cards:%s[点数%s],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,本局闲输赢:%.2f,本局庄输赢:%.2f,", common.Cards_2_string(player_cards), player_points_txt, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), f.player.Get_profit(), f.dealer.Get_profit())
				f.push_message(msg)
			}

			f.player.Next_hand()
		}
	} else {
		for i := 0; i < f.player.Hand_cnt(); i++ {
			player_card_type := f.player.Current_hand().Card_type()

			player_cards := f.player.Current_hand().Cards()
			_, player_points_txt := logic.Points(player_cards)

			if player_card_type == CARD_TYPE.BLACK_JACK {
				f.dealer.Update_score(-1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.BLACK_JACK.Odds())
				f.player.Update_score(1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.BLACK_JACK.Odds())

				msg := fmt.Sprintf("闲家Blackjack,player_cards:%s[点数%s],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,本局闲输赢:%.2f,本局庄输赢:%.2f,", common.Cards_2_string(player_cards), player_points_txt, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), f.player.Get_profit(), f.dealer.Get_profit())
				f.push_message(msg)
			}
			f.player.Next_hand()
		}
	}

	//2 投降
	f.player.Reset_hand_index()
	for i := 0; i < f.player.Hand_cnt(); i++ {
		player_cards := f.player.Current_hand().Cards()
		_, player_points_txt := logic.Points(player_cards)
		if f.player.Current_hand().Is_surrender() {
			f.dealer.Update_score(1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.SURRENDER.Odds())
			f.player.Update_score(-1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.SURRENDER.Odds())
			msg := fmt.Sprintf("庄家赢,闲家投降,player_cards:%s[点数%s],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,本局闲输赢:%.2f,本局庄输赢:%.2f,", common.Cards_2_string(player_cards), player_points_txt, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), f.player.Get_profit(), f.dealer.Get_profit())
			f.push_message(msg)
		}
		f.player.Next_hand()
	}

	//3 爆点
	f.player.Reset_hand_index()
	for i := 0; i < f.player.Hand_cnt(); i++ {
		player_cards := f.player.Current_hand().Cards()
		player_points, player_points_txt := logic.Points(player_cards)
		if player_points[0] >= define.POINT_BUST { //闲家爆点
			f.dealer.Update_score(1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.POINT.Odds())
			f.player.Update_score(-1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.POINT.Odds())
			msg := fmt.Sprintf("庄家赢,闲家爆点,player_cards:%s[点数%s],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,本局闲输赢:%.2f,本局庄输赢:%.2f,", common.Cards_2_string(player_cards), player_points_txt, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), f.player.Get_profit(), f.dealer.Get_profit())
			f.push_message(msg)
		}
		f.player.Next_hand()
	}

	//4 其他
	f.player.Reset_hand_index()
	if dealer_point >= define.POINT_BUST { //庄家爆点
		for i := 0; i < f.player.Hand_cnt(); i++ {
			if f.player.Current_hand().Is_result() { //已经结算
				continue
			}
			player_cards := f.player.Current_hand().Cards()
			_, player_points_txt := logic.Points(player_cards)

			f.dealer.Update_score(-1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.POINT.Odds())
			f.player.Update_score(1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.POINT.Odds())

			msg := fmt.Sprintf("闲家赢,庄家爆点,player_cards:%s[点数%s],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,本局闲输赢:%.2f,本局庄输赢:%.2f,", common.Cards_2_string(player_cards), player_points_txt, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), f.player.Get_profit(), f.dealer.Get_profit())
			f.push_message(msg)

			f.player.Next_hand()
		}
	} else {
		for i := 0; i < f.player.Hand_cnt(); i++ {
			if f.player.Current_hand().Is_result() { //已经结算
				continue
			}
			player_cards := f.player.Current_hand().Cards()
			player_points, _ := logic.Points(player_cards)
			player_point := logic.Player_pick_best_point_to_compare(player_points)

			if dealer_point > player_point { //庄赢
				f.dealer.Update_score(1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.POINT.Odds())
				f.player.Update_score(-1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.POINT.Odds())

				msg := fmt.Sprintf("庄家赢,player_cards:%s[点数%d],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,本局闲输赢:%.2f,本局庄输赢:%.2f,", common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), f.player.Get_profit(), f.dealer.Get_profit())
				f.push_message(msg)

			} else if dealer_point == player_point { //push
				f.dealer.Update_score(0)
				f.player.Update_score(0)

				msg := fmt.Sprintf("Push,player_cards:%s[点数%d],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,本局闲输赢:%.2f,本局庄输赢:%.2f,", common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), f.player.Get_profit(), f.dealer.Get_profit())
				f.push_message(msg)

			} else { //闲赢
				f.dealer.Update_score(-1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.POINT.Odds())
				f.player.Update_score(1 * float64(f.player.Current_hand().Get_bet_amount()) * CARD_TYPE.POINT.Odds())

				msg := fmt.Sprintf("闲家赢,player_cards:%s[点数%d],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,本局闲输赢:%.2f,本局庄输赢:%.2f,", common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), f.player.Get_profit(), f.dealer.Get_profit())
				f.push_message(msg)
			}
			f.player.Next_hand()
		}
	}
}

// Round_end 本轮结束
func (f *FlowControl) Round_end() {
	msg := fmt.Sprintf("=====本轮结束=====")
	f.push_message(msg)

	//写入日志
	for _, v := range f.messages {
		xlog_entry.Debugf("%s", v)
	}

	//下注策略
	Current_bets := make([]int, 0)
	Game_scores := make([]float64, 0)
	f.player.Reset_hand_index()
	for i := 0; i < f.player.Hand_cnt(); i++ {
		Current_bets = append(Current_bets, f.player.Current_hand().Get_bet_amount())
		Game_scores = append(Game_scores, f.player.Current_hand().Get_score())
		f.player.Next_hand()
	}

	// result_node := &strategy_bet_amount.ResultNode{
	// 	Current_chip:   f.player.Get_chip(),
	// 	Current_bets:   Current_bets,
	// 	Current_scores: Game_scores,
	// }
	// f.bet_amount_strategy.Result_node_append(result_node)
}

// Game_over 游戏结束
func (f *FlowControl) Game_over() {
	player_stat := f.player.Extract_user_stat()
	dealer_stat := f.dealer.Extract_user_stat()

	xlog_entry.Debugf("=====游戏结束,总轮数:%d,shoe_card_cnt:%d=====", f.Deal_times(), len(f.shoe_cards))
	xlog_entry.Debugf("=====游戏结束,player_stat:%s=====", player_stat.String())
	xlog_entry.Debugf("=====游戏结束,dealer_stat:%s=====", dealer_stat.String())
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
	min_bet, max_bet, bet_amount_strategy := f.bet_amount_strategy.Query_option()
	user_stat := f.player.Extract_user_stat()
	shoe_stat := &ShoeStat{
		shoe_index:          f.shoe_index,
		min_bet:             min_bet,
		max_bet:             max_bet,
		bet_amount_strategy: bet_amount_strategy,
		deal_times:          f.Deal_times(),
		player_init_chip:    f.player.Get_init_chip(),
		player_chip:         f.player.Get_chip(),
		player_total_bets:   user_stat.Total_bets,
		player_lose_hands:   user_stat.Lose_hands,
		player_push_hands:   user_stat.Push_hands,
		player_win_hands:    user_stat.Win_hands,
	}

	return shoe_stat
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
