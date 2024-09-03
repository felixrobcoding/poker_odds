/*
功能：百家乐任务-统计庄家胜率
说明：流程控制
*/
package baccarat_task_winrate

import (
	"Odds/baccarat/define"
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/define/CARD_TYPE"
	"Odds/baccarat/logic"
	"Odds/baccarat/strategy_bet_amount"
	"Odds/baccarat/strategy_bet_area"
	"Odds/baccarat/strategy_bet_area/big_road"
	"Odds/baccarat/strategy_bet_area/suggestion"
	"Odds/baccarat/user_info"
	"Odds/common"
	"Odds/common/BET_AMOUNT_STRATEGY"
	"Odds/common/USER_TYPE"
	"Odds/common/algorithm"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/poker-x-studio/x/ximage"
)

const (
	PLAYER_INIT_CHIP = 10000 //
)

var shoe_index int //靴牌索引

type FlowControl struct {
	shoe_index          int                                    //靴牌索引
	shoe_cards          []byte                                 //牌靴里的牌
	shoe_cards_all      []byte                                 //牌靴里的牌
	deal_times          int                                    //发牌次数
	player              *user_info.UserInfo                    //闲家
	dealer              *user_info.UserInfo                    //庄家
	messages            []string                               //复盘信息
	win_bet_areas       [][]BET_AREA.TYPE                      //获胜区域
	max_bet_amount      int                                    //最大下注额
	bet_area_strategy   *strategy_bet_area.Strategy            //下注区域策略
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
	f.win_bet_areas = make([][]BET_AREA.TYPE, 0)

	f.bet_area_strategy = strategy_bet_area.NewStrategy()

	//f.bet_amount_strategy = strategy_bet_amount.NewBetAmountStrategy(BET_AMOUNT_STRATEGY.ALL_IN, PLAYER_INIT_CHIP)
	//f.bet_amount_strategy = strategy_bet_amount.NewBetAmountStrategy(BET_AMOUNT_STRATEGY.FIXED_AMOUNT, PLAYER_INIT_CHIP)
	//f.bet_amount_strategy = strategy_bet_amount.NewBetAmountStrategy(BET_AMOUNT_STRATEGY.MARTEGAL, PLAYER_INIT_CHIP)
	//f.bet_amount_strategy = strategy_bet_amount.NewBetAmountStrategy(BET_AMOUNT_STRATEGY.FIBONACCI, PLAYER_INIT_CHIP)
	//f.bet_amount_strategy = strategy_bet_amount.NewBetAmountStrategy(BET_AMOUNT_STRATEGY.KELLY, PLAYER_INIT_CHIP)
	f.bet_amount_strategy = strategy_bet_amount.NewBetAmountStrategy(BET_AMOUNT_STRATEGY.DANNY, PLAYER_INIT_CHIP)
}

// 洗牌
func (f *FlowControl) Shuffle() {
	f.shoe_index = shoe_index
	shoe_index++
	f.shoe_cards_all = make([]byte, 0)
	f.shoe_cards_all = algorithm.Shuffle_cards(DECKS)
	f.shoe_cards = make([]byte, len(f.shoe_cards_all))
	copy(f.shoe_cards, f.shoe_cards_all)
	f.deal_times = 0

	//测试
	//f.shoe_cards = []byte{0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03}
	//f.shoe_cards = []byte{0x01, 0x11, 0x21, 0x31, 0x01, 0x11, 0x21, 0x31, 0x01, 0x11, 0x21, 0x31, 0x01, 0x11, 0x21, 0x31}
}

func (f *FlowControl) Shuffle_from_outside(shoe_cards []byte) {
	f.shoe_index = shoe_index
	shoe_index++
	f.shoe_cards_all = make([]byte, len(shoe_cards))
	copy(f.shoe_cards_all, shoe_cards)
	f.shoe_cards = make([]byte, len(f.shoe_cards_all))
	copy(f.shoe_cards, f.shoe_cards_all)
	f.deal_times = 0
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

	//查询下注区域
	suggestion_bet_area := f.bet_area_strategy.Query_bet_area()
	//下注策略
	bet_amount, err := f.bet_amount_strategy.Query_bet_amount()
	if err != nil {
		msg := fmt.Sprintf("本局结束,deal_times:%d,shoe_card_cnt:%d,err:%s,", f.deal_times, shoe_card_cnt, err.Error())
		f.push_message(msg)
		return errors.New(msg)
	}
	f.deal_times++

	f.player.Deal(player_cards, suggestion_bet_area.Bet_area, bet_amount) //闲家押注
	f.dealer.Deal(dealer_cards, BET_AREA.ERROR, 0)

	player_point := logic.Points(f.player.Current_hand().Cards())
	dealer_point := logic.Points(f.dealer.Current_hand().Cards())
	shoe_card_cnt = len(f.shoe_cards)
	msg := fmt.Sprintf("发牌,shoe_card_cnt:%d,player_cards:%s[点数:%d],dealer_cards:%s[点数:%d],", shoe_card_cnt, common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point)
	f.push_message(msg)

	msg = fmt.Sprintf("%s,闲家押注区域:%s,下注额:%d,", suggestion_bet_area.Comment, suggestion_bet_area.Bet_area.String(), bet_amount)
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

		msg := fmt.Sprintf("比较结果:庄赢,player_cards:%s[点数%d],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,本局闲输赢:%.2f,本局庄输赢:%.2f,", common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), player_profit, dealer_profit)
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

		msg := fmt.Sprintf("比较结果:Tie,player_cards:%s[点数%d],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,本局闲输赢:%.2f,本局庄输赢:%.2f,", common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), player_profit, dealer_profit)
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

		msg := fmt.Sprintf("比较结果:闲赢,player_cards:%s[点数%d],dealer_cards:%s[点数%d],闲家筹码:%.2f,庄家筹码:%.2f,本局闲输赢:%.2f,本局庄输赢:%.2f,", common.Cards_2_string(player_cards), player_point, common.Cards_2_string(dealer_cards), dealer_point, f.player.Get_chip(), f.dealer.Get_chip(), player_profit, dealer_profit)
		f.push_message(msg)
	}
	f.win_bet_areas = append(f.win_bet_areas, win_bet_areas)
}

