package slice_go

import "fmt"

func Test1() {
	s := make([]int, 0)
	oldCap := cap(s)
	for i := 0; i < 2048; i++ {
		s = append(s, i)
		newCap := cap(s)
		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}
}

func Test2() {
	s := []int{1, 2}
	s = append(s, 3, 4, 5)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func Test3() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := slice[2:4:6]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s = append(s, 20, 22)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(slice)
	s = append(s, 25)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(slice)
}
