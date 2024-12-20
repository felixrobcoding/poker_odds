/*
功能：策略表-转svg输出
说明：
*/
package outputer

import (
	"bytes"
	"fmt"
	"github.com/felixrobcoding/poker_oddsblackjack/define/ACTION_TYPE"
	"github.com/felixrobcoding/poker_oddsblackjack/define/HAND_TYPE"
	"github.com/felixrobcoding/poker_oddsblackjack/strategy/node"

	svg "github.com/ajstarks/svgo"
	"github.com/poker-x-studio/x/xmath"
)

const (
	TXT_HEADING_BASIC_STRATEGY = "Blackjack Basic Strategy"
)

type SvgMakerStrategy struct {
	SvgMaker
}

var strategy_svg SvgMakerStrategy

// 创建
func Strategy_svg_make(strategy_map map[string]*node.Node) string {
	//设置参数
	strategy_svg.y_axis_cnt = 32
	strategy_svg.header = TXT_HEADING_BASIC_STRATEGY

	return strategy_svg.make(strategy_map)
}

// 生成svg图
func (s *SvgMakerStrategy) make(strategy_map map[string]*node.Node) string {
	svg_writer := new(bytes.Buffer) //写入缓冲
	canvas := svg.New(svg_writer)

	width := s.canvas_width()
	height := s.canvas_height()
	canvas.Startview(int(float64(width)*SCALE), int(float64(height)*SCALE), 0, 0, width, height)

	//defs
	s.make_defs(canvas)
	//背景
	s.make_bg(canvas)
	//header
	s.make_header(canvas)
	//body
	s.make_body(canvas, strategy_map)
	//footer
	s.make_footer(canvas)

	canvas.End()

	//测试
	//fmt.Println(svg_writer)
	return svg_writer.String()
}

func (s *SvgMakerStrategy) make_body(canvas *svg.SVG, strategy_map map[string]*node.Node) {
	x_axis_headings := X_axis_headings()
	x_axis_cell_cnt := len(x_axis_headings)

	//起点
	y_base := H_HEADER
	origin_rect_hard := *xmath.NewRect(MARGIN_LEFT+WIDTH_Y_AXIS_HEADING, y_base, GRID_WIDTH, GRID_HEIGHT)

	//hard
	s.make_x_axis_headings(canvas, y_base-5)
	for i := 8; i <= 17; i++ {
		rect := origin_rect_hard
		rect.Y_move((i - 8) * GRID_HEIGHT)
		for j := 0; j < x_axis_cell_cnt; j++ {
			key := node.Make_key(HAND_TYPE.HARD, i, x_axis_headings[j])
			v, ok := strategy_map[key]
			if ok {
				if j > 0 {
					rect.X_move(GRID_WIDTH)
				}
				s.make_grid(canvas, &rect, v.Action)
			}
		}
	}
	s.make_y_axis_headings_hard(canvas, y_base-5, 8, 17, false)

	//soft
	y_base = H_HEADER + 12*GRID_HEIGHT
	origin_rect_soft := *xmath.NewRect(MARGIN_LEFT+WIDTH_Y_AXIS_HEADING, y_base, GRID_WIDTH, GRID_HEIGHT)

	s.make_x_axis_headings(canvas, y_base-5)
	for i := 13; i <= 19; i++ {
		rect := origin_rect_soft
		rect.Y_move((i - 13) * GRID_HEIGHT)
		for j := 0; j < x_axis_cell_cnt; j++ {
			key := node.Make_key(HAND_TYPE.SOFT, i, x_axis_headings[j])
			v, ok := strategy_map[key]
			if ok {
				if j > 0 {
					rect.X_move(GRID_WIDTH)
				}
				s.make_grid(canvas, &rect, v.Action)
			}
		}
	}
	s.make_y_axis_headings_soft(canvas, y_base-5, 13, 19, false)

	//splits
	y_base = H_HEADER + 21*GRID_HEIGHT
	origin_rect_splits := *xmath.NewRect(MARGIN_LEFT+WIDTH_Y_AXIS_HEADING, y_base, GRID_WIDTH, GRID_HEIGHT)

	s.make_x_axis_headings(canvas, y_base-5)
	for i := 2; i <= 11; i++ {
		rect := origin_rect_splits
		rect.Y_move((i - 2) * GRID_HEIGHT)
		for j := 0; j < x_axis_cell_cnt; j++ {
			key := node.Make_key(HAND_TYPE.SPLITS, i, x_axis_headings[j])
			v, ok := strategy_map[key]
			if ok {
				if j > 0 {
					rect.X_move(GRID_WIDTH)
				}
				s.make_grid(canvas, &rect, v.Action)
			}
		}
	}
	s.make_y_axis_headings_splits(canvas, y_base-5)
}

// 一格
func (s *SvgMakerStrategy) make_grid(canvas *svg.SVG, rect *xmath.Rect, action ACTION_TYPE.TYPE) {
	txt := fmt.Sprintf("translate(%d,%d)", rect.X_left, rect.Y_top)
	canvas.Gtransform(txt)
	defer canvas.Gend()

	//背景
	actions := []ACTION_TYPE.TYPE{
		ACTION_TYPE.HIT, ACTION_TYPE.DOUBLE_DOWN, ACTION_TYPE.STAND, ACTION_TYPE.SPLIT, ACTION_TYPE.SURRENDER,
	}
	style_bgs := []string{
		STYLE_BG_HIT, STYLE_BG_DOUBLE, STYLE_BG_STAND, STYLE_BG_SPLIT, STYLE_BG_SURRENDER,
	}

	style := STYLE_BG_UNKOWN
	for k, v := range actions {
		if v == action {
			style = style_bgs[k]
		}
	}
	canvas.Rect(0, 0, rect.Witdh(), rect.Height(), style)

	//标题
	canvas.Text(rect.Witdh()/2-10, rect.Height()/2+10, action.String_short(), STYLE_TXT_ACTION_18)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
