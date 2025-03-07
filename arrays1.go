package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	arr1 :=[3]string{primary,secondary,tertiary}
	arr2 :=[3]int{0,0,0}
	for i:=0; i<len(arr1);i++{
	if i==0{
	 arr2[i]=len(arr1[i])
	 continue
	}
	 arr2[i]=len(arr1[i])+arr2[i-1]
	}
	return arr1,arr2
}
