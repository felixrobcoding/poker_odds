/*
功能：玩法策略-测试
说明：
*/
package strategy

import (
	"Odds/blackjack/strategy/outputer"
	"Odds/blackjack/strategy/standard_strategy"
	"fmt"
	"os"
	"testing"

	"github.com/poker-x-studio/x/ximage"
)

var index int

func Test1(t *testing.T) {
	//chart_2_table(chart_map)
	push_messagef("%s", "dddd")

	txt := fmt.Sprintf("%4d", 1234567)
	fmt.Println(txt)
}

func push_messagef(format string, a ...string) {
	str := fmt.Sprintf(format, a)
	msg := fmt.Sprintf("轮数:%d,%s", index, str)
	fmt.Println(msg)
	index++
}

func TestSVG(t *testing.T) {
	svg_content := outputer.Strategy_svg_make(standard_strategy.Get_strategy_map())

	png_filepath, svg_filepath, err := ximage.Svg_2_png(svg_content)
	if err != nil {
		return
	}
	os.Remove(svg_filepath)
	fmt.Println(png_filepath)

	jpeg_filepath, err := ximage.Png_2_jpeg(png_filepath)
	if err != nil {
		return
	}
	os.Remove(png_filepath)
	fmt.Println(jpeg_filepath)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
