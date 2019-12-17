package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"reflect"

	"flag"
	"github.com/seehuhn/mt19937"
)

type Info struct {
	Name    string
	Usage   string
	Version string
}

func main() {
	flag.Parse()
	f := flag.Arg(0)

	app := Info{
		Name:    "Fire Gandamn",
		Usage:   "Let's Burning",
		Version: "0.0.1",
	}

	fmt.Println(app.Name)

	switch f {
	case "web":
		fmt.Println("Start Web Server/n http://localhost:8080")
		http.HandleFunc("/", Handler)
		http.ListenAndServe(":8080", nil)
	default:
		fmt.Println(Logic())
	}

}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, Logic())
}

func Logic() string {
	var result string
	a := []int{}

	for {
		if BurnCheck(a) {
			result += "🤖ガンダム"
			break
		}

		if Random() {
			result += "🔥燃え上がれ"
			BurnQueue(&a, 0)
		} else {
			result += "💧"
			BurnQueue(&a, 1)
		}

		result += "\n"
	}
	return result
}

func BurnCheck(a []int) bool {
	if reflect.DeepEqual(a, []int{0, 0, 0}) {
		return true
	}
	return false
}

func BurnQueue(a *[]int, b int) {
	if len(*a) == 3 {
		*a = (*a)[1:]
	}
	*a = append(*a, b)
}

func Random() bool {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rng := rand.New(mt19937.New())
	rng.Seed(seed.Int64())

	var rand int64 = rng.Int63n(2)

	if rand == 1 {
		return true
	}

	return false
}
