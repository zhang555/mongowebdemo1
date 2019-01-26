package util

import (
	"time"
	"math/rand"
	"strconv"
)

var (
	Random *rand.Rand
)

func InitRandom() {
	Random = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandomUint() uint {
	return uint(Random.Uint32())
}

//生成六位的随机数
func Random6NumToString() string {
	var num int
	num = Random.Int()
	for num < 100000 {
		num = Random.Int()
	}
	return strconv.Itoa(num)[0:6]
}

func Random3NumWithout4ToString() string {
	//var (
	//	//num   int
	//	s     = strconv.Itoa(Random.Int())
	//	le    = len(s)
	//	count int
	//	by    bytes.Buffer
	//)
	//
	//for _, value := range s {
	//	if string(value) != "4" {
	//
	//		by.WriteString(string(value))
	//
	//		count ++
	//		if count > 6 {
	//			break
	//		}
	//	}
	//}
	//
	//return by.String()
	return ""
}
