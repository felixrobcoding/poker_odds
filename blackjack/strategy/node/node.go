/*
功能：策略助手-节点
说明：
*/
package node

import (
	"Odds/blackjack/define/ACTION_TYPE"
	"Odds/blackjack/define/HAND_TYPE"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	hand_type    HAND_TYPE.TYPE   //
	player_point int              //闲家牌点数
	dealer_value string           //庄家牌值
	Action       ACTION_TYPE.TYPE //
}

func NewNode(hand_type HAND_TYPE.TYPE, player_point int, dealer_value string, action ACTION_TYPE.TYPE) *Node {
	node := &Node{hand_type, player_point, dealer_value, action}
	return node
}

func (n *Node) Make_key() string {
	return Make_key(n.hand_type, n.player_point, n.dealer_value)
}

func (n *Node) Dealer_value() string {
	return n.dealer_value
}

func Make_key(hand_type HAND_TYPE.TYPE, player_point int, dealer_value string) string {
	return fmt.Sprintf("%d_%d_%s", int(hand_type), player_point, dealer_value)
}

// 分离
func Split_key(key string) (hand_type HAND_TYPE.TYPE, player_point int, dealer_value string) {
	split_txts := strings.Split(key, "_")
	if len(split_txts) != 3 {
		return HAND_TYPE.ERROR, 0, ""
	}
	tmp, err := strconv.Atoi(split_txts[0])
	if err != nil {
		return HAND_TYPE.ERROR, 0, ""
	}
	hand_type = HAND_TYPE.TYPE(tmp)

	tmp, err = strconv.Atoi(split_txts[1])
	if err != nil {
		return HAND_TYPE.ERROR, 0, ""
	}
	player_point = tmp
	dealer_value = split_txts[2]
	return
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
