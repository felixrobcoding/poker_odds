/*
功能：5张奥马哈任务-计算胜率
说明：
*/
package plo5_task_win_rate

import (
	"Odds/common/COMPARE_TYPE"
	"fmt"
	"sync"
)

type Plo5WinRate struct {
	rw_mutex         sync.RWMutex
	play1_hole_cards []byte
	play2_hole_cards []byte
	board_cards      []byte
	smaller_cnt      int
	equal_cnt        int
	bigger_cnt       int
}

// Set_result 设置结果
func (w *Plo5WinRate) Set_result(ct COMPARE_TYPE.TYPE) {
	w.rw_mutex.Lock()
	defer w.rw_mutex.Unlock()

	if ct == COMPARE_TYPE.BIGGER {
		w.bigger_cnt++
	} else if ct == COMPARE_TYPE.EQUAL {
		w.equal_cnt++
	} else {
		w.smaller_cnt++
	}
}

// String 转字符串
func (w *Plo5WinRate) String() string {
	w.rw_mutex.RLock()
	defer w.rw_mutex.RUnlock()

	//总和
	sum := w.bigger_cnt + w.equal_cnt + w.smaller_cnt
	str := "{"

	str += fmt.Sprintf("sum:%d,", sum)

	ratio := float64(w.bigger_cnt) / float64(sum*1.0)
	str += fmt.Sprintf("bigger_cnt:%d,ratio:%.2f%%,", w.bigger_cnt, ratio*100.0)

	ratio = float64(w.equal_cnt) / float64(sum*1.0)
	str += fmt.Sprintf("equal_cnt:%d,ratio:%.2f%%,", w.equal_cnt, ratio*100.0)

	ratio = float64(w.smaller_cnt) / float64(sum*1.0)
	str += fmt.Sprintf("smaller_cnt:%d,ratio:%.2f%%", w.smaller_cnt, ratio*100.0)

	str += "}"
	return str
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
