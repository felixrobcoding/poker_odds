# 测试条件:6副牌,闲家起始筹码:10000,最小下注:10,最大下注:1000


## 玩法策略[6副牌]

## 下注策略

### 一下注策略:全下, 10个goroutine，每个goroutine循环1000次,每次是6副牌,日志输出结果:

time="2023-07-12 16:18:06" level=info msg="min_bet:10,max_bet:1000,betting_strategy:all_in,player_min_profit:-30000.00,player_max_profit:117639.50," TAG=blackjack_task_auto_run
time="2023-07-12 16:18:06" level=info msg="sum_round:560024,sum_profit:-11121364.50,avarg_profit_per_round:-19.8587,avarg_profit_per_round/min_bet:-1.985873" TAG=blackjack_task_auto_run
time="2023-07-12 16:18:06" level=info msg="end_tick:1689155286,时间周期,diff=10,时间=0:10"


### 二 下注策略:固定额度, 10个goroutine，每个goroutine循环1000次,每次是6副牌,日志输出结果:

time="2023-07-12 16:19:05" level=info msg="min_bet:10,max_bet:1000,betting_strategy:fixed_amount,player_min_profit:-310.00,player_max_profit:325.00," TAG=blackjack_task_auto_run
time="2023-07-12 16:19:05" level=info msg="sum_round:559994,sum_profit:-124925.00,avarg_profit_per_round:-0.2231,avarg_profit_per_round/min_bet:-0.022308" TAG=blackjack_task_auto_run
time="2023-07-12 16:19:05" level=info msg="end_tick:1689155345,时间周期,diff=10,时间=0:10"

### 三 下注策略:马丁格尔策略, 10个goroutine，每个goroutine循环1000次,每次是6副牌,日志输出结果:

time="2023-07-12 16:20:02" level=info msg="min_bet:10,max_bet:1000,betting_strategy:martegal,player_min_profit:-10925.00,player_max_profit:5345.00," TAG=blackjack_task_auto_run
time="2023-07-12 16:20:02" level=info msg="sum_round:559595,sum_profit:-585920.00,avarg_profit_per_round:-1.0470,avarg_profit_per_round/min_bet:-0.104704" TAG=blackjack_task_auto_run
time="2023-07-12 16:20:02" level=info msg="end_tick:1689155402,时间周期,diff=10,时间=0:10"

### 四 下注策略:斐波那契策略, 10个goroutine，每个goroutine循环1000次,每次是6副牌,日志输出结果:

time="2023-07-12 16:21:45" level=info msg="min_bet:10,max_bet:1000,betting_strategy:fibonacci,player_min_profit:-12975.00,player_max_profit:34695.00," TAG=blackjack_task_auto_run
time="2023-07-12 16:21:45" level=info msg="sum_round:509839,sum_profit:-9736750.00,avarg_profit_per_round:-19.0977,avarg_profit_per_round/min_bet:-1.909770" TAG=blackjack_task_auto_run
time="2023-07-12 16:21:45" level=info msg="end_tick:1689155505,时间周期,diff=10,时间=0:10"

### 五 下注策略:凯利策略, 10个goroutine，每个goroutine循环1000次,每次是6副牌,日志输出结果:

time="2023-07-12 16:22:38" level=info msg="min_bet:10,max_bet:1000,betting_strategy:kelly,player_min_profit:-8525.00,player_max_profit:26425.50," TAG=blackjack_task_auto_run
time="2023-07-12 16:22:38" level=info msg="sum_round:559985,sum_profit:-5552970.00,avarg_profit_per_round:-9.9163,avarg_profit_per_round/min_bet:-0.991628" TAG=blackjack_task_auto_run
time="2023-07-12 16:22:38" level=info msg="end_tick:1689155558,时间周期,diff=10,时间=0:10"