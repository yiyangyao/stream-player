package main

import "log"

/**
限流：bucket token 算法
client发送请求时，先从bucket中获取一个token,获取成功则获取stream连接
*/

type ConnLimiter struct {
	limitCount int
	bucket     chan int
}

func NewConnLimiter(limitCount int) *ConnLimiter {
	return &ConnLimiter{
		limitCount: limitCount,
		bucket:     make(chan int, limitCount),
	}
}

func (cl *ConnLimiter) GetStreamConn() bool {
	if len(cl.bucket) >= cl.limitCount {
		log.Printf("reached the rate limitation of stream connection. please contact nathan2012@163.com")
		return false
	}

	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseStreamConn() {
	c := <-cl.bucket
	log.Printf("new connection is coming: %d", c)
}
