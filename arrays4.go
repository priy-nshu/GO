package main

func sum(nums ...int) int {
	 s:=0
	for i:=0;i<len(nums);i++{
	 s +=nums[i];
	}
	return s
}
