package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}

func sliceOps() {
	fmt.Println("Creating slice")
	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	//copy的话类型须一致
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	//解构append
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}

func main() {
	// arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// fmt.Println("arr[2:6] =", arr[2:6])
	// fmt.Println("arr[:6] =", arr[:6])
	// s1 := arr[2:]
	// fmt.Println("s1 =", s1)
	// s2 := arr[:]
	// fmt.Println("s2 =", s2)

	// fmt.Println("After updateSlice(s1)")
	// updateSlice(s1)
	// fmt.Println(s1)
	// fmt.Println(arr)

	// fmt.Println("Reslice")
	// fmt.Println(s2)
	// s2 = s2[:5]
	// fmt.Println(s2)
	// s2 = s2[2:]
	// fmt.Println(s2)

	// fmt.Println("Extending slice")
	// arr[0], arr[2] = 0, 2
	// fmt.Println("arr =", arr)
	// s1 = arr[2:6]
	// s2 = s1[3:5] // [s1[3], s1[4]]
	// fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
	// 	s1, len(s1), cap(s1))
	// fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
	// 	s2, len(s2), cap(s2))

	// slice有点类似 c++ 的vector,超过cap后会重新分配
	// s3 := append(s2, 10)
	// s4 := append(s3, 11)
	// s5 := append(s4, 12)
	// fmt.Println("s3, s4, s5 =", s3, s4, s5)
	// // s4 and s5 no longer view arr.
	// fmt.Println("arr =", arr)

	sliceOps()
}
