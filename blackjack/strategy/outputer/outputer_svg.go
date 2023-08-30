/*
功能：转svg输出-基类
说明：
*/
package outputer

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
	"github.com/poker-x-studio/x/xmath"
	"github.com/poker-x-studio/x/xutils"
)

const (
	TXT_FOOTER = xutils.POKER_X_STUDIO //footer 文字

	GRID_WIDTH           = 36          //
	GRID_HEIGHT          = 36          //
	X_AXIS_GRID_CNT      = 10          //X轴方向个数
	WIDTH_Y_AXIS_HEADING = 40          //标题
	SCALE                = 1.5         //缩放比
	MARGIN_LEFT          = 20          //左边距
	MARGIN_RIGHT         = MARGIN_LEFT //右边距
	H_HEADER             = 120         //header高
	H_FOOTER             = 40          //footer高

	STYLE_HEADER             = "font-size:22;fill:#000"
	STYLE_FOOTER             = "font-size:22;fill:#000"
	STYLE_Y_AXIS_HEADING     = "font-size:18;fill:#000"
	STYLE_BG_HIT             = "fill:#FF0000;stroke:rgb(00,00,00);stroke-width:1" //背景
	STYLE_BG_DOUBLE          = "fill:#0000FF;stroke:rgb(00,00,00);stroke-width:1"
	STYLE_BG_STAND           = "fill:#FFFF00;stroke:rgb(00,00,00);stroke-width:1"
	STYLE_BG_SPLIT           = "fill:#00FF00;stroke:rgb(00,00,00);stroke-width:1"
	STYLE_BG_SURRENDER       = "fill:#FFFFFF;stroke:rgb(00,00,00);stroke-width:1"
	STYLE_BG_UNKOWN          = "fill:#DCDCDC;stroke:rgb(00,00,00);stroke-width:1"
	STYLE_TXT_Y_AXIS_HEADING = "font-size:18;fill:#000000" //文字
	STYLE_TXT_ACTION_18      = "font-size:18;fill:#000000" //文字
	STYLE_TXT_ACTION_12      = "font-size:12;fill:#000000" //文字
)

type SvgMaker struct {
	y_axis_cnt int    //y轴个数
	header     string //标题
}

// 宽度
func (s *SvgMaker) canvas_width() int {
	width := (X_AXIS_GRID_CNT*GRID_WIDTH + WIDTH_Y_AXIS_HEADING + MARGIN_LEFT + MARGIN_RIGHT)
	return width
}

// 高度
func (s *SvgMaker) canvas_height() int {
	height := H_HEADER + s.y_axis_cnt*GRID_HEIGHT + H_FOOTER //高度
	return height
}

// defs
func (s *SvgMaker) make_defs(canvas *svg.SVG) {
	canvas.Def()
	defer canvas.DefEnd()
}

// 整个svg图的背景
func (s *SvgMaker) make_bg(canvas *svg.SVG) {
	canvas.Gtransform("translate(0,0)")
	defer canvas.Gend()

	//背景填充
	canvas.Rect(0, 0, s.canvas_width(), s.canvas_height(), "fill:#E6E6FA")
}

// 头
func (s *SvgMaker) make_header(canvas *svg.SVG) {
	canvas.Text(s.canvas_width()/2-150, H_HEADER/2, s.header, STYLE_HEADER)
}

func (s *SvgMaker) make_body(canvas *svg.SVG) {
}

// 一格
func (s *SvgMaker) make_grid(canvas *svg.SVG) {
}

// x轴标题
func (s *SvgMaker) make_x_axis_headings(canvas *svg.SVG, y_base int) {
	const x_base int = MARGIN_LEFT + WIDTH_Y_AXIS_HEADING + GRID_HEIGHT/2 - 10

	canvas.Gtransform("translate(0,0)")
	defer canvas.Gend()

	x_axis_heaings := X_axis_headings()
	for k, v := range x_axis_heaings {
		canvas.Text(x_base+k*GRID_WIDTH, y_base, v, STYLE_Y_AXIS_HEADING)
	}
}

// y轴标题-hard
func (s *SvgMaker) make_y_axis_headings_hard(canvas *svg.SVG, y int, start_point int, end_point int, is_all bool) {
	//起点
	origin_rect := *xmath.NewRect(MARGIN_LEFT, y, GRID_WIDTH, GRID_HEIGHT)

	rect := origin_rect
	rect.Y_move(-1 * GRID_HEIGHT)
	s.make_y_axis_heading(canvas, &rect, HARD_HAND_HEADING)

	for i := start_point; i <= end_point; i++ {
		rect := origin_rect
		rect.Y_move((i - start_point) * GRID_HEIGHT)
		if is_all {
			s.make_y_axis_heading(canvas, &rect, Y_axis_headings_hard_all()[i-start_point])
		} else {
			s.make_y_axis_heading(canvas, &rect, Y_axis_headings_hard()[i-start_point])
		}
	}
}

// y轴标题-soft
func (s *SvgMaker) make_y_axis_headings_soft(canvas *svg.SVG, y int, start_point int, end_point int, is_all bool) {
	//起点
	origin_rect := *xmath.NewRect(MARGIN_LEFT, y, GRID_WIDTH, GRID_HEIGHT)

	rect := origin_rect
	rect.Y_move(-1 * GRID_HEIGHT)
	s.make_y_axis_heading(canvas, &rect, SOFT_HAND_HEADING)

	for i := start_point; i <= end_point; i++ {
		rect := origin_rect
		rect.Y_move((i - start_point) * GRID_HEIGHT)
		if is_all {
			s.make_y_axis_heading(canvas, &rect, Y_axis_headings_soft_all()[i-start_point])
		} else {
			s.make_y_axis_heading(canvas, &rect, Y_axis_headings_soft()[i-start_point])
		}
	}
}

// y轴标题-splits
func (s *SvgMaker) make_y_axis_headings_splits(canvas *svg.SVG, y int) {
	//起点
	origin_rect := *xmath.NewRect(MARGIN_LEFT, y, GRID_WIDTH, GRID_HEIGHT)

	rect := origin_rect
	rect.Y_move(-1 * GRID_HEIGHT)
	s.make_y_axis_heading(canvas, &rect, SPLITS_HAND_HEADING)

	for i := 2; i <= 11; i++ {
		rect := origin_rect
		rect.Y_move((i - 2) * GRID_HEIGHT)
		s.make_y_axis_heading(canvas, &rect, Y_axis_headings_splits()[i-2])
	}
}

// y轴标题-单独一个
func (s *SvgMaker) make_y_axis_heading(canvas *svg.SVG, rect *xmath.Rect, heading string) {
	txt := fmt.Sprintf("translate(%d,%d)", rect.X_left, rect.Y_top)
	canvas.Gtransform(txt)
	defer canvas.Gend()

	//背景
	//style := STYLE_BG_UNKOWN
	//canvas.Rect(0, 0, rect.Witdh(), rect.Height(), style)

	len := len(heading)
	if len >= 6 { //特别处理一下,否则 splits 会超出
		len = 5
	}
	//标题
	canvas.Text(rect.Witdh()-len*10, rect.Height()/2+10, heading, STYLE_TXT_Y_AXIS_HEADING)
}

// 页脚
func (s *SvgMaker) make_footer(canvas *svg.SVG) {
	height := s.canvas_height()
	txt := fmt.Sprintf("translate(0,%d)", height-H_FOOTER)
	canvas.Gtransform(txt)
	defer canvas.Gend()

	canvas.Text(s.canvas_width()/2-100, H_FOOTER/2+10, TXT_FOOTER, STYLE_FOOTER)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
