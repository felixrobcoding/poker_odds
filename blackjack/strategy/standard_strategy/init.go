/*
功能：标准策略表
说明：
*/
package standard_strategy

import (
	"github.com/felixrobcoding/poker_oddsblackjack/define/ACTION_TYPE"
	"github.com/felixrobcoding/poker_oddsblackjack/define/HAND_TYPE"
	"github.com/felixrobcoding/poker_oddsblackjack/strategy/node"
	"github.com/felixrobcoding/poker_oddsblackjack/strategy/outputer"
)

type Chart struct {
	chart_map map[string]*node.Node //策略表map
}

var instance Chart

// 对外接口
func Get_strategy_map() map[string]*node.Node {
	if instance.get_map() == nil {
		instance.init()
	}
	return instance.get_map()
}

// 获取map
func (c *Chart) get_map() map[string]*node.Node {
	return c.chart_map
}

// 初始化
func (c *Chart) init() {
	if c.chart_map == nil {
		c.chart_map = make(map[string]*node.Node, 0)
	}
	c.init_hard_hand()
	c.init_soft_hand()
	c.init_splits_hand()
}

// 初始化-hard
func (c *Chart) init_hard_hand() {
	x_axis_headings := outputer.X_axis_headings()
	x_axis_cell_cnt := len(x_axis_headings)

	//hit
	for i := 5; i <= 21; i++ {
		for j := 0; j < x_axis_cell_cnt; j++ {
			nd := node.NewNode(HAND_TYPE.HARD, i, x_axis_headings[j], ACTION_TYPE.HIT)
			key := nd.Make_key()
			//fmt.Println(key)
			c.chart_map[key] = nd
		}
	}

	//stand
	for i := 12; i <= 16; i++ {
		for j := 0; j < 5; j++ {
			if i == 12 && j <= 1 {
				continue
			}
			if j > 4 {
				continue
			}
			nd := node.NewNode(HAND_TYPE.HARD, i, x_axis_headings[j], ACTION_TYPE.STAND)
			key := nd.Make_key()
			c.chart_map[key] = nd
		}
	}

	for i := 17; i <= 21; i++ {
		for j := 0; j < x_axis_cell_cnt; j++ {
			nd := node.NewNode(HAND_TYPE.HARD, i, x_axis_headings[j], ACTION_TYPE.STAND)
			key := nd.Make_key()
			c.chart_map[key] = nd
		}
	}

	//double down
	for i := 9; i <= 11; i++ {
		for j := 0; j < x_axis_cell_cnt; j++ {
			if i == 9 && (j == 0 || j > 4) {
				continue
			}
			if i == 10 && (j > 7) {
				continue
			}
			if i == 11 && (j > 8) {
				continue
			}
			nd := node.NewNode(HAND_TYPE.HARD, i, x_axis_headings[j], ACTION_TYPE.DOUBLE_DOWN)
			key := nd.Make_key()
			c.chart_map[key] = nd
		}
	}

	//surrender
	for i := 14; i <= 16; i++ {
		for j := 0; j < x_axis_cell_cnt; j++ {
			if (j != 7) && (j != 8) {
				continue
			}
			if (i == 14 || i == 15) && (j == 7) {
				continue
			}
			nd := node.NewNode(HAND_TYPE.HARD, i, x_axis_headings[j], ACTION_TYPE.SURRENDER)
			key := nd.Make_key()
			c.chart_map[key] = nd
		}
	}
}

// 初始化-soft
func (c *Chart) init_soft_hand() {
	x_axis_headings := outputer.X_axis_headings()
	x_axis_cell_cnt := len(x_axis_headings)

	//hit
	for i := 13; i <= 21; i++ {
		for j := 0; j < x_axis_cell_cnt; j++ {
			nd := node.NewNode(HAND_TYPE.SOFT, i, x_axis_headings[j], ACTION_TYPE.HIT)
			key := nd.Make_key()
			//fmt.Println(key)
			c.chart_map[key] = nd
		}
	}

	//stand
	for i := 18; i <= 21; i++ {
		for j := 0; j < x_axis_cell_cnt; j++ {
			if i == 18 && j > 6 {
				continue
			}
			nd := node.NewNode(HAND_TYPE.SOFT, i, x_axis_headings[j], ACTION_TYPE.STAND)
			key := nd.Make_key()
			//fmt.Println(key)
			c.chart_map[key] = nd
		}
	}

	//double down
	for i := 13; i <= 18; i++ {
		for j := 0; j < x_axis_cell_cnt; j++ {
			if (i == 13 || i == 14) && (j <= 2 || j >= 5) {
				continue
			}
			if (i == 15 || i == 16) && (j <= 1 || j >= 5) {
				continue
			}
			if (i == 17 || i == 18) && (j <= 0 || j >= 5) {
				continue
			}
			nd := node.NewNode(HAND_TYPE.SOFT, i, x_axis_headings[j], ACTION_TYPE.DOUBLE_DOWN)
			key := nd.Make_key()
			//fmt.Println(key)
			c.chart_map[key] = nd
		}
	}
}

// 初始化-splits
func (c *Chart) init_splits_hand() {
	x_axis_headings := outputer.X_axis_headings()
	x_axis_cell_cnt := len(x_axis_headings)

	//hit
	for i := 2; i <= 11; i++ {
		for j := 0; j < x_axis_cell_cnt; j++ {
			nd := node.NewNode(HAND_TYPE.SPLITS, i, x_axis_headings[j], ACTION_TYPE.HIT)
			key := nd.Make_key()
			//fmt.Println(key)
			c.chart_map[key] = nd
		}
	}

	//split
	for i := 2; i <= 11; i++ {
		for j := 0; j < x_axis_cell_cnt; j++ {
			if (i == 2 || i == 3) && (j >= 6) {
				continue
			}
			if (i == 4) && (j <= 2 || j >= 5) {
				continue
			}
			if i == 5 {
				continue
			}
			if (i == 6) && (j >= 5) {
				continue
			}
			if (i == 7) && (j >= 6) {
				continue
			}

			nd := node.NewNode(HAND_TYPE.SPLITS, i, x_axis_headings[j], ACTION_TYPE.SPLIT)
			key := nd.Make_key()
			//fmt.Println(key)
			c.chart_map[key] = nd
		}
	}

	//double
	for i := 5; i <= 5; i++ {
		for j := 0; j < x_axis_cell_cnt; j++ {
			if j >= 8 {
				continue
			}

			nd := node.NewNode(HAND_TYPE.SPLITS, i, x_axis_headings[j], ACTION_TYPE.DOUBLE_DOWN)
			key := nd.Make_key()
			//fmt.Println(key)
			c.chart_map[key] = nd
		}
	}

	//stand
	for i := 9; i <= 10; i++ {
		for j := 0; j < x_axis_cell_cnt; j++ {
			if (i == 9) && (j <= 4 || j >= 6 && j <= 7) {
				continue
			}

			nd := node.NewNode(HAND_TYPE.SPLITS, i, x_axis_headings[j], ACTION_TYPE.STAND)
			key := nd.Make_key()
			//fmt.Println(key)
			c.chart_map[key] = nd
		}
	}

	//surrender
	for i := 7; i <= 8; i++ {
		for j := 0; j < x_axis_cell_cnt; j++ {
			if j != 8 {
				continue
			}
			nd := node.NewNode(HAND_TYPE.SPLITS, i, x_axis_headings[j], ACTION_TYPE.SURRENDER)
			key := nd.Make_key()
			c.chart_map[key] = nd
		}
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
