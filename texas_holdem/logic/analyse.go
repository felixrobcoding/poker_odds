/*
功能：分析牌型[对外接口]
说明：
*/
package logic

// 分析牌型
func Analyse(cards []byte) *AnalyseItem {
	if len(cards) <= 0 {
		return nil
	}

	item := NewAnalyseItem(cards)
	if item == nil {
		return nil
	}
	return item
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
