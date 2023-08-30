/*
功能：生成获胜区域统计svg图
说明：
*/
package svg_utils

import (
	"bytes"

	"fmt"

	svg "github.com/ajstarks/svgo"
	"github.com/poker-x-studio/x/xmath"
)

const (
	W_TOTAL_HANDS_HEADING      = 100 //总手数
	W_WIN_BET_AREA_HEADING     = 120 //压注区域
	W_WIN_BET_AREA_CNT_HEADING = 160 //压注区域个数
	W_PERCENTAGE_HEADING       = 100 //百分比
)

type WinBetAreaStatSvg struct {
	y_axis_cnt int //y轴个数
}

var instance_stat *WinBetAreaStatSvg

// 单例
func Instance_win_bet_area_stat_svg() *WinBetAreaStatSvg {
	if instance_stat == nil {
		instance_stat = &WinBetAreaStatSvg{}
	}
	return instance_stat
}

// 生成svg图
// 返回:svg图字符串
func (s *WinBetAreaStatSvg) Make_svg(stats []*WinBetAreaStat) string {
	svg_writer := new(bytes.Buffer) //写入缓冲
	canvas := svg.New(svg_writer)

	s.y_axis_cnt = len(stats)

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
	s.make_body(canvas, stats)
	//footer
	s.make_footer(canvas, height)
	canvas.End()

	//测试
	//fmt.Println(svg_writer)
	return svg_writer.String()
}

// 宽度
func (s *WinBetAreaStatSvg) canvas_width() int {
	return W_TOTAL_HANDS_HEADING + W_WIN_BET_AREA_HEADING + W_WIN_BET_AREA_CNT_HEADING + W_PERCENTAGE_HEADING
}

// 高度
func (s *WinBetAreaStatSvg) canvas_height() int {
	return H_HEADER + s.y_axis_cnt*H_PER_ROW + H_HEADING + H_FOOTER //高度
}

// defs
func (s *WinBetAreaStatSvg) make_defs(canvas *svg.SVG) {
	canvas.Def()
	defer canvas.DefEnd()
}

// 整个svg图的背景
func (s *WinBetAreaStatSvg) make_bg(canvas *svg.SVG) {
	canvas.Gtransform("translate(0,0)")
	defer canvas.Gend()

	//背景填充
	canvas.Rect(0, 0, s.canvas_width(), s.canvas_height(), STYLE_BG)
}

// 头
func (s *WinBetAreaStatSvg) make_header(canvas *svg.SVG) {
	canvas.Gtransform("translate(0,0)")
	defer canvas.Gend()

	canvas.Text(s.canvas_width()/2-120, H_HEADER/2+10, TXT_HEADER, STYLE_HEADER_TXT)
}

func (s *WinBetAreaStatSvg) make_body(canvas *svg.SVG, stats []*WinBetAreaStat) {
	//标题
	s.make_headings(canvas)

	txt := fmt.Sprintf("translate(0,%d)", H_HEADER+H_HEADING)
	canvas.Gtransform(txt)

	rect := &xmath.Rect{
		X_left:   0,
		X_right:  s.canvas_width(),
		Y_top:    0,
		Y_bottom: H_PER_ROW,
	}

	for k, v := range stats {
		rct_tmp := xmath.NewRectWithCopy(rect)
		if k > 0 {
			rct_tmp.Y_move(H_PER_ROW * k)
		}
		s.make_row(canvas, rct_tmp, v)
	}
	canvas.Gend()
}

