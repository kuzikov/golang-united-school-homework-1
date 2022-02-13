package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	var greet = emoji.Sprint("Hello :world_map:!")
	fmt.Printf("%s\n", greet)
}
