/*
功能：21点
说明：
*/
package blackjack

import (
	"github.com/poker-x-studio/x/xlog"

	"github.com/sirupsen/logrus"
)

const (
	_TAG  = "blackjack" //游戏名称
	DECKS = 6           //牌副数
)

var xlog_entry = xlog.New_entry(_TAG)
var is_debug = true

func Logrus() *logrus.Entry {
	return xlog_entry
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
