package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
    for i ,j:=range msg{
      for _,k:=range badWords{
         if k==j{
          return i
         }
      }
    }
	return -1
	}
