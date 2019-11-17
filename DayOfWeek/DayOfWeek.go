package main

import "fmt"

func solution(S string, k int) string {
	daysToNum := map[string]int{
		"Sun": 0,
		"Mon": 1,
		"Tue": 2,
		"Wed": 3,
		"Thu": 4,
		"Fri": 5,
		"Sat": 6,
	}
	numToDays := map[int]string{
		0: "Sun",
		1: "Mon",
		2: "Tue",
		3: "Wed",
		4: "Thu",
		5: "Fri",
		6: "Sat",
	}
	return numToDays[(daysToNum[S]+k)%7]
}

func main() {
	fmt.Println(solution("Wed", 2))
	fmt.Println(solution("Sat", 23))
}
