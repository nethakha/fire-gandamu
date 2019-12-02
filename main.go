package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"reflect"

	"github.com/seehuhn/mt19937"
)

// 3回以上燃え上がったらガンダム
func main() {
	a := []int{}

	for {
		if rnd() == 0 {
			fmt.Println("🔥燃え上がれ")
			a = burn(a, 0)
		} else if check(a) == false {
			fmt.Println("💧鎮火")
			a = burn(a, 1)
			continue
		} else {
			fmt.Println("🤖ガンダム")
			break
		}
	}
}

func check(a []int) bool {
	if reflect.DeepEqual(a, []int{0, 0, 0}) {
		return true
	}
	return false
}

func burn(a []int, b int) []int {
	if len(a) == 3 {
		a = a[1:]
	}
	a = append(a, b)
	return a
}

func rnd() int64 {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rng := rand.New(mt19937.New())
	rng.Seed(seed.Int64())
	return rng.Int63n(2)
}
