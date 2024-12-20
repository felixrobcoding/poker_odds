/*
功能：用户信息
说明：
*/
package user_info

import (
	"github.com/felixrobcoding/poker_oddsbaccarat/define/BET_AREA"
	"github.com/felixrobcoding/poker_oddscommon/GAME_RESULT"
	"github.com/felixrobcoding/poker_oddscommon/USER_TYPE"
)

type UserInfo struct {
	t         USER_TYPE.TYPE //用户类型
	hand      *HandCard      //手牌
	init_chip float64        //起始筹码
	chip      float64        //筹码
	user_stat *UserStat      //用户统计
}

func NewUserInfo(user_type USER_TYPE.TYPE, init_chip float64) *UserInfo {
	user := &UserInfo{
		t:         user_type,
		init_chip: init_chip,
		chip:      init_chip,
	}
	user.init()
	return user
}

func (u *UserInfo) init() {
	u.hand = NewHandCard()
	u.user_stat = &UserStat{}
}

// Deal 发牌
func (u *UserInfo) Deal(cards []byte, bet_area BET_AREA.TYPE, bet_amount int) {
	u.hand = NewHandCard()
	u.hand.Append_cards(cards)
	u.hand.Set_bet(bet_area, bet_amount)
}

func (u *UserInfo) Current_hand() *HandCard {
	return u.hand
}

// Update_score 更新积分
func (u *UserInfo) Update_score(score float64, win_bet_areas []BET_AREA.TYPE) {
	if u.t == USER_TYPE.BANKER { //庄家
	}

	game_result := u.hand.update_score(score, win_bet_areas)
	u.chip += score

	_, bet_amount, _ := u.hand.Get_bet()
	u.user_stat.Total_bets += bet_amount

	if game_result == GAME_RESULT.WIN {
		u.user_stat.Win_hands++
	} else if game_result == GAME_RESULT.LOSE {
		u.user_stat.Lose_hands++
	} else {
		u.user_stat.Tie_hands++
	}
}

// Get_init_chip 获取起始筹码
func (u *UserInfo) Get_init_chip() float64 {
	return u.init_chip
}

// Get_chip 获取筹码
func (u *UserInfo) Get_chip() float64 {
	return u.chip
}

// Get_profit 利润
func (u *UserInfo) Get_profit() float64 {
	return u.chip - u.init_chip
}

// Extract_user_stat 提取用户统计
func (u *UserInfo) Extract_user_stat() *UserStat {
	return u.user_stat
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
