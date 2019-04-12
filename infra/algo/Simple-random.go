package algo

import (
	"math/rand"
	"time"
)

//每个红包最小保证金额
const min = int64(1)

//简单随机算法
//红包数量，红包金额(单位：分)
func SimpleRand(count, amount int64) int64 {
	//当红包数量剩余一个的时候，就直接返回剩余红包金额
	if count == 1 {
		return amount
	}
	//计算最大可调度金额(红包总金额-红包数量*每个红包最小保证金额)
	max := amount - min*count
	//添加随机数种子
	rand.Seed(time.Now().UnixNano())
	time.Sleep(1)
	x := rand.Int63n(max) + min
	return x

}
