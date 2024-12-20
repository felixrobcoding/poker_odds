/*
功能：
说明：
*/
package svg_utils

import (
	"fmt"
	"github.com/felixrobcoding/poker_oddsbaccarat/define/BET_AREA"

	"github.com/poker-x-studio/x/xutils"
)

const (
	SCALE             = 1.5                      //缩放比
	H_PER_ROW         = 30                       //每行高
	H_HEADER          = 40                       //header高
	H_HEADING         = int(1.5 * H_PER_ROW)     //标题
	H_FOOTER          = 40                       //footer高
	TXT_FOOTER        = xutils.POKER_X_STUDIO    //页脚 文字
	TXT_HEADER        = "Baccarat压注区域统计"         //
	STYLE_BG          = "fill:#8B4500"           //svg背景
	STYLE_HEADER_TXT  = "font-size:22;fill:#FFF" //
	STYLE_HEADING_BG  = "fill:#838B83"           //svg标题背景
	STYLE_HEADING_TXT = "font-size:18;fill:#FFF" //
	STYLE_FOOTER_TXT  = "font-size:22;fill:#FFF" //
)

type WinBetAreaStat struct {
	Total_hands      int           //总手数
	Win_bet_area     BET_AREA.TYPE //压注区域
	Win_bet_area_cnt int           //压注区域个数
	Percentage       float64       //百分比
}

func (w *WinBetAreaStat) String() string {
	txt := fmt.Sprintf("[Total_hands:%d,Win_bet_area:%s,Win_bet_area_cnt:%d,Percentage:%.4f%%]", w.Total_hands, w.Win_bet_area.String(), w.Win_bet_area_cnt, w.Percentage*100)
	return txt
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
