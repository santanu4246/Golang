package main

import "fmt"


func main() {

	// var n int;

	// fmt.Println("Enter the size:-");
	// fmt.Scan(&n);

	// numbers := make([]int,n)

	// fmt.Println("Enter the elements:-");
	// for i := 0; i < n; i++{
	// 	fmt.Scan(&numbers[i]);
	// }

	// fmt.Println(numbers)

	// loop over a slice
	// numbers := []int{10, 20, 30, 40}

	// for i, value := range numbers {
	// 	fmt.Println(i, value)
	// }

	//delete eliment from index 2

	// numbers := []int{10, 20, 30, 40, 50}

	// numbers = append(numbers[:2], numbers[3:]...)

	// fmt.Println(numbers)

	//copy function
	nums := make([]int, 0, 5)
	nums = append(nums, 1)
	nums2 := make([]int, len(nums))

	copy(nums2, nums)

	fmt.Println(nums, nums2)
}
