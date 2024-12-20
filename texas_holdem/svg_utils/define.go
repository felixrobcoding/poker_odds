/*
功能：
说明：
*/
package svg_utils

import (
	"github.com/felixrobcoding/poker_oddstexas_holdem/define/CARD_TYPE"

	"github.com/poker-x-studio/x/xutils"
)

const (
	SCALE                   = 1.5                         //缩放比
	H_PER_ROW               = 30                          //每行高
	H_HEADER                = 40                          //header高
	H_HEADING               = int(1.5 * H_PER_ROW)        //标题
	H_FOOTER                = 40                          //footer高
	TXT_FOOTER              = xutils.POKER_X_STUDIO       //页脚 文字
	TXT_HEADER_UTLIMATE     = "Ultimate Texas holdem牌型统计" //
	TXT_HEADER_TEXAS_HOLDEM = "Texas holdem牌型统计"          //
	STYLE_BG                = "fill:#8B4500"              //svg背景
	STYLE_HEADER_TXT        = "font-size:22;fill:#FFF"    //
	STYLE_HEADING_BG        = "fill:#838B83"              //svg标题背景
	STYLE_HEADING_TXT       = "font-size:18;fill:#FFF"    //
	STYLE_FOOTER_TXT        = "font-size:22;fill:#FFF"    //
)

type CardTypeStat struct {
	Total_run_times int            //总运行次数
	Deal_card_cnt   int            //发牌张数
	Type            CARD_TYPE.TYPE //牌型
	Type_cnt        int            //当前牌型个数
	Percentage      float64        //百分比
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
