# 说明

## 一 功能说明:

- 1.1 庄家闲家胜率统计
- 1.2 增加百家乐svg输出大路图
- 1.3 特殊形态有提示，比如长龙，单跳，双跳

## 二 下注区域胜率对比:

GO_ROUTINE_CNT      = 10    //goroutine个数
LOOP_TIMES          = 10    //每个goroutine循环次数

### 2.1 去除断龙检测

msg="sum_deal_times:7059,max_bet_amount:80,sum_hands:7059,sum_bets:77550,sum_profit:-1933.00,hands_per_shoe:70.59,profit_per_shoe:-19.3300,profit_per_hand:-0.2738,sum_profit/sum_bets:-0.0249" func=Odds/baccarat/baccarat_task_winrate.stat file="D:/private/poker_odds/baccarat/baccarat_task_winrate/task.go:204" TAG=baccarat_task_winrate
msg="sum_player_lose_hands:3226,sum_player_push_hands:654,sum_player_win_hands:3179,player_lose_hands_ratio:45.7005%,player_push_hands_ratio:9.2648%,player_win_hands_ratio:45.0347%," func=Odds/baccarat/baccarat_task_winrate.stat file="D:/private/poker_odds/baccarat/baccarat_task_winrate/task.go:205" TAG=baccarat_task_winrate

### 2.2 增加 断龙检测

msg="sum_deal_times:7047,max_bet_amount:80,sum_hands:7047,sum_bets:77250,sum_profit:-2481.50,hands_per_shoe:70.47,profit_per_shoe:-24.8150,profit_per_hand:-0.3521,sum_profit/sum_bets:-0.0321" func=Odds/baccarat/baccarat_task_winrate.stat file="D:/private/poker_odds/baccarat/baccarat_task_winrate/task.go:204" TAG=baccarat_task_winrate
msg="sum_player_lose_hands:3299,sum_player_push_hands:652,sum_player_win_hands:3096,player_lose_hands_ratio:46.8142%,player_push_hands_ratio:9.2522%,player_win_hands_ratio:43.9336%," func=Odds/baccarat/baccarat_task_winrate.stat file="D:/private/poker_odds/baccarat/baccarat_task_winrate/task.go:205" TAG=baccarat_task_winrate

增加 断龙检测后，胜率下降了


## 三 大路上列统计:

统计条件：1000次循环，每次循环中一靴牌中有8副扑克，也就是一共1000靴牌的情况
|列节点数|列数/总列数|百分比|说明|
|---|---|---|---|
|1|16346/32446|50.3791%|总手数63868[去掉和局的情况]|
|2|8294/32446|25.5625||
|3|3963/32446|12.2141%||
|4|1993/32446|6.1425%||
|5|912/32446|2.8108%||
|6|504/32446|1.5534%||
|7|211/32446|0.6503%||
|8|122/32446|0.3760%||
|9|42/32446|0.1294%||
|10|27/32446|0.0832%||
|11|16/32446|0.0493%||
|12|6/32446|0.0185%||
|13|5/32446|0.0154%||
|14|2/32446|0.0062%||
|15|1/32446|0.0031%||
|16|2/32446|0.0062%||

