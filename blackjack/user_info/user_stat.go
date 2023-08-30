/*
功能：用户统计
说明：
*/
package user_info

import "fmt"

type UserStat struct {
	Total_bets int //总投注
	Lose_hands int //输的手数
	Push_hands int //和的手数
	Win_hands  int //赢的手数
}

func (u UserStat) String() string {
	str := fmt.Sprintf("[Total_bets:%d,Lose_hands:%d,Push_hands:%d,Win_hands:%d]", u.Total_bets, u.Lose_hands, u.Push_hands, u.Win_hands)
	return str
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
