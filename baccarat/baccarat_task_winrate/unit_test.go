/*
功能：百家乐任务-单元测试
说明：
*/
package baccarat_task_winrate

import (
	"fmt"
	"path"
	"path/filepath"
	"testing"
)

func TestXxx(t *testing.T) {
	// dir := xpath.Executable_dir()
	// filename := "test.json"
	// path := path.Join(dir, filename)

	// shoe_cards := []byte{1, 2, 3, 4}
	// Save(path, shoe_cards)
}

func TestXxx2(t *testing.T) {
	// dir := xpath.Executable_dir()
	// filename := "test.json"
	// path := path.Join(dir, filename)

	// shoe_cards, err := Load(path)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(shoe_cards)
}

func TestXxx3(t *testing.T) {
	file := "E:\\data\\test.txt"
	dir, filename := filepath.Split(file)
	fmt.Println(dir, filename)       //获取路径中的目录及文件名 E:\data\  test.txt
	fmt.Println(filepath.Base(file)) //获取路径中的文件名test.txt
	fmt.Println(path.Ext(file))      //获取路径中的文件的后缀 .txt
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
