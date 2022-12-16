package main

import (
	"fmt"
	"github.com/Sigumaa/shukujitsu-jp"
)

func main() {
	entries, err := shukujitsu.AllEntries()
	if err != nil {
		panic(err)
	}
	for _, e := range entries {
		fmt.Printf("%s = %s\n", e.YMD, e.Name)
	}
}
