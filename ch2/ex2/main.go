// Converter converts feet and inches.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := Feet(t)
		i := Inch(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, f.ToInches(), i, i.ToFeet())
	}
}
