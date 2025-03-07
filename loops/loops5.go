package main

func countConnections(groupSize int) int {
	sum :=0
	for i:=0;i<groupSize;i++{
	  sum +=i
	}
	return sum
}


