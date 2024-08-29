/*
功能：手牌
说明：
*/
package user_info

import (
	"Odds/baccarat/define/BET_AREA"
	"Odds/baccarat/define/CARD_TYPE"
	"Odds/baccarat/logic"
	"Odds/common/GAME_RESULT"
)

type HandCard struct {
	cards         []byte
	game_result   GAME_RESULT.TYPE //手牌结果
	bet_area      BET_AREA.TYPE    //下注区域
	bet_amount    int              //下注额
	win_bet_areas []BET_AREA.TYPE  //获胜下注区域
	score         float64          //得分
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

// Append_card 追加
func (h *HandCard) Append_card(card byte) {
	if h.cards == nil {
		panic("")
	}
	h.cards = append(h.cards, card)
}

func (h *HandCard) Append_cards(cards []byte) {
	if h.cards == nil {
		panic("")
	}
	h.cards = append(h.cards, cards...)
}

// Card_type 获取牌型
func (h *HandCard) Card_type() CARD_TYPE.TYPE {
	return logic.Card_type(h.cards)
}

// Set_bet 设置下注额
func (h *HandCard) Set_bet(bet_area BET_AREA.TYPE, bet_amount int) {
	h.bet_area = bet_area
	h.bet_amount = bet_amount
}

// Get_bet 获取下注
func (h *HandCard) Get_bet() (bet_area BET_AREA.TYPE, bet_amount int, win_bet_areas []BET_AREA.TYPE) {
	bet_area = h.bet_area
	bet_amount = h.bet_amount
	win_bet_areas = h.win_bet_areas
	return
}

// Get_score 获取积分
func (h *HandCard) Get_score() float64 {
	return h.score
}

// update_score 更新积分
func (h *HandCard) update_score(score float64, win_bet_areas []BET_AREA.TYPE) (game_result GAME_RESULT.TYPE) {
	game_result = GAME_RESULT.PUSH
	if score > 0 {
		game_result = GAME_RESULT.WIN
	} else if score < 0 {
		game_result = GAME_RESULT.LOSE
	}

	h.game_result = game_result
	h.score = score
	h.win_bet_areas = win_bet_areas
	return
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
