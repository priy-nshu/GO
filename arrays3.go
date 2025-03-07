package main

func getMessageCosts(messages []string) []float64 {
	cost :=make([]float64, cap(messages))//does the same thing as len
	for i:=0;i<len(messages);i++{
	 cost[i] = 0.01*float64(len(messages[i]))
	}
	return cost
}

