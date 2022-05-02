package main

import (
	"fmt"
	_ "strconv"
)

/*
Task: Given a list of 4 integers, find all possible valid 24 hour times (eg: 12:34) that the given integers can be assembled into and return the total number of valid times.
You can not use the same number twice.
Times such as 34:12 and 12:60 are not valid.
Provided integers can not be negative.
Notes: Input integers can not be negative.
Input integers can yeald 0 possible valid combinations.
Example:
	Input: [1, 2, 3, 4]
	Valid times: ["12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "23:14", "23:41", "21:34", "21:43"]
	Return: 10
*/

func possibleTimes(digits []int) int {
	// Since the input will always be 4 digits, we can get away with constant time
	possibleTimes := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if i == j || i == k || j == k {
					continue
				}
				l := 6 - i - j - k

				// Time format: IJ:KL

				if digits[i]*10+digits[j] > 23 {
					continue
				}
				if digits[k]*10+digits[l] > 59 {
					continue
				}
				possibleTimes++
			}
		}
	}

	return possibleTimes
}

func main() {
	// Example test cases.
	fmt.Println(possibleTimes([]int{1, 2, 3, 4}))
	fmt.Println(possibleTimes([]int{9, 1, 2, 0}))
	fmt.Println(possibleTimes([]int{2, 2, 1, 9}))
	fmt.Println(possibleTimes([]int{6, 7, 8, 9}))
}
