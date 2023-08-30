# 任务说明


## 一 功能说明:
    - Blackjack统计

## 二 庄家闲家胜率统计

 - 运行环境:vscode/ubuntu
 - 运行条件:Blackjack基础策略[6副牌,庄家17点stand,闲家可以double,可以surrender,split之后,可以继续hit]
  - 开启10个goroutine,每个goroutine循环1万次,也就是一共运行10万次
  - 每次循环就是1靴牌,每靴牌就是6副牌的游戏，剩余20张的时候,1靴牌结束
  - 最小下注10,最大下注1000,原始筹码10000
  - 投注策略为固定筹码
  - 运行时间:2023-07-13 18:05:10
  - 运行时长:2:03[2分3秒]  
  - 运行结果: 

```
msg="min_bet:10,max_bet:1000,betting_strategy:fixed_amount,player_min_profit:-350.00,player_max_profit:375.00,"
msg="sum_deal_times:5500421,sum_hands:5638827,sum_bets:63034250,sum_profit:-1306545.00,hands_per_shoe:56.39,profit_per_shoe:-13.0655,profit_per_hand:-0.2317,sum_profit/sum_bets:-0.020728" 
msg="sum_player_lose_hands:2863533,sum_player_push_hands:445554,sum_player_win_hands:2329740,player_lose_hands_ratio:0.5078,player_push_hands_ratio:0.0790,player_win_hands_ratio:0.4132," 
```
