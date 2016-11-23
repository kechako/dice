package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func help() (int, error) {
	return 1, fmt.Errorf("%s [Max value]", os.Args[0])
}

func _main() (int, error) {
	var max int = 1000

	args := os.Args[1:]
	if len(args) > 0 {
		_, err := fmt.Sscanf(args[0], "%d", &max)
		if err != nil || max < 1 {
			return help()
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Int()%max + 1)

	return 0, nil
}

func main() {
	code, err := _main()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(code)
	}
}
