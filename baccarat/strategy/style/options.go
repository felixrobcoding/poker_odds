/*
功能：形态-选项
说明：
*/
package style

type Option func(*StyleOption)

// 形态选项
type StyleOption struct {
	long_node_cnt         int
	long_bet_times        float64
	single_jump_bet_times float64
	double_jump_bet_times float64
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

func WithLongBetTimes(bet_times float64) Option {
	return func(s *StyleOption) {
		s.long_bet_times = bet_times
	}
}

func WithSingleJumpBetTimes(bet_times float64) Option {
	return func(s *StyleOption) {
		s.single_jump_bet_times = bet_times
	}
}

func WithDoubleJumpBetTimes(bet_times float64) Option {
	return func(s *StyleOption) {
		s.double_jump_bet_times = bet_times
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
