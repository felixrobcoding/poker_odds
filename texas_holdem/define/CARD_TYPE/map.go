/*
功能：牌型map
说明：
*/
package CARD_TYPE

import "sort"

type Item struct {
	t       TYPE
	txt_eng string
	txt_chs string
}

var (
	type_map = make(map[TYPE]Item, 0)
)

func init() {
	items := []Item{
		{EX_FLUSH_3, "Flush 3", "三张同花"},
		{EX_FLUSH_4, "Flush 4", "四张同花"},
		{EX_FLUSH_5, "Flush 5", "五张同花"},
		{EX_MORE_ONE_PAIR, "More one pair", "至少一对"},
		{HIGH_CARD, "High card", "高牌"},
		{ONE_PAIR, "One pair", "一对"},
		{TWO_PAIR, "Two pair", "两对"},
		{THREE_OF_A_KIND, "Three of a Kind", "三条"},
		{STRAIGHT, "Straight", "顺子"},
		{FLUSH, "Flush", "同花"},
		{FULL_HOUSE, "Full house", "葫芦"},
		{FOUR_OF_A_KIND, "Four of a Kind", "四条"},
		{STRAIGHT_FLUSH, "Straight Flush", "同花顺"},
		{ROYAL_STRAIGHT_FLUSH, "Royal Straight Flush", "皇家同花顺"},
	}
	for k, v := range items {
		type_map[v.t] = items[k]
	}
}

// 牌型
func Card_types() []TYPE {
	types := make([]TYPE, 0)
	for _, v := range type_map {
		types = append(types, v.t)
	}

	//排序,从小到大
	sort.SliceStable(types, func(i, j int) bool {
		return types[i] < types[j]
	})

	return types
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
