package unsafe_go

import (
	"fmt"
	"leaning-golang-source-test/unsafe-local-test"
	"unsafe"
)

func Test1() {
	/*p := unsafe_local_test.Programmer{
		name:     "yang",
		age:      18,
		language: "chinese",
	}*/
	p := unsafe_local_test.GetLocalValue()
	fmt.Println("old p: ", p)
	name := (*string)(unsafe.Pointer(&p))
	*name = "taoge"
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + (unsafe.Sizeof(string("")))))
	*age = 26
	language := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + (unsafe.Sizeof(string(""))) + (unsafe.Sizeof(int(0)))))
	*language = "english"
	fmt.Println("new p: ", p)
}

func Test2() {
	s := make([]int, 9, 20)
	len := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println("len: ", *len)
	cap := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println("cap: ", *cap)

	mp := make(map[string]int)
	mp["qcrao"] = 100
	mp["stefno"] = 18
	count := (**int)(unsafe.Pointer(&mp)) //因为make()返回的就是hmap的地址，所以&mp是对地址取指针，需要（**int）
	fmt.Println("count: ", **count)
}
