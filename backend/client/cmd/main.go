package main

import (
	"github.com/par1ram/client/internal"
)

func main() {
	client := internal.NewClient()
	client.Run()
}
