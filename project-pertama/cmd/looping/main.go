package main

import (
	"fmt"
	"time"
)

func loopingWithFor() {
	for i := 0; i < 5; i++ {
		println(i)
	}
}
func loopingWithForWhile() {
	i := 0
	for i < 5 {
		println(i)
		i++
	}
}
func loopingWithBreak() {
	nilai := 3
	i := 0
	for {

		println(i)
		if i == nilai {
			break
		}
		i++
	}
}

func loopingWithContinue() {
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			continue
		}
		println(i)
	}
}

// linear search O(n)
func loopingWithInput(n int, search int) {
	for i := 0; i < n; i++ {
		if i == search {
			println(i)
			break
		}
	}
}

// binary search
func loopingWithInputBinary(n int, search int) {
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if mid == search {
			println(mid)
			break
		} else if mid < search {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
}

// O(n^2)
func searchWithLinear(input []int, search int) {

	var result [][]int

	// [0][1], [0,2] [0,3] [0,4] [0,5] [0,6] [0,7] [0,8] [0,9]
	// [1,1], [1,2] [1,3] [1,4] [1,5] [1,6] [1,7] [1,8] [1,9]
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if input[i]+input[j] == search {
				result = append(result, []int{input[i], input[j]})
			}
		}
	}
	fmt.Println(result)
}

// O(n)
func searchWithMap(input []int, search int) [][]int {
	var result [][]int
	var m = make(map[int]int)

	for i := 0; i < len(input); i++ {
		if _, ok := m[search-input[i]]; ok {
			result = append(result, []int{input[i], search - input[i]})
		}
		m[input[i]] = i
	}
	return result
}

func testLopping() {
	// input := []int{10, 30, 20, 11, 89, 91, 60, 55, 9, 63} // n = 10
	input := []int{}

	for i := range 100000 {
		input = append(input, i)
	}

	start := time.Now()
	fmt.Println("start: ", start)
	// searchWithLinear(input, 500)
	searchWithMap(input, 500)
	elapsed := time.Since(start)
	fmt.Println("elapsed: ", elapsed)

	// loopingWithFor()
	// start := time.Now()
	// fmt.Println("start: ", start)
	// loopingWithInput(1000000, 9999)

	// elapsed := time.Since(start)
	// fmt.Println("elapsed: ", elapsed)

	// start := time.Now()
	// fmt.Println("start: ", start)
	// loopingWithInputBinary(1000000, 9999)

	// elapsed := time.Since(start)

	// fmt.Println("elapsed: ", elapsed)
}

func main() {

	// example with n + 1 problems
	// studentList = select * from students; 300ms , n = 100

	// n + 1 problems
	// for i , value := range studentList {
	//   studentPoint = select * from student_points where student_id = value.id; 30ms , 30ms * 100 = 3000ms
	//}

	// solution with n + 1 problems
	// studentList = select * from students; 300ms , n = 100
	// studentListId = []int{}
	// studentPoint = select * from student_points where student_id in (studentListId); 500ms
}
