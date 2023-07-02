package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now().Add(-time.Hour)
	time.Sleep(1 * time.Second)
	t2 := time.Now()

	fmt.Println(t1, t2)
	if t2.Sub(t1).Hours() >= 1 {
		fmt.Println(t2.Sub(t1).Hours())
	}
}
