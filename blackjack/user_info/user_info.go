/*
功能：用户信息
说明：
*/
package user_info

import (
	"fmt"
	"github.com/felixrobcoding/poker_oddsblackjack/define"
	"github.com/felixrobcoding/poker_oddsblackjack/define/ACTION_TYPE"
	"github.com/felixrobcoding/poker_oddsblackjack/logic"
	"github.com/felixrobcoding/poker_oddscommon/GAME_RESULT"
	"github.com/felixrobcoding/poker_oddscommon/USER_TYPE"
)

type UserInfo struct {
	t          USER_TYPE.TYPE //用户类型
	hand_index int            //当前手牌索引
	hands      []*HandCard    //手牌
	init_chip  float64        //起始筹码
	chip       float64        //筹码
	user_stat  *UserStat      //用户统计
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
	u.hands = make([]*HandCard, 0)
	u.hand_index = 0
	u.user_stat = &UserStat{}
}

// 发牌
func (u *UserInfo) Deal(cards []byte, bet_amount int) {
	hand := NewHandCard()
	hand.Append_cards(cards)
	hand.Set_bet_amount(bet_amount)

	u.hands = make([]*HandCard, 0)
	u.hand_index = 0

	u.hands = append(u.hands, hand)
}

// 手牌数
func (u *UserInfo) Hand_cnt() int {
	return len(u.hands)
}

// 下一手
func (u *UserInfo) Next_hand() error {
	if (u.hand_index + 1) > (len(u.hands) - 1) {
		return fmt.Errorf("")
	}
	u.hand_index++
	return nil
}

// 重置
func (u *UserInfo) Reset_hand_index() {
	u.hand_index = 0
}

func (u *UserInfo) Current_hand() *HandCard {
	if u.hand_index > (len(u.hands) - 1) {
		panic("")
	}
	return u.hands[u.hand_index]
}

// 分牌
func (u *UserInfo) Split_card(right_card byte) ([]byte, []byte) {
	if len(u.hands[u.hand_index].Cards()) != 2 {
		panic("")
	}
	//新的hand保留第二张
	new_hand := NewHandCard()
	new_hand.Append_card(u.hands[u.hand_index].cards[1])
	new_hand.bet_amount = u.hands[u.hand_index].bet_amount
	new_hand.action = int(ACTION_TYPE.SPLIT)

	//原来的hand保留第一张
	u.hands[u.hand_index].cards = u.hands[u.hand_index].cards[0:1]
	u.hands[u.hand_index].Append_card(right_card)
	u.hands[u.hand_index].action = int(ACTION_TYPE.SPLIT)

	u.hands = slice_insert(u.hands, u.hand_index+1, new_hand)

	return u.hands[u.hand_index].cards, u.hands[u.hand_index+1].cards
}

// slice保持顺序插入
func slice_insert(all []*HandCard, index int, data *HandCard) []*HandCard {
	new_all := make([]*HandCard, 0)
	if index == len(all) { //插入最后
		new_all = append(new_all, all...)
		new_all = append(new_all, data)
		return new_all
	}

	//插入中间
	for i := 0; i < len(all); i++ {
		if i == index {
			new_all = append(new_all, data)
		}
		new_all = append(new_all, all[i])
	}
	return new_all
}

// 更新积分
func (u *UserInfo) Update_score(score float64) {
	if u.t == USER_TYPE.BANKER { //庄家
	}

	game_result := u.hands[u.hand_index].update_score(score)
	u.chip += score
	u.user_stat.Total_bets += u.hands[u.hand_index].Get_bet_amount()

	if game_result == GAME_RESULT.WIN {
		u.user_stat.Win_hands++
	} else if game_result == GAME_RESULT.LOSE {
		u.user_stat.Lose_hands++
	} else {
		u.user_stat.Push_hands++
	}
}

// 获取起始筹码
func (u *UserInfo) Get_init_chip() float64 {
	return u.init_chip
}

// 获取筹码
func (u *UserInfo) Get_chip() float64 {
	return u.chip
}

// 利润
func (u *UserInfo) Get_profit() float64 {
	return u.chip - u.init_chip
}

// 是否全部爆点/投降
func (u *UserInfo) Is_all_bust_or_surrender() (is bool, points []int, surrender_cnt int, bust_cnt int) {
	points = make([]int, 0)
	for i := 0; i < len(u.hands); i++ {
		if u.hands[i].Is_surrender() {
			surrender_cnt++
		} else {
			tmp_points, _ := logic.Points(u.hands[i].Cards())
			points = append(points, tmp_points[0])
			if tmp_points[0] >= define.POINT_BUST {
				bust_cnt++
			}
		}
	}
	is = (surrender_cnt + bust_cnt) == len(u.hands)
	return
}

// 用户统计
func (u *UserInfo) Extract_user_stat() *UserStat {
	return u.user_stat
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
