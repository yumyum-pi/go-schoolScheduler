package main

import (
	"math/rand"
	"time"

	"github.com/yumyum-pi/go-schoolScheduler/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cmd.Execute()
}
