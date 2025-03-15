package main

import "time"

func processMessages(messages []string) []string {
	// ?
	reciever :=make([]string,0,len(messages))
	ch := make(chan string,len(messages))
	for _,msg:=range messages{
	 go func(m string){
	  ch<-process(m)
	 }(msg)
	  reciever=append(reciever,<-ch)
	}
	return reciever
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