// 本轮结束
func (f *FlowControl) Round_end() {
	msg := fmt.Sprintf("=====本轮结束=====")
	f.push_message(msg)

	//输出信息
	f.dump_messages()

	//下注区域策略
	bet_area, bet_amount, win_bet_areas := f.player.Current_hand().Get_bet()
	//排序
	sort.SliceStable(win_bet_areas, func(i, j int) bool {
		return win_bet_areas[i] < win_bet_areas[j]
	})
	suggestion_feedback_node := &suggestion.FeedbackNode{
		Current_chip:       f.player.Get_chip(),
		Current_bet_area:   bet_area,
		Current_bet_amount: bet_amount,
		Result_area:        win_bet_areas[0], //只需要庄闲和,策略评估不需要庄对闲对
		Result_score:       f.player.Current_hand().Get_score(),
	}
	f.bet_area_strategy.Feedback_node_append(suggestion_feedback_node)

	//下注额策略
	feedback_node := &strategy_bet_amount.FeedbackNode{
		Current_chip: f.player.Get_chip(),
		Bet_amount:   bet_amount,
		Result_area:  win_bet_areas[0], //只需要庄闲和,策略评估不需要庄对闲对
		Result_score: f.player.Current_hand().Get_score(),
	}
	f.bet_amount_strategy.Feedback_node_append(feedback_node)

	//记录最大下注额
	if f.max_bet_amount < bet_amount {
		f.max_bet_amount = bet_amount
	}
}

// 游戏结束
func (f *FlowControl) Game_over() {
	//输出信息
	f.dump_messages()

	player_stat := f.player.Extract_user_stat()
	dealer_stat := f.dealer.Extract_user_stat()

	//输出svg,svg和dat文件保存在一起，方便比对调试
	data_path := f.Make_svg()
	//保存靴牌
	Save_file(f.shoe_cards_all, data_path)

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

// 输出信息
func (f *FlowControl) dump_messages() {
	for _, v := range f.messages {
		xlog_entry.Debugf("%s", v)
	}
}

// 输出svg
func (f *FlowControl) Make_svg() string {
	bigroad := f.bet_area_strategy.Query_big_road()

	instance_svg := big_road.Instance_big_road_svg()
	svg_content := instance_svg.Make_svg(bigroad, true)

	jpeg_filepath, svg_filepath, err := ximage.Svg_2_jpeg(svg_content)
	if err != nil {
		return "tmp.dat"
	}
	os.Remove(svg_filepath)
	fmt.Println(jpeg_filepath)

	return strings.Replace(jpeg_filepath, "jpeg", "dat", -1)
}

// 提取每靴牌的统计
func (f *FlowControl) Extract_shoe_stat() *ShoeStat {
	//下注策略
	option_min_bet, option_max_bet, bet_amount_strategy := f.bet_amount_strategy.Query_option()
	//大路统计
	bigroad := f.bet_area_strategy.Query_big_road()
	bigroad_stat := bigroad.Extract_bigroad_stat()
	//用户统计
	user_stat := f.player.Extract_user_stat()
	shoe_stat := &ShoeStat{
		shoe_index:          f.shoe_index,
		option_min_bet:      option_min_bet,
		option_max_bet:      option_max_bet,
		bet_amount_strategy: bet_amount_strategy,
		deal_times:          f.Deal_times(),
		max_bet_amount:      f.max_bet_amount,
		player_init_chip:    f.player.Get_init_chip(),
		player_chip:         f.player.Get_chip(),
		player_total_bets:   user_stat.Total_bets,
		player_lose_hands:   user_stat.Lose_hands,
		player_push_hands:   user_stat.Tie_hands,
		player_win_hands:    user_stat.Win_hands,
		win_bet_areas:       f.win_bet_areas,
		bigroad_stat:        bigroad_stat,
	}

	return shoe_stat
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
