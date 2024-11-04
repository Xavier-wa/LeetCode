package main

import (
	"math/rand/v2"

	"github.com/siddontang/go/log"
)

func main() {
	dp := make([][]int, 4)
	for i, _ := range dp {
		dp[i] = make([]int, 2)
		for j, _ := range dp[i] {
			dp[i][j] = rand.Int()
		}
	}
	log.Warnf("%v", dp)
	m := dp[:]
	log.Warnf("%v", m)
}
