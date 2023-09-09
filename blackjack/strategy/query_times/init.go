/*
功能：查询次数
说明：
*/
package query_times

import (
	"Odds/blackjack/strategy/node"
	"Odds/blackjack/strategy/outputer"
	"Odds/blackjack/strategy/standard_strategy"
	"sync"
)

type QueryTimes struct {
	times_map map[string]int //查询次数统计
	mutex     sync.Mutex     //同步
}

var instance QueryTimes

// 单例
func Instance() *QueryTimes {
	if instance.times_map == nil {
		instance.init(standard_strategy.Get_strategy_map())
	}
	return &instance
}

// 查询次数初始化
func (q *QueryTimes) init(strategy_map map[string]*node.Node) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.times_map == nil {
		q.times_map = make(map[string]int, 0)
	}

	for _, v := range strategy_map {
		key := v.Make_key()
		q.times_map[key] = 0
	}
}

// Increase 查询次数增加
func (q *QueryTimes) Increase(key string) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.times_map[key]++
}

// Output_svg_query_times 输出svg
func (q *QueryTimes) Output_svg_query_times() string {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	svg_content := outputer.Query_times_svg_make(standard_strategy.Get_strategy_map(), q.times_map)
	return svg_content
}

// Output_svg_never_query_times 从没有被查询到
func (q *QueryTimes) Output_svg_never_query_times() string {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	//没查询的key
	zero_times_map := make(map[string]int, 0)
	strategy_map := standard_strategy.Get_strategy_map()
	for k, _ := range strategy_map {
		times, ok := q.times_map[k]
		if ok && (times <= 0) {
			zero_times_map[k] = 0
		}
	}

	svg_content := outputer.Query_times_svg_make(standard_strategy.Get_strategy_map(), zero_times_map)
	return svg_content
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
