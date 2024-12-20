/*
功能：策略表-转gotable输出
说明：
*/
package outputer

import (
	"fmt"
	"github.com/felixrobcoding/poker_oddsblackjack/define/HAND_TYPE"
	"github.com/felixrobcoding/poker_oddsblackjack/strategy/node"

	"github.com/liushuochen/gotable"
	"github.com/poker-x-studio/x/xdebug"
)

type PointRange struct {
	min int
	max int
}

// 图转表,console输出
func Gotable_make(chart_map map[string]*node.Node) error {

	points := []PointRange{
		{5, 21},
		{13, 21},
		{2, 11},
	}
	hand_headings := []string{HARD_HAND_HEADING, SOFT_HAND_HEADING, SPLITS_HAND_HEADING}
	hand_types := []HAND_TYPE.TYPE{HAND_TYPE.HARD, HAND_TYPE.SOFT, HAND_TYPE.SPLITS}

	for k, v := range hand_types {
		cols := make([]string, 0)
		cols = append(cols, hand_headings[k])
		cols = append(cols, X_axis_headings()...)

		table, err := gotable.Create(cols...)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		rows := make([]map[string]string, 0)
		for i := points[k].min; i <= points[k].max; i++ {
			row := extract_row(hand_headings[k], chart_map, v, i)
			if row != nil {
				rows = append(rows, row)
			}
		}

		for _, row := range rows {
			err := table.AddRow(row)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
		}

		//table.String()
		if xdebug.Is_debug() {
			//fmt.Println(table.String())
		}
	}
	return nil
}

// 提取行
func extract_row(heading string, chart_map map[string]*node.Node, hand_type HAND_TYPE.TYPE, player_point int) map[string]string {
	nodes := make([]*node.Node, 0)

	x_axis_headings := X_axis_headings()
	for _, v := range x_axis_headings {
		key := node.Make_key(hand_type, player_point, v)
		node, ok := chart_map[key]
		if ok {
			nodes = append(nodes, node)
		}
	}

	y_axis_headings_splits := Y_axis_headings_splits()

	row := make(map[string]string, 0)
	row[heading] = fmt.Sprintf("%d", player_point)
	if hand_type == HAND_TYPE.SPLITS {
		row[heading] = y_axis_headings_splits[player_point-2]
	}
	for _, node := range nodes {
		row[node.Dealer_value()] = node.Action.String_short()
	}
	return row
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
