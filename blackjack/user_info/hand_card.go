/*
功能：手牌
说明：
*/
package user_info

import (
	"Odds/blackjack/define/ACTION_TYPE"
	"Odds/blackjack/define/CARD_TYPE"
	"Odds/blackjack/define/HAND_TYPE"
	"Odds/blackjack/logic"
	"Odds/common/GAME_RESULT"
)

type HandCard struct {
	cards       []byte
	game_result GAME_RESULT.TYPE //手牌结果
	bet         int              //下注额
	score       float64          //得分
	action      int              //记录 分牌/投降
}

func NewHandCard() *HandCard {
	hand := HandCard{}
	hand.init()
	return &hand
}

// 初始化
func (h *HandCard) init() {
	h.cards = make([]byte, 0)
	h.game_result = GAME_RESULT.ERROR
}

func (h *HandCard) Cards() []byte {
	return h.cards
}

func (h *HandCard) Card_cnt() int {
	return len(h.cards)
}

// 追加
func (h *HandCard) Append_card(card byte) {
	if h.cards == nil {
		panic("")
	}
	h.cards = append(h.cards, card)

	//计算hand型
	h.cal_hand_type()
	//计算牌型
	h.cal_card_type()
}

func (h *HandCard) Append_cards(cards []byte) {
	if h.cards == nil {
		panic("")
	}
	h.cards = append(h.cards, cards...)

	//计算hand型
	h.cal_hand_type()
	//计算牌型
	h.cal_card_type()
}

// 获取hand型
func (h *HandCard) Hand_type() HAND_TYPE.TYPE {
	return h.cal_hand_type()
}

// 获取牌型
func (h *HandCard) Card_type() CARD_TYPE.TYPE {
	return h.cal_card_type()
}

// 计算hand型
func (h *HandCard) cal_hand_type() HAND_TYPE.TYPE {
	return logic.Analyse_hand_type(h.cards)
}

// 计算牌型
func (h *HandCard) cal_card_type() CARD_TYPE.TYPE {
	card_type := logic.Analyse_card_type(h.cards)
	if (card_type == CARD_TYPE.BLACK_JACK) && (h.Is_split()) { //分牌之后不能是Blackjack
		return CARD_TYPE.POINT
	}
	return card_type
}

// 是否分牌
func (h *HandCard) Is_split() bool {
	return (h.action & int(ACTION_TYPE.SPLIT)) > 0
}

// 投降
func (h *HandCard) Surrender() {
	h.action |= int(ACTION_TYPE.SURRENDER)
}

// 是否投降
func (h *HandCard) Is_surrender() bool {
	return (h.action & int(ACTION_TYPE.SURRENDER)) > 0
}

// 设置下注额
func (h *HandCard) Set_bet(bet int) {
	h.bet = bet
}

// 获取下注额
func (h *HandCard) Get_bet() int {
	return h.bet
}

// 获取积分
func (h *HandCard) Get_score() float64 {
	return h.score
}

// 更新积分
func (h *HandCard) update_score(score float64) (game_result GAME_RESULT.TYPE) {
	game_result = GAME_RESULT.PUSH
	if score > 0 {
		game_result = GAME_RESULT.WIN
	} else if score < 0 {
		game_result = GAME_RESULT.LOSE
	}

	h.game_result = game_result
	h.score = score
	return
}

// 是否结算完成
func (h *HandCard) Is_result() bool {
	return h.game_result != GAME_RESULT.ERROR
}

// 加倍停牌
func (h *HandCard) Double_down() {
	h.bet = 2 * h.bet
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
