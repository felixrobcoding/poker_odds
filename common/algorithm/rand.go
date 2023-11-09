/*
功能：随机
说明：
*/
package algorithm

import (
	"math/rand"
	"time"
)

const (
	_LOOP_CNT = 5
)

// Shuffle_cards 随机洗牌
func Shuffle_cards(decks int) []byte {
	cards := Raw_cards(decks)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Seed(time.Now().UnixNano())

	for i := 0; i < _LOOP_CNT; i++ {
		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
	}

	return cards
}

// Raw_cards 得到牌
func Raw_cards(decks int) []byte {
	cards := []byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, //方块
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, //梅花
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, //红心
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, //黑桃
	}
	shoe_cards := make([]byte, 0)
	for i := 0; i < decks; i++ {
		shoe_cards = append(shoe_cards, cards...)
	}
	return shoe_cards
}

func Shuffle_cards_ex(cards []byte) []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Seed(time.Now().UnixNano())

	for i := 0; i < _LOOP_CNT; i++ {
		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
	}

	return cards
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
