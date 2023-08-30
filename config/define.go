/*
功能：配置定义
说明：
*/
package config

const (
	CONFIG_FOLDER = "config" //config目录
	TOML_FILE_APP = "app.toml"
)

type logFile struct {
	Filename string `toml:"filename"`
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
