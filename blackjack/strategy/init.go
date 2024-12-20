/*
功能：玩法策略-初始化
说明：
*/
package strategy

import (
	"fmt"
	"github.com/felixrobcoding/poker_oddsblackjack/strategy/outputer"
	"github.com/felixrobcoding/poker_oddsblackjack/strategy/standard_strategy"
	"os"

	"github.com/poker-x-studio/x/ximage"
)

const (
	is_output = false //是否输出
)

func init() {
	output_gotable_jpeg()
}

// 输出gotable/jpeg文件
func output_gotable_jpeg() {
	if !is_output {
		return
	}
	strategy_map := standard_strategy.Get_strategy_map()

	//gotable
	outputer.Gotable_make(strategy_map)

	//svg
	svg_content := outputer.Strategy_svg_make(strategy_map)
	png_filepath, svg_filepath, err := ximage.Svg_2_png(svg_content)
	if err != nil {
		return
	}
	//删除临时文件
	os.Remove(svg_filepath)
	fmt.Println(png_filepath)

	jpeg_filepath, err := ximage.Png_2_jpeg(png_filepath)
	if err != nil {
		return
	}
	//删除临时文件
	os.Remove(png_filepath)
	fmt.Println(jpeg_filepath)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
