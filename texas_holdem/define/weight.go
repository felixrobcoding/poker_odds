/*
功能：权重定义
说明：
*/
package define

const (
	WEIGHT_HIGH_CARD            int64 = 1          //高牌权重
	WEIGHT_ONE_PAIR             int64 = 10         //一对权重
	WEIGHT_TWO_PAIR             int64 = 100        //两对权重
	WEIGHT_THREE_OF_A_KIND      int64 = 1000       //三条权重
	WEIGHT_STRAIGHT             int64 = 10000      //顺子权重
	WEIGHT_FLUSH                int64 = 100000     //同花权重
	WEIGHT_FULL_HOUSE           int64 = 1000000    //葫芦权重
	WEIGHT_FOUR_OF_A_KIND       int64 = 10000000   //四条权重
	WEIGHT_STRAIGHT_FLUSH       int64 = 100000000  //同花顺权重
	WEIGHT_ROYAL_STRAIGHT_FLUSH int64 = 1000000000 //皇家同花顺权重
)

//-----------------------------------------------
//					the end
//-----------------------------------------------
