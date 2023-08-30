/*
功能：app配置-定义
说明：
*/
package config

type debug struct {
}

// 配置文件
type ConfigApp struct {
	Debug debug `toml:"debug"`
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
