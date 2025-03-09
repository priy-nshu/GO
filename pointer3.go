package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(analytic *Analytics,message Message){
   if message.Success{
    (*analytic).MessagesTotal++
    (*analytic).MessagesSucceeded++
   }else{
       (*analytic).MessagesTotal++
           (*analytic).MessagesFailed++
}
}

