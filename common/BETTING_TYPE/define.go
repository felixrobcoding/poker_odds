/*
功能：下注类型
说明：

https://www.pinnacle.com/zh-cn/betting-articles/betting-strategy/staking-one-method-to-improve-your-betting/2962vhe9w3jpj7x7

策略1：每一次都奋力一搏
每一次都投注所有资金。优点是可以快速获得丰厚收益。缺点？只要输一次就会输光所有资金，退出游戏。

策略2：固定注额
每一次都投注固定的金额，不管赢了多少钱，也不改变。在此示例中，注额是$100。如果你的胜率是55%，赔率是2，此方法意味着你大幅降低了输光所有注额的机会。遗憾的是，这意味着你的彩金只能以“缓慢而稳定”的方式增加。

策略3：马丁格尔策略
任何一次失败后，加倍下注，利用下一次注额的彩金弥补之前的损失。这种方式的彩金增加速度比固定注额方法快（因为你加倍下注，弥补之前的损失）。但是，如果出现连续失败，你仍必须继续加倍下注，弥补之前的损失，你很快就会投注了一大笔资金。

策略 4：斐波纳契策略
按照斐波纳契顺序增加注额，使用下一次注额的彩金弥补之前的损失。此方法和马丁格尔体育博彩方法有类似的缺点，但是它可以减少连续失败时注额的增加速度（因此，赢奖速度也下降）。

策略 5：比例投注
投注与你的优势成比例的资金在此模拟中，我们使用凯利准则公式计算比例投注。采用这种方法，注额等于优势除于赔率。在此示例中，优势是10%，赔率是一比一，10 / 1等于10。
因此，注额是$1000的10%：$100。如果投注成功，下一个注额将提高到$110，亦即1100总资金的10%。这意味着彩金增加速度比固定注额系统快，损失速度也有所下降。
*/
package BETTING_TYPE

type TYPE int

const (
	ERROR        TYPE = 0
	ALL_IN       TYPE = 1 //全下
	FIXED_AMOUNT TYPE = 2 //固定额度
	MARTEGAL     TYPE = 3 //马丁格尔:输了加倍
	FIBONACCI    TYPE = 4 //斐波那契:累加【每一个数都是前二个数的和。头二项是0和1，此数列的前几项如下： 0, 1, 1, 2, 3, 5, 】
	KELLY        TYPE = 5 //凯利准则:每次投注总筹码的某一百分比
	MIN          TYPE = ALL_IN
	MAX          TYPE = KELLY
)

//是否有效
func Is_valid(t TYPE) bool {
	if t >= MIN && t <= MAX {
		return true
	}
	return false
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
