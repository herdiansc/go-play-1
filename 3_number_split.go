package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
)

func split(num int) []int {
    numStr := strconv.Itoa(num)
    nums := strings.Split(numStr, "")
    results := []int{}
    fmt.Println(len(nums))
    for i, v := range nums {
        digit, _ := strconv.Atoi(v)
        n := (len(nums)-(i+1))
        results = append(results, digit*int(math.Pow10(n)))
    }
    return results    
}

func main() {
	fmt.Println("Integer")
	num := 1234
	fmt.Println(split(num))
}
