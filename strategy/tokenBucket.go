package strategy

import "fmt"

type TokenBucket struct {
	capacity   uint16 // max amount of tokens that can be held in the bucket
	refillRate uint16 // amount of tokens refilled per second
}

func Hello() {
	fmt.Println("Hello")
}

func ConfigureTokenBucket(capacity, refillRate uint16) {
	tb := TokenBucket{capacity, refillRate}
	fmt.Println(tb)
}
