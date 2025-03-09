package main

func getNameCounts(names []string) map[rune]map[string]int {
	sorter := make(map[rune]map[string]int)
	for _,name:= range names{
	 firstLetter :=[]rune(name)[0]//this gives a rune && rune(name[0]) gives a byte
	 if _,ok :=sorter[firstLetter];!ok{
	   sorter[firstLetter] =make(map[string]int)
	 }
	 sorter[firstLetter][name]++
	}
	return sorter
}
