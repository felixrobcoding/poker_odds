/*
功能：大路-svg
说明：
*/
package big_road

import (
	"Odds/baccarat/define/BET_AREA"
	"bytes"

	"fmt"

	svg "github.com/ajstarks/svgo"
	"github.com/poker-x-studio/x/xmath"
	"github.com/poker-x-studio/x/xutils"
)

const (
	SCALE             = 1.5                      //缩放比
	GRID_WIDTH        = 22                       //格子的宽度
	GRID_HEIGHT       = 22                       //格子的高度
	NODE_CIRCLE_R     = 10                       //节点半径
	LEFT_SPACE        = 20                       //左边留白
	RIGHT_SPACE       = 20                       //右边留白
	MIN_ROW_CNT       = 6                        //最小行数
	MIN_WIDTH         = 200                      //最小宽度
	H_HEADER          = 40                       //header高
	H_HEADING         = 20                       //标题
	H_FOOTER          = 40                       //footer高
	TXT_FOOTER        = xutils.POKER_X_STUDIO    //页脚 文字
	TXT_HEADER        = "百家乐大路图"                 //
	STYLE_BG          = "fill:#FFF"              //svg背景
	STYLE_HEADER_TXT  = "font-size:22;fill:#000" //
	STYLE_HEADING_BG  = "fill:#FFF"              //svg标题背景
	STYLE_HEADING_TXT = "font-size:18;fill:#000" //
	STYLE_FOOTER_TXT  = "font-size:22;fill:#000" //
)

type BigRoadSvg struct {
	x_axis_cnt    int    //x轴个数
	y_axis_cnt    int    //y轴个数
	heading_txt   string //文字
	is_view_index bool   //是否显示序号
}

var instance_big_road_svg *BigRoadSvg

// 单例
func Instance_big_road_svg() *BigRoadSvg {
	if instance_big_road_svg == nil {
		instance_big_road_svg = &BigRoadSvg{}
	}
	return instance_big_road_svg
}

// 生成svg图
// 返回:svg图字符串
func (s *BigRoadSvg) Make_svg(bigroad *BigRoad, is_view_index bool) string {
	svg_writer := new(bytes.Buffer) //写入缓冲
	canvas := svg.New(svg_writer)

	s.x_axis_cnt = bigroad.Col_cnt()
	s.y_axis_cnt = bigroad.Col_max_node_cnt()
	if s.y_axis_cnt < MIN_ROW_CNT {
		s.y_axis_cnt = MIN_ROW_CNT
	}
	s.heading_txt = fmt.Sprintf("总手数:%d[不包含Tie]", bigroad.Total_cnt())
	s.is_view_index = is_view_index

	widht := s.canvas_width()
	height := s.canvas_height()
	canvas.Startview(int(float64(widht)*SCALE), int(float64(height)*SCALE), 0, 0, widht, height)

	//defs
	s.make_defs(canvas)
	//背景
	s.make_bg(canvas)
	//header
	s.make_header(canvas)
	//body
	s.make_body(canvas, bigroad)
	//footer
	s.make_footer(canvas, height)
	canvas.End()

	//测试
	//fmt.Println(svg_writer)
	return svg_writer.String()
}

// 宽度
func (s *BigRoadSvg) canvas_width() int {
	space := LEFT_SPACE + RIGHT_SPACE
	if GRID_WIDTH*s.x_axis_cnt > MIN_WIDTH {
		return GRID_WIDTH*s.x_axis_cnt + space
	}
	return MIN_WIDTH + space
}

// 高度
func (s *BigRoadSvg) canvas_height() int {
	return H_HEADER + H_HEADING + s.y_axis_cnt*GRID_HEIGHT + H_FOOTER //高度
}

// defs
func (s *BigRoadSvg) make_defs(canvas *svg.SVG) {
	canvas.Def()
	defer canvas.DefEnd()
}

