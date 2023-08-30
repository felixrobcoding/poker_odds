/*
功能：app配置
说明：
*/
package config

import (
	"path"

	"github.com/poker-x-studio/x/xlog"
	"github.com/poker-x-studio/x/xpath"

	"github.com/BurntSushi/toml"
	"github.com/poker-x-studio/x/xdebug"
)

// Load_config_app 加载配置文件
func Load_config_app() (cfg *ConfigApp, err error) {
	xlog_entry := xlog.New_entry(xdebug.Funcname())

	config := &ConfigApp{}

	filepath := path.Join(xpath.Executable_dir(), CONFIG_FOLDER, TOML_FILE_APP)
	_, err = toml.DecodeFile(filepath, config)
	if err != nil {
		xlog_entry.Errorf("配置文件 %s 失败,%s", filepath, err.Error())
		return
	}
	cfg = config

	xlog_entry.Infof("配置文件 %s is ok,", filepath)
	return
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
