package exercise

import (
	"fmt"
	"time"
)

func question1() {
	for num := 6; num < 11; num++ {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}
}

func question2() {
	for num := 10; num > 2; num-- {
		if num%2 == 1 {
			fmt.Println(num)
		}
	}

	num := 10
	for num > 2 {
		if num%2 == 1 {
			fmt.Println(num)
		}
		num--
	}
}

func checkWeekday() {
	day := time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println(day, " is a Weekend")
	default:
		fmt.Println(day, " is a Weekday")
	}
}

func Exercise2() {
	question1()
	fmt.Println()

	question2()
	fmt.Println()

	checkWeekday()
}
