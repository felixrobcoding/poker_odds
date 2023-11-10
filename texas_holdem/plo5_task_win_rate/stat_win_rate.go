/*
功能：5张奥马哈任务-统计plo5胜率
说明：
*/
package plo5_task_win_rate

import (
	"Odds/common/COMPARE_TYPE"
	"fmt"
	"sync"
)

// 统计plo5胜率
type StatPlo5WinRate struct {
	rw_mutex         sync.RWMutex
	play1_hole_cards []byte
	play2_hole_cards []byte
	board_cards      []byte
	smaller_cnt      int
	equal_cnt        int
	bigger_cnt       int
}

// Set_result 设置结果
func (s *StatPlo5WinRate) Set_result(ct COMPARE_TYPE.TYPE) {
	s.rw_mutex.Lock()
	defer s.rw_mutex.Unlock()

	if ct == COMPARE_TYPE.BIGGER {
		s.bigger_cnt++
	} else if ct == COMPARE_TYPE.EQUAL {
		s.equal_cnt++
	} else {
		s.smaller_cnt++
	}
}

// String 转字符串
func (s *StatPlo5WinRate) String() string {
	s.rw_mutex.RLock()
	defer s.rw_mutex.RUnlock()

	//总和
	sum := s.bigger_cnt + s.equal_cnt + s.smaller_cnt
	str := "{"

	str += fmt.Sprintf("sum:%d,", sum)

	ratio := float64(s.bigger_cnt) / float64(sum*1.0)
	str += fmt.Sprintf("bigger_cnt:%d,ratio:%.2f%%,", s.bigger_cnt, ratio*100.0)

	ratio = float64(s.equal_cnt) / float64(sum*1.0)
	str += fmt.Sprintf("equal_cnt:%d,ratio:%.2f%%,", s.equal_cnt, ratio*100.0)

	ratio = float64(s.smaller_cnt) / float64(sum*1.0)
	str += fmt.Sprintf("smaller_cnt:%d,ratio:%.2f%%", s.smaller_cnt, ratio*100.0)

	str += "}"
	return str
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
