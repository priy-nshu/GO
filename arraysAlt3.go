package main

import(
    "strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i,msg:=range messages{
	 messages[i].tags =tagger(msg)
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
    lower :=strings.ToLower(msg.content)
    if strings.Contains(lower,"urgent"){
     tags =append(tags,"Urgent")
    }
    if strings.Contains(lower,"sale") || strings.Contains(lower,"promo"){
     tags =append(tags,"Promo")
    }
	return tags
}
