package main

import(
 "strings"
)

func countDistinctWords(messages []string) int {
	list :=make(map[string]bool)
	for _,message :=range messages{
	 words :=strings.Fields(strings.ToLower(message))
	  for _,word:=range words{
	   list[word]=true
	  }
	 }
	return len(list)
}
