/*
功能：map转table输出
说明：
*/
package common

import (
	"fmt"
	"sort"
	"strings"

	"github.com/liushuochen/gotable"
)

type item struct {
	key  string
	data int
}

type MapStrInt map[string]int

var cols = []string{"key", "data"}

func Map_2_table(m MapStrInt) string {
	table, err := gotable.Create(cols...)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	order_items := map_2_slice_and_order(m)
	for _, v := range order_items {
		row := map[string]string{
			cols[0]: v.key,
			cols[1]: fmt.Sprintf("%d", v.data),
		}

		err := table.AddRow(row)
		if err != nil {
			fmt.Println(err.Error())
			return err.Error()
		}
	}

	return table.String()
}

// 排序
func map_2_slice_and_order(m MapStrInt) []item {
	items := make([]item, 0)
	for k, v := range m {
		e := item{
			key:  k,
			data: v,
		}
		items = append(items, e)
	}

	//排序,大的在前
	sort.SliceStable(items, func(i, j int) bool {
		if items[j].data == items[i].data {
			return strings.Compare(items[j].key, items[i].key) < 0
		}
		return items[j].data < items[i].data
	})
	return items
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
