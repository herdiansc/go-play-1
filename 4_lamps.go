package main

import (
	"fmt"
)

func main() {
	fmt.Println("Switches")
	switches := []int{}
	for i := 0;i<100;i++ {
	    switches = append(switches, (i+1)*-1)
	}
	for i:=0;i<100;i++ {
	    for j:=0;j<len(switches);j++ {
	        if i==0 {
	            switches[j] = switches[j]*-1
	        }
	        if i==1 && (j+1)%2==0 {
	            switches[j] = switches[j]*-1
	        }
	        if i>1 && (j+1)%3==0 {
	            switches[j] = switches[j]*-1
	        }
	    }
	}
	on := 0
	for _, v := range switches {
	    if v > 0 {
	        on++
	    }
	}
	
	fmt.Println(on)
}
