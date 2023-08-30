# 项目说明


## 一 功能说明:
    - 百家乐游戏-胜率统计

## 二 统计结果

- 统计条件:goroutine个数:50,每个goroutine循环次数:10000,每次循环牌靴中扑克副数:8
- 统计时间:time="2023-07-18 19:29:17",耗时7:15(7分15秒)
- 统计结果:

```
time="2023-07-18 19:29:17" level=info msg="min_bet:10,max_bet:1000,betting_strategy:fixed_amount,player_min_profit:-394.50,player_max_profit:394.00," TAG=baccarat_task_winrate
time="2023-07-18 19:29:17" level=info msg="sum_deal_times:40295556,sum_hands:40295556,sum_bets:402955560,sum_profit:-4567106.50,hands_per_shoe:80.59,profit_per_shoe:-9.1342,profit_per_hand:-0.1133,sum_profit/sum_bets:-0.011334" TAG=baccarat_task_winrate
time="2023-07-18 19:29:17" level=info msg="sum_player_lose_hands:18201698,sum_player_push_hands:3832132,sum_player_win_hands:18261726,player_lose_hands_ratio:0.4517,player_push_hands_ratio:0.0951,player_win_hands_ratio:0.4532," TAG=baccarat_task_winrate
time="2023-07-18 19:29:17" level=info msg="end_tick:1689685157,生命周期,diff=435,时间=7:15(7分15秒)"
```