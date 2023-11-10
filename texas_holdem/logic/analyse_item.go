/*
功能：分析项
说明：
*/
package logic

import (
	"Odds/common"
	"Odds/texas_holdem/define/CARD_TYPE"

	"fmt"
)

// 相同值
type sameValue struct {
	single_cnt int //单牌个数
	pair_cnt   int //对子个数
	three_cnt  int //三条个数
	four_cnt   int //四条个数
}

// AnalyseItem 分析项
type AnalyseItem struct {
	cards         []byte           //牌
	suit_cnts     []int            //花色统计
	value_cnts    []int            //牌值统计
	same_value    sameValue        //相同值
	Card_types    []CARD_TYPE.TYPE //牌型
	compare_cards []byte           //比较牌
}

type typeHandler struct {
	t       CARD_TYPE.TYPE
	handler HandlerAnalyseCardType
}

// NewAnalyseItem 创建
func NewAnalyseItem(cards []byte) *AnalyseItem {
	if len(cards) <= 0 {
		return nil
	}
	item := &AnalyseItem{}
	item.init()

	item.cards = cards[:]
	for _, v := range cards {
		suit := common.Suit(v)   //花色
		value := common.Value(v) //值
		item.suit_cnts[suit]++
		item.value_cnts[value-1]++
	}
	//计算相同值项
	item.cal_same_value_item()
	//计算牌型
	item.cal_card_type()
	return item
}

func (a *AnalyseItem) init() {
	a.cards = make([]byte, 0)
	a.suit_cnts = make([]int, common.SUIT_CNT)
	a.value_cnts = make([]int, common.VALUE_CNT_LOGIC)

	a.Card_types = make([]CARD_TYPE.TYPE, 0)
	a.compare_cards = make([]byte, 0)
}

// Cards 牌列表
func (a *AnalyseItem) Cards() []byte {
	return a.cards
}

// Compare_cards 比较牌列表
func (a *AnalyseItem) Compare_cards() []byte {
	return a.compare_cards
}

// Card_cnt 牌张数
func (a *AnalyseItem) Card_cnt() int {
	return len(a.cards)
}

// Cal_compare_card_type 计算比较牌型
func (a *AnalyseItem) Cal_compare_card_type() CARD_TYPE.TYPE {
	//牌型中提取比较牌型
	types := []CARD_TYPE.TYPE{
		CARD_TYPE.ONE_PAIR,
		CARD_TYPE.TWO_PAIR,
		CARD_TYPE.THREE_OF_A_KIND,
		CARD_TYPE.STRAIGHT,
		CARD_TYPE.FLUSH,
		CARD_TYPE.FULL_HOUSE,
		CARD_TYPE.FOUR_OF_A_KIND,
		CARD_TYPE.STRAIGHT_FLUSH,
		CARD_TYPE.ROYAL_STRAIGHT_FLUSH,
	}
	for _, v1 := range a.Card_types {
		for _, v2 := range types {
			if v1 == v2 {
				return v2
			}
		}
	}
	return CARD_TYPE.HIGH_CARD
}

// cal_same_value_item 计算相同值项
func (a *AnalyseItem) cal_same_value_item() {
	for _, v := range a.value_cnts {
		if v == 1 {
			a.same_value.single_cnt++
			continue
		}
		if v == 2 {
			a.same_value.pair_cnt++
			continue
		}
		if v == 3 {
			a.same_value.three_cnt++
			continue
		}
		if v == 4 {
			a.same_value.four_cnt++
			continue
		}
	}
}

// cal_card_type 计算牌型
func (a *AnalyseItem) cal_card_type() {
	type_handlers := []typeHandler{
		{CARD_TYPE.EX_FLUSH_3, is_ex_flush_3},
		{CARD_TYPE.EX_FLUSH_4, is_ex_flush_4},
		{CARD_TYPE.EX_FLUSH_5, is_ex_flush_5},
		{CARD_TYPE.EX_MORE_ONE_PAIR, is_ex_more_one_pair},
		{CARD_TYPE.ONE_PAIR, is_one_pair},
		{CARD_TYPE.TWO_PAIR, is_two_pair},
		{CARD_TYPE.THREE_OF_A_KIND, is_three_of_a_kind},
		{CARD_TYPE.STRAIGHT, is_straight},
		{CARD_TYPE.FLUSH, is_flush},
		{CARD_TYPE.FULL_HOUSE, is_full_house},
		{CARD_TYPE.FOUR_OF_A_KIND, is_four_of_a_kind},
		{CARD_TYPE.STRAIGHT_FLUSH, is_straight_flush},
		{CARD_TYPE.ROYAL_STRAIGHT_FLUSH, is_royal_straight_flush},
	}

	for i := 0; i < len(type_handlers); i++ {
		if is, compare_cards := type_handlers[i].handler(a); is { //判断牌型
			a.Card_types = append(a.Card_types, type_handlers[i].t)
			a.compare_cards = compare_cards[:]
		}
	}
	//高牌
	if len(a.Card_types) == 0 {
		a.Card_types = append(a.Card_types, CARD_TYPE.HIGH_CARD)
		a.compare_cards = sort_high_card(a)
	}
	if len(a.Card_types) == 1 && ((a.Card_types[0] == CARD_TYPE.EX_FLUSH_3) || (a.Card_types[0] == CARD_TYPE.EX_FLUSH_4)) {
		a.Card_types = append(a.Card_types, CARD_TYPE.HIGH_CARD)
	}
}

// 描述字符串
// is_include_ex:是否包含扩展牌型
func (a *AnalyseItem) String(is_include_ex bool) string {
	cards := common.Cards_2_string(a.cards)
	card_type := CARD_TYPE.String(a.Card_types) //默认是包含扩展牌型
	if !is_include_ex {
		card_type = a.Cal_compare_card_type().String()
	}
	compare_cards := common.Cards_2_string(a.compare_cards)
	return fmt.Sprintf("AnalyseItem:[原始牌:%s,牌型:%s,比较牌:%s]", cards, card_type, compare_cards)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
