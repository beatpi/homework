package main

import "fmt"

func main() {
	// s := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	s := []int{0, 0, 1, 1, 1, 1, 2, 3}
	//	k := 2
	//	fmt.Println(removeDuplicates(s))
	// fmt.Println(removeElement(s, k))
	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(nums []int) int {
	counter := 0
	double := false
	for i := 1; i < len(nums); i++ {
		if !double {
			counter++
			double = true
			nums[counter] = nums[i]
		} else {
			if nums[i] != nums[counter] {
				double = false
				counter++
				nums[counter] = nums[i]
			} else {
				double = true
			}
		}
	}
	return counter + 1
}

/*
func removeDuplicates(nums []int) int {
	counter := 0
	double := false
	for i := 1; i < len(nums); i++ {
		if !double {
			counter++
			double = true
			nums[counter] = nums[i]
		} else {
			if nums[i] != nums[counter] {
				double = false
				counter++
				nums[counter] = nums[i]
			} else {
				continue
			}
		}
	}
	return counter + 1
}
*/
/*
	func removeDuplicates(nums []int) int {
		counter := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[counter] {
				counter++
				nums[counter] = nums[i]
			}
		}

		return counter + 1
	}
*/
func removeElement(nums []int, val int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[counter] = nums[i]
			counter++
		}

	}
	return counter
}
