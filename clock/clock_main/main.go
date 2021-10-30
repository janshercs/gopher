package main

import (
	"os"
	"time"

	"clock"
)

func main() {
	t := time.Now()
	clock.SVGWriter(os.Stdout, t)
}
