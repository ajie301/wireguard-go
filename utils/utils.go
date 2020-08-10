package utils

import (
	"math/rand"
	"time"
)

// 获取随机非0整型数
func getRandomUint32() uint32 {
	var i uint32 = 0
	rand.Seed(time.Now().UnixNano())
	for {
		i = rand.Uint32()
		if i != 0 {
			break
		}
	}
	return i
}

/*
  使用随机数异或混淆消息类型
  t: 原始消息类型
  m: 混淆后消息类型
  r: 随机数
*/
func MixType(t uint32) (m uint32, r uint32) {
	r = getRandomUint32()
	m = r ^ t
	return
}

/*
  还原消息类型
  m: 混淆后消息类型
  r: 随机数
  返回值: 原始消息类型
*/
func DivideType(m, r uint32) uint32 {
	return m ^ r
}
