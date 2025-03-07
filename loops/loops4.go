package main

import (
	"fmt"
)

func printPrimes(max int) {
	for i:=2;i<max;i++{
	  count :=0
	  for j:= 2;j<i/2;j++{
	    if i%j==0{
	      count++
	    }
	  }
	  if count== 1{
	   fmt.Println(i)
	  }
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}

