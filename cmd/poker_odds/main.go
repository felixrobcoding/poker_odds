/*
功能：
说明：
*/
package main

import (
	"github.com/felixrobcoding/poker_oddsbaccarat/baccarat_task_winrate"

	"fmt"

	"github.com/poker-x-studio/x/xdebug"
	"github.com/poker-x-studio/x/xlog"

	"github.com/jessevdk/go-flags"
)

// 命令行参数
type CmdParams struct {
	xdebug.CmdParamsBase
}

const (
	VERSION = "poker odds v2024.8.28" //版本信息
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(VERSION)

	var err error

	//解析参数
	var cmd_params CmdParams
	if _, err = flags.Parse(&cmd_params); err != nil {
		panic("Parse err:" + err.Error())
	}

	if !cmd_params.Debug {
		xdebug.Set_release()
	}
	//初始化日志
	xlog.Init_logrus(cmd_params.Log_filename)

	//加载配置文件
	if err = load_config_file(); err != nil {
		panic("加载配置文件,失败,err:" + err.Error())
	}

	//德州扑克任务-统计公共牌
	//go texas_holdem_task_board_cards.Start()
	//终极德州扑克任务-统计7张牌出现对子的概率win_bet_area_stats
	//go ultimate_texas_holdem_task_7_cards_pair.Start()
	//blackjack任务-胜率统计
	//go blackjack_task_winrate.Start()
	//blackjack任务-dealer爆牌率
	//go blackjack_task_dealer_bust_rate.Start()

	//baccarat任务-胜率统计
	go baccarat_task_winrate.Start()
	//baccarat任务-统计庄家/闲家点数分布
	//go baccarat_task_points.Start()

	//4张奥马哈任务-计算胜率
	//go plo4_task_win_rate.Start()
	//5张奥马哈任务-计算胜率
	//go plo5_task_win_rate.Start()

	// 等待进程结束信号
	xdebug.Wait_for_signal()
}

// load_config_file 加载配置文件
func load_config_file() error {
	var err error

	//
	app.Config_app, err = config.Load_config_app()
	if err != nil {
		return fmt.Errorf("读取配置文件错误,%s", err.Error())
	}

	return nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
