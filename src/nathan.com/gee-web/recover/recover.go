package main

import "fmt"

func recoverRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover panic error %s", err)
		}
	}()

	fmt.Println("before panic")
	panic("crash")
	fmt.Println("after panic")
}

func main() {
	recoverRecover()
}
