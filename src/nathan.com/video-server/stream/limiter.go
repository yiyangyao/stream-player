package main

import "log"

// bucket token 算法

type ConnLimiter struct {
	connNum int
	bucket  chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		cc, make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.connNum {
		log.Printf("Reached the rate limitation.")
		return false
	}

	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connection coming: %d", c)
}