// 整个svg图的背景
func (s *BigRoadSvg) make_bg(canvas *svg.SVG) {
	canvas.Gtransform("translate(0,0)")
	defer canvas.Gend()

	//const style_line = "stroke:rgb(88,88,88);stroke-width:1"
	const style_line = "stroke:#DCDCDC;stroke-width:1"

	//背景填充
	canvas.Rect(0, 0, s.canvas_width(), s.canvas_height(), STYLE_BG)

	//分隔线
	y_top := H_HEADER + H_HEADING
	for y := 0; y <= s.y_axis_cnt; y++ {
		canvas.Line(LEFT_SPACE, y_top+y*GRID_HEIGHT, s.canvas_width()-RIGHT_SPACE, y_top+y*GRID_HEIGHT, style_line)
	}

	x_left := LEFT_SPACE
	for x := 0; x <= s.x_axis_cnt; x++ {
		canvas.Line(x_left+x*GRID_WIDTH, y_top, x_left+x*GRID_WIDTH, s.canvas_height()-H_FOOTER, style_line)
	}
}

// 头
func (s *BigRoadSvg) make_header(canvas *svg.SVG) {
	canvas.Gtransform("translate(0,0)")
	defer canvas.Gend()

	canvas.Text(s.canvas_width()/2-120, H_HEADER/2+10, TXT_HEADER, STYLE_HEADER_TXT)
}

func (s *BigRoadSvg) make_body(canvas *svg.SVG, bigroad *BigRoad) {
	//标题
	s.make_headings(canvas)

	txt := fmt.Sprintf("translate(0,%d)", H_HEADER+H_HEADING)
	canvas.Gtransform(txt)

	//起始节点
	init_rect := &xmath.Rect{
		X_left:   LEFT_SPACE,
		X_right:  LEFT_SPACE + GRID_WIDTH,
		Y_top:    0,
		Y_bottom: 0 + GRID_HEIGHT,
	}

	index := 1

	//列
	for col_index := 0; col_index < bigroad.Col_cnt(); col_index++ {
		col := bigroad.Get_col(col_index)

		//行
		for row_index := 0; row_index < col.Cnt(); row_index++ {
			if col.Bet_area() == BET_AREA.PLAYER {
				rct_tmp := xmath.NewRectWithCopy(init_rect)
				rct_tmp.X_move(GRID_WIDTH * col_index)
				rct_tmp.Y_move(GRID_HEIGHT * row_index)
				s.make_node(canvas, rct_tmp, col.Bet_area(), index)
				index++
				continue
			}
			if col.Bet_area() == BET_AREA.BANKER {
				rct_tmp := xmath.NewRectWithCopy(init_rect)
				rct_tmp.X_move(GRID_WIDTH * col_index)
				rct_tmp.Y_move(GRID_HEIGHT * row_index)
				s.make_node(canvas, rct_tmp, col.Bet_area(), index)
				index++
				continue
			}
		}
	}

	canvas.Gend()
}

// 标题
func (s *BigRoadSvg) make_headings(canvas *svg.SVG) {
	txt := fmt.Sprintf("translate(%d,%d)", LEFT_SPACE, H_HEADER)
	canvas.Gtransform(txt)
	defer canvas.Gend()

	//背景填充
	canvas.Rect(0, 0, s.canvas_width(), H_HEADING, STYLE_HEADING_BG)

	//头部区域的标题
	canvas.Text(0, 0, s.heading_txt)
}

// 节点
func (s *BigRoadSvg) make_node(canvas *svg.SVG, rect *xmath.Rect, bet_area BET_AREA.TYPE, index int) {
	const style_player = "stroke:blue;stroke-width:2;fill:none"
	const style_banker = "stroke:red;stroke-width:2;fill:none"

	if bet_area == BET_AREA.PLAYER {
		canvas.Circle(rect.Center_X(), rect.Center_y(), NODE_CIRCLE_R, style_player)
	}
	if bet_area == BET_AREA.BANKER {
		canvas.Circle(rect.Center_X(), rect.Center_y(), NODE_CIRCLE_R, style_banker)
	}

	const x_space = -6
	const y_space = 4
	if s.is_view_index {
		canvas.Text(rect.Center_X()+x_space, rect.Center_y()+y_space, fmt.Sprintf("%d", index))
	}
}

// 页脚
func (s *BigRoadSvg) make_footer(canvas *svg.SVG, total_height int) {
	txt := fmt.Sprintf("translate(0,%d)", total_height-H_FOOTER)
	canvas.Gtransform(txt)
	defer canvas.Gend()

	canvas.Text(s.canvas_width()/2-120, H_FOOTER/2+10, TXT_FOOTER, STYLE_FOOTER_TXT)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
