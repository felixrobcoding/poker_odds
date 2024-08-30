/*
功能：百家乐任务-读写文件
说明：
*/
package baccarat_task_winrate

import (
	"io"
	"os"
)

// 保存文件
func Save_file(shoe_cards []byte, path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(shoe_cards)
	if err != nil {
		return err
	}
	//fmt.Println(len)
	return nil
}

// Load 加载文件
func Load_file(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
