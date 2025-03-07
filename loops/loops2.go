package main

func maxMessages(thresh int) int {
	total :=0
	msg :=0
	for i:=0;total+(100+i) <=thresh;i++{
	 msg++
	 total +=100+i
	}
	return msg
}
