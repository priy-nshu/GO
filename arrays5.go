package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
    cost := []float64{}
	for _,c :=range costs {
	 if c.day ==day{ 
	  cost =append(cost,c.value) }
	}
	return cost
}
