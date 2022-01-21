package main

import unsafe_go "leaning-golang-source-test/unsafe-go"

type Programmer struct {
	name     string
	age      int
	language string
}

func main() {
	//context_go.Test1()
	//context_go.Test2()
	//context_go.Test3()
	//unsafe_go.Test1()
	unsafe_go.Test2()
}

type MaxQueue struct {
	List    []int
	MaxList []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		List:    make([]int, 0),
		MaxList: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	length := len(this.MaxList)
	if length == 0 {
		return -1
	}
	return this.MaxList[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.List = append(this.List, value)
	length := len(this.MaxList)
	for length != 0 && this.MaxList[length-1] < value {
		this.MaxList = this.MaxList[:length-1]
	}
	this.MaxList = append(this.MaxList, value)
}

func (this *MaxQueue) Pop_front() int {
	lenList := len(this.List)
	if lenList == 0 {
		return -1
	}
	val := this.List[0]
	this.List = this.List[1:]
	if val == this.MaxList[0] {
		this.MaxList = this.MaxList[1:]
	}
	return val
}
