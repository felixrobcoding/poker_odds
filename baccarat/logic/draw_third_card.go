/*
功能：补牌
说明：

1 https://en.wikipedia.org/wiki/Baccarat
*/
package logic

import "github.com/poker-x-studio/x/xutils"

// 补闲家的牌
//If the player has an initial total of 5 or less, they draw a third card. If the player has an initial total of 6 or 7, they stand.
func Draw_card_for_player(player_cards []byte, dealer_cards []byte) bool {
	len := len(player_cards)
	if len != 2 {
		return false
	}
	point := Points(player_cards)
	if point <= 5 { //必须补牌
		return true
	}
	return false
}

//补庄家的牌
func Draw_card_for_dealer(player_cards []byte, dealer_cards []byte) bool {
	dealer_len := len(dealer_cards)
	if dealer_len != 2 {
		return false
	}
	point := Points(dealer_cards)
	if point <= 2 { //必须补牌
		return true
	} else if point >= 3 && point <= 6 {
		//闲家是否已经补牌
		is_player_draw := len(player_cards) > 2

		player_third_point := 0
		if is_player_draw {
			player_third_point = Point(player_cards[2])
		}

		if point == 3 { //若闲家补得第三张牌是8，不得补牌
			points := []int{8}
			if (is_player_draw) && (xutils.Is_contains(points, player_third_point)) {
				return false
			}
			return true
		} else if point == 4 { //若闲家补得第三张牌是0,1,8,9，不得补牌
			points := []int{0, 1, 8, 9}
			if (is_player_draw) && (xutils.Is_contains(points, player_third_point)) {
				return false
			}
			return true
		} else if point == 5 { //若闲家补得第三张牌是0,1,2,3,8,9，不得补牌
			points := []int{0, 1, 2, 3, 8, 9}
			if (is_player_draw) && (xutils.Is_contains(points, player_third_point)) {
				return false
			}
			return true
		} else if point == 6 { //若闲家补得第三张牌是6,7，必须补牌
			points := []int{6, 7}
			if (is_player_draw) && (xutils.Is_contains(points, player_third_point)) {
				return true
			}
			return false
		}
	} else if point >= 7 { //不得补牌
		return false
	}

	return false
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
