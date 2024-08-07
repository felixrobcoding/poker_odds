/*
功能：比较牌型
说明：
*/
package logic

import (
	"Odds/common"
	"Odds/common/COMPARE_TYPE"
	"Odds/texas_holdem/define/CARD_TYPE"
)

// Compare 比较牌型
func Compare(source *AnalyseItem, target *AnalyseItem) COMPARE_TYPE.TYPE {
	card_type_s := source.Cal_compare_card_type()
	card_type_t := target.Cal_compare_card_type()

	if card_type_t > card_type_s {
		return COMPARE_TYPE.BIGGER
	} else if card_type_t < card_type_s {
		return COMPARE_TYPE.SMALLER
	} else { //牌型相等
		if card_type_t == CARD_TYPE.ONE_PAIR {
			return compare_one_pair(source, target)
		} else if card_type_t == CARD_TYPE.TWO_PAIR {
			return compare_two_pair(source, target)
		} else if card_type_t == CARD_TYPE.THREE_OF_A_KIND {
			return compare_three_of_a_kind(source, target)
		} else if card_type_t == CARD_TYPE.STRAIGHT {
			return compare_straight(source, target)
		} else if card_type_t == CARD_TYPE.FLUSH {
			return compare_flush(source, target)
		} else if card_type_t == CARD_TYPE.FULL_HOUSE {
			return compare_full_house(source, target)
		} else if card_type_t == CARD_TYPE.FOUR_OF_A_KIND {
			return compare_four_of_a_kind(source, target)
		} else if card_type_t == CARD_TYPE.STRAIGHT_FLUSH {
			return compare_straight_flush(source, target)
		} else if card_type_t == CARD_TYPE.ROYAL_STRAIGHT_FLUSH {
			return compare_royal_straight_flush(source, target)
		}
		return COMPARE_TYPE.EQUAL
	}
}

func compare_one_pair(source *AnalyseItem, target *AnalyseItem) COMPARE_TYPE.TYPE {
	source_cards := source.Compare_cards()
	target_cards := target.Compare_cards()

	is_equal := false
	for i := 0; i < len(target_cards); i++ {
		ct := compare_one_card(source_cards[i], target_cards[i])
		if ct == COMPARE_TYPE.EQUAL {
			is_equal = true
		} else {
			return ct
		}
	}
	if is_equal {
		return COMPARE_TYPE.EQUAL
	}
	return COMPARE_TYPE.SMALLER
}

func compare_two_pair(source *AnalyseItem, target *AnalyseItem) COMPARE_TYPE.TYPE {
	source_cards := source.Compare_cards()
	target_cards := target.Compare_cards()

	is_equal := false
	for i := 0; i < len(target_cards); i++ {
		ct := compare_one_card(source_cards[i], target_cards[i])
		if ct == COMPARE_TYPE.EQUAL {
			is_equal = true
		} else {
			return ct
		}
	}
	if is_equal {
		return COMPARE_TYPE.EQUAL
	}
	return COMPARE_TYPE.SMALLER
}

func compare_three_of_a_kind(source *AnalyseItem, target *AnalyseItem) COMPARE_TYPE.TYPE {
	return compare_one_card(source.Compare_cards()[0], target.Compare_cards()[0])
}

func compare_straight(source *AnalyseItem, target *AnalyseItem) COMPARE_TYPE.TYPE {
	source_cards := source.Compare_cards()
	target_cards := target.Compare_cards()

	is_equal := false
	for i := 0; i < len(target_cards); i++ {
		ct := compare_one_card(source_cards[i], target_cards[i])
		if ct == COMPARE_TYPE.EQUAL {
			is_equal = true
		} else {
			return ct
		}
	}
	if is_equal {
		return COMPARE_TYPE.EQUAL
	}
	return COMPARE_TYPE.SMALLER
}

func compare_flush(source *AnalyseItem, target *AnalyseItem) COMPARE_TYPE.TYPE {
	source_cards := source.Compare_cards()
	target_cards := target.Compare_cards()

	is_equal := false
	for i := 0; i < len(target_cards); i++ {
		ct := compare_one_card(source_cards[i], target_cards[i])
		if ct == COMPARE_TYPE.EQUAL {
			is_equal = true
		} else {
			return ct
		}
	}
	if is_equal {
		return COMPARE_TYPE.EQUAL
	}
	return COMPARE_TYPE.SMALLER
}

func compare_full_house(source *AnalyseItem, target *AnalyseItem) COMPARE_TYPE.TYPE {
	return compare_one_card(source.Compare_cards()[0], target.Compare_cards()[0])
}

func compare_four_of_a_kind(source *AnalyseItem, target *AnalyseItem) COMPARE_TYPE.TYPE {
	return compare_one_card(source.Compare_cards()[0], target.Compare_cards()[0])
}

func compare_straight_flush(source *AnalyseItem, target *AnalyseItem) COMPARE_TYPE.TYPE {
	source_cards := source.Compare_cards()
	target_cards := target.Compare_cards()

	is_equal := false
	for i := 0; i < len(target_cards); i++ {
		ct := compare_one_card(source_cards[i], target_cards[i])
		if ct == COMPARE_TYPE.EQUAL {
			is_equal = true
		} else {
			return ct
		}
	}
	if is_equal {
		return COMPARE_TYPE.EQUAL
	}
	return COMPARE_TYPE.SMALLER
}

func compare_royal_straight_flush(source *AnalyseItem, target *AnalyseItem) COMPARE_TYPE.TYPE {
	return COMPARE_TYPE.EQUAL
}

// compare_one_card 比较单张牌
// _2_logic_card:是否转换为逻辑牌
func compare_one_card(source byte, target byte) COMPARE_TYPE.TYPE {
	s_value := common.Value(source)
	t_value := common.Value(target)

	if t_value > s_value {
		return COMPARE_TYPE.BIGGER
	} else if t_value < s_value {
		return COMPARE_TYPE.SMALLER
	} else {
		return COMPARE_TYPE.EQUAL
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