// 标题
func (s *WinBetAreaStatSvg) make_headings(canvas *svg.SVG) {
	txt := fmt.Sprintf("translate(0,%d)", H_HEADER)
	canvas.Gtransform(txt)
	defer canvas.Gend()

	//背景填充
	canvas.Rect(0, 0, s.canvas_width(), H_HEADING, STYLE_HEADING_BG)
	const y_base = 32

	//头部区域的标题
	w := 0

	canvas.Text(w+10, y_base, "总手数", STYLE_HEADING_TXT)
	w += W_TOTAL_HANDS_HEADING

	canvas.Text(w+10, y_base, "获胜压注区域", STYLE_HEADING_TXT)
	w += W_WIN_BET_AREA_HEADING

	canvas.Text(w+10, y_base, "获胜压注区域个数", STYLE_HEADING_TXT)
	w += W_WIN_BET_AREA_CNT_HEADING

	canvas.Text(w+10, y_base, "百分比", STYLE_HEADING_TXT)
	w += W_PERCENTAGE_HEADING
}

// 行
func (s *WinBetAreaStatSvg) make_row(canvas *svg.SVG, rect *xmath.Rect, stat *WinBetAreaStat) {
	y_base := rect.Center_y() + 6

	const style_txt = "font-size:16;fill:#000"
	const style_bg_1 = "fill:#DCDCDC"
	const style_bg_2 = "fill:#fff"

	//底色
	rct_total_hands := xmath.NewRect(rect.X_left, rect.Y_top, W_TOTAL_HANDS_HEADING, rect.Height())
	canvas.Rect(rct_total_hands.X_left, rct_total_hands.Y_top, rct_total_hands.Witdh(), rct_total_hands.Height(), style_bg_1)

	rct_win_bet_area := xmath.NewRectWithCopy(rct_total_hands)
	rct_win_bet_area = rct_win_bet_area.X_move_to(rct_win_bet_area.X_right).Update_width(W_WIN_BET_AREA_HEADING)
	canvas.Rect(rct_win_bet_area.X_left, rct_win_bet_area.Y_top, rct_win_bet_area.Witdh(), rct_win_bet_area.Height(), style_bg_1)

	rct_win_bet_area_cnt := xmath.NewRectWithCopy(rct_win_bet_area)
	rct_win_bet_area_cnt = rct_win_bet_area_cnt.X_move_to(rct_win_bet_area_cnt.X_right).Update_width(W_WIN_BET_AREA_CNT_HEADING)
	canvas.Rect(rct_win_bet_area_cnt.X_left, rct_win_bet_area_cnt.Y_top, rct_win_bet_area_cnt.Witdh(), rct_win_bet_area_cnt.Height(), style_bg_2)

	rct_percentage := xmath.NewRectWithCopy(rct_win_bet_area_cnt)
	rct_percentage = rct_percentage.X_move_to(rct_percentage.X_right).Update_width(W_PERCENTAGE_HEADING)
	canvas.Rect(rct_percentage.X_left, rct_percentage.Y_top, rct_percentage.Witdh(), rct_percentage.Height(), style_bg_1)

	canvas.Text(rct_total_hands.X_left+10, y_base, fmt.Sprintf("%d", stat.Total_hands), style_txt)
	canvas.Text(rct_win_bet_area.X_left+10, y_base, stat.Win_bet_area.String(), style_txt)
	canvas.Text(rct_win_bet_area_cnt.X_left+10, y_base, fmt.Sprintf("%d", stat.Win_bet_area_cnt), style_txt)
	canvas.Text(rct_percentage.X_left+10, y_base, fmt.Sprintf("%.4f%%", stat.Percentage*100.0), style_txt)

	//分隔线
	canvas.Line(rect.X_left, rect.Y_bottom, rect.X_right, rect.Y_bottom, "stroke:rgb(88,88,88);stroke-width:1")
}

// 页脚
func (s *WinBetAreaStatSvg) make_footer(canvas *svg.SVG, total_height int) {
	txt := fmt.Sprintf("translate(0,%d)", total_height-H_FOOTER)
	canvas.Gtransform(txt)
	defer canvas.Gend()

	canvas.Text(s.canvas_width()/2-120, H_FOOTER/2+10, TXT_FOOTER, STYLE_FOOTER_TXT)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
