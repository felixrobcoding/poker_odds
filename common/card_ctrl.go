/*
功能：牌控件
说明：
*/
package common

import "fmt"

// CardCtrl 牌控件
type CardCtrl struct {
	Cards    []byte //
	full_cnt int    //张数
}

// NewCardCtrlWithCard 创建
func NewCardCtrlWithCard(card byte, full_cnt int) *CardCtrl {
	ctrl := &CardCtrl{}
	ctrl.init()

	ctrl.Cards = append(ctrl.Cards, card)
	ctrl.full_cnt = full_cnt
	return ctrl
}

// NewCardCtrlWithCards 创建
func NewCardCtrlWithCards(cards []byte, full_cnt int) *CardCtrl {
	ctrl := &CardCtrl{}
	ctrl.init()

	if cards != nil {
		ctrl.Cards = append(ctrl.Cards, cards...)
	}
	ctrl.full_cnt = full_cnt
	return ctrl
}

// init 初始化
func (c *CardCtrl) init() {
	c.Cards = make([]byte, 0)
	c.full_cnt = 0
}

// Is_full 是否充满
func (c *CardCtrl) Is_full() bool {
	return c.full_cnt == len(c.Cards) && c.full_cnt > 0
}

// Append 追加
func (c *CardCtrl) Append_card(card byte) {
	if c.Cards == nil {
		return
	}
	c.Cards = append(c.Cards, card)
}

func (c *CardCtrl) Append_cards(cards []byte) {
	if cards == nil {
		return
	}
	if c.Cards == nil {
		return
	}
	c.Cards = append(c.Cards, cards...)
}

//转成字符串
func (c *CardCtrl) String() string {
	str := "CardCtrl:[full_cnt:"
	str += fmt.Sprintf("%d,Cards:", c.full_cnt)
	for k, v := range c.Cards {
		str += fmt.Sprintf("%d", v)
		if k < (len(c.Cards) - 1) {
			str += ","
		}
	}
	str += "]"
	return str
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
