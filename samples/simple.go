package main

import (
	"fmt"

	".."
)

func main() {
	fuzzy := fuzzy.New()
	fuzzy.Add("work")
	fuzzy.Add("world")
	fuzzy.Add("wolf")
	fuzzy.Add("well")

	fmt.Printf("for '%s' result is '%s'\n", "wolf", fuzzy.Search("wolf"))
	fmt.Printf("for '%s' result is '%s'\n", "welf", fuzzy.Search("welf"))
	fmt.Printf("for '%s' result is '%s'\n", "ell", fuzzy.Search("ell"))
}
