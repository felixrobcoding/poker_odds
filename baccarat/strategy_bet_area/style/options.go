/*
功能：形态-选项
说明：
*/
package style

type Option func(*StyleOption)

// 形态选项
type StyleOption struct {
	long_node_cnt int
}

var style_option *StyleOption

func NewStyleOption(opts ...Option) *StyleOption {
	s := &StyleOption{}
	for _, v := range opts {
		v(s)
	}
	return s
}

func WithLongNodeCnt(node_cnt int) Option {
	return func(s *StyleOption) {
		s.long_node_cnt = node_cnt
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
