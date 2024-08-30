/*
功能：大路-列
说明：
*/
package big_road

import (
	"Odds/baccarat/define/BET_AREA"
)

// Col 列,列中的所有元素都相同
type Col struct {
	nodes []Node
}

// push 插入列中
func (c *Col) push(node Node) {
	c.nodes = append(c.nodes, node)
}

func (c *Col) Result_area() BET_AREA.TYPE {
	return c.nodes[0].result_area
}

// Cnt 列中的个数
func (c *Col) Cnt() int {
	return len(c.nodes)
}

// Get_node 获取节点
func (c *Col) Get_node(index int) *Node {
	if index < 0 || index >= len(c.nodes) {
		return nil
	}
	return &c.nodes[index]
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
