package main

import (
	"fmt"
	"strings"
)

func countingValleys(n int, s string) int {
    steps := strings.Split(s, " ")
	state := 0
	valley := 0
	for _, v := range steps {
	    fmt.Println(v)
	    if v == "U" {
	        state++
	    }
	    if v == "D" {
	        state--
	    }

	    if state == 0 && v == "U" {
	        valley++
	    }
	}

    return valley
}

func main() {
	fmt.Println("Count valleys")
	n := 8
	s := "U D D D U D U U"
	fmt.Println(countingValleys(n, s))
}
