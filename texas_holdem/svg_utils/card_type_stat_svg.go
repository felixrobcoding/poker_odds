/*
功能：生成牌型统计svg图
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
	W_TOTAL_RUN_TIMES_HEADING = 100 //总运行次数
	W_DEAL_CARD_CNT_HEADING   = 80  //发牌张数
	W_CARD_TYPE_HEADING       = 100 //牌型
	W_TYPE_CNT_HEADING        = 120 //当前牌型个数
	W_PERCENTAGE_HEADING      = 100 //百分比
)

type CardTypeStatSvg struct {
	header     string //header
	y_axis_cnt int    //y轴个数
}

var instance_stat *CardTypeStatSvg

// 单例
func Instance_card_type_stat_svg() *CardTypeStatSvg {
	if instance_stat == nil {
		instance_stat = &CardTypeStatSvg{}
	}
	return instance_stat
}

// 生成svg图
// 返回:svg图字符串
func (c *CardTypeStatSvg) Make_svg(header string, stats []*CardTypeStat) string {
	svg_writer := new(bytes.Buffer) //写入缓冲
	canvas := svg.New(svg_writer)

	c.header = header
	c.y_axis_cnt = len(stats)

	widht := c.canvas_width()
	height := c.canvas_height()
	canvas.Startview(int(float64(widht)*SCALE), int(float64(height)*SCALE), 0, 0, widht, height)

	//defs
	c.make_defs(canvas)
	//背景
	c.make_bg(canvas)
	//header
	c.make_header(canvas)
	//body
	c.make_body(canvas, stats)
	//footer
	c.make_footer(canvas, height)
	canvas.End()

	//测试
	fmt.Println(svg_writer)
	return svg_writer.String()
}

// 宽度
func (c *CardTypeStatSvg) canvas_width() int {
	return W_TOTAL_RUN_TIMES_HEADING + W_DEAL_CARD_CNT_HEADING + W_CARD_TYPE_HEADING + W_TYPE_CNT_HEADING + W_PERCENTAGE_HEADING
}

// 高度
func (c *CardTypeStatSvg) canvas_height() int {
	return H_HEADER + c.y_axis_cnt*H_PER_ROW + H_HEADING + H_FOOTER //高度
}

// defs
func (c *CardTypeStatSvg) make_defs(canvas *svg.SVG) {
	canvas.Def()
	defer canvas.DefEnd()
}

// 整个svg图的背景
func (c *CardTypeStatSvg) make_bg(canvas *svg.SVG) {
	canvas.Gtransform("translate(0,0)")
	defer canvas.Gend()

	//背景填充
	canvas.Rect(0, 0, c.canvas_width(), c.canvas_height(), STYLE_BG)
}

// 头
func (c *CardTypeStatSvg) make_header(canvas *svg.SVG) {
	canvas.Gtransform("translate(0,0)")
	defer canvas.Gend()

	canvas.Text(c.canvas_width()/2-120, H_HEADER/2+10, c.header, STYLE_HEADER_TXT)
}

func (c *CardTypeStatSvg) make_body(canvas *svg.SVG, stats []*CardTypeStat) {
	//标题
	c.make_headings(canvas)

	txt := fmt.Sprintf("translate(0,%d)", H_HEADER+H_HEADING)
	canvas.Gtransform(txt)

	rect := &xmath.Rect{
		X_left:   0,
		X_right:  c.canvas_width(),
		Y_top:    0,
		Y_bottom: H_PER_ROW,
	}

	for k, v := range stats {
		rct_tmp := xmath.NewRectWithCopy(rect)
		if k > 0 {
			rct_tmp.Y_move(H_PER_ROW * k)
		}
		c.make_row(canvas, rct_tmp, v)
	}
	canvas.Gend()
}

// 标题
func (c *CardTypeStatSvg) make_headings(canvas *svg.SVG) {
	txt := fmt.Sprintf("translate(0,%d)", H_HEADER)
	canvas.Gtransform(txt)
	defer canvas.Gend()

	//背景填充
	canvas.Rect(0, 0, c.canvas_width(), H_HEADING, STYLE_HEADING_BG)
	const y_base = 32

	//头部区域的标题
	w := 0

	canvas.Text(w+10, y_base, "总运行次数", STYLE_HEADING_TXT)
	w += W_TOTAL_RUN_TIMES_HEADING

	canvas.Text(w+10, y_base, "发牌张数", STYLE_HEADING_TXT)
	w += W_DEAL_CARD_CNT_HEADING

	canvas.Text(w+10, y_base, "牌型", STYLE_HEADING_TXT)
	w += W_CARD_TYPE_HEADING

	canvas.Text(w+10, y_base, "当前牌型个数", STYLE_HEADING_TXT)
	w += W_TYPE_CNT_HEADING

	canvas.Text(w+10, y_base, "百分比", STYLE_HEADING_TXT)
	w += W_PERCENTAGE_HEADING
}

// 行
func (c *CardTypeStatSvg) make_row(canvas *svg.SVG, rect *xmath.Rect, stat *CardTypeStat) {
	y_base := rect.Center_y() + 6

	const style_txt = "font-size:16;fill:#000"
	const style_bg_1 = "fill:#DCDCDC"
	const style_bg_2 = "fill:#fff"

	//底色
	rct_total_run_times := xmath.NewRect(rect.X_left, rect.Y_top, W_TOTAL_RUN_TIMES_HEADING, rect.Height())
	canvas.Rect(rct_total_run_times.X_left, rct_total_run_times.Y_top, rct_total_run_times.Witdh(), rct_total_run_times.Height(), style_bg_1)

	rct_deal_card_cnt := xmath.NewRectWithCopy(rct_total_run_times)
	rct_deal_card_cnt = rct_deal_card_cnt.X_move_to(rct_deal_card_cnt.X_right).Update_width(W_DEAL_CARD_CNT_HEADING)
	canvas.Rect(rct_deal_card_cnt.X_left, rct_deal_card_cnt.Y_top, rct_deal_card_cnt.Witdh(), rct_deal_card_cnt.Height(), style_bg_2)

	rct_card_type := xmath.NewRectWithCopy(rct_deal_card_cnt)
	rct_card_type = rct_card_type.X_move_to(rct_card_type.X_right).Update_width(W_CARD_TYPE_HEADING)
	canvas.Rect(rct_card_type.X_left, rct_card_type.Y_top, rct_card_type.Witdh(), rct_card_type.Height(), style_bg_1)

	rct_card_type_cnt := xmath.NewRectWithCopy(rct_card_type)
	rct_card_type_cnt = rct_card_type_cnt.X_move_to(rct_card_type_cnt.X_right).Update_width(W_TYPE_CNT_HEADING)
	canvas.Rect(rct_card_type_cnt.X_left, rct_card_type_cnt.Y_top, rct_card_type_cnt.Witdh(), rct_card_type_cnt.Height(), style_bg_2)

	rct_percentage := xmath.NewRectWithCopy(rct_card_type_cnt)
	rct_percentage = rct_percentage.X_move_to(rct_percentage.X_right).Update_width(W_PERCENTAGE_HEADING)
	canvas.Rect(rct_percentage.X_left, rct_percentage.Y_top, rct_percentage.Witdh(), rct_percentage.Height(), style_bg_1)

	canvas.Text(rct_total_run_times.X_left+10, y_base, fmt.Sprintf("%d", stat.Total_run_times), style_txt)
	canvas.Text(rct_deal_card_cnt.X_left+10, y_base, fmt.Sprintf("%d", stat.Deal_card_cnt), style_txt)
	canvas.Text(rct_card_type.X_left+10, y_base, stat.Type.String(), style_txt)
	canvas.Text(rct_card_type_cnt.X_left+10, y_base, fmt.Sprintf("%d", stat.Type_cnt), style_txt)
	canvas.Text(rct_percentage.X_left+10, y_base, fmt.Sprintf("%.2f%%", stat.Percentage), style_txt)

	//分隔线
	canvas.Line(rect.X_left, rect.Y_bottom, rect.X_right, rect.Y_bottom, "stroke:rgb(88,88,88);stroke-width:1")
}

// 页脚
func (c *CardTypeStatSvg) make_footer(canvas *svg.SVG, total_height int) {
	txt := fmt.Sprintf("translate(0,%d)", total_height-H_FOOTER)
	canvas.Gtransform(txt)
	defer canvas.Gend()

	canvas.Text(c.canvas_width()/2-120, H_FOOTER/2+10, TXT_FOOTER, STYLE_FOOTER_TXT)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
