package main

func bulkSend(numMessages int) float64 {
	var total float64 =0.0
	for i := 0; i < numMessages; i++{
	   total =total + (1+0.01*float64(i))
	}
	return total
}
