package exercise

import (
	"fmt"
)

func Exercise3() {
	nums := make([]int, 0, 6)
	for i := 1; i < 10; i += 2 {
		nums = append(nums, i)
	}

	nums = append([]int{0}, nums...)
	fmt.Println(nums)

	removeItem := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	fmt.Println(removeItem)
	fmt.Println(nums)

	favoriteFood := make(map[string][]string)
	favoriteFood["Alex"] = []string{"icecream", "prata"}
	favoriteFood["Bob"] = []string{"pizza", "taco"}

	for key, value := range favoriteFood {
		fmt.Println(key, value)
	}
}
