package main

import (
	"fmt"
)

func sockMerchant(n int, arr []int) int {
	occ := make(map[int]int)
	pair := 0
	for i := 0; i<len(arr); i++ {
		num := arr[i]
		if _, ok := occ[num]; ok {
		    pair++
		    delete(occ, num)
		} else {package main

import (
	"fmt"
)

func sockMerchant(n int, arr []int) int {
	occ := make(map[int]int)
	pair := 0
	for i := 0; i<len(arr); i++ {
		num := arr[i]
		if _, ok := occ[num]; ok {
		    pair++
		    delete(occ, num)
		} else {
			occ[num] = 1
		}
		
	}
	return pair
}

func main() {
	fmt.Println("Find pair of socks")
	n := 11
	arr := []int{1,2,3,4,2,3,1,6,7,4,1}
	fmt.Println(sockMerchant(n, arr))
}
