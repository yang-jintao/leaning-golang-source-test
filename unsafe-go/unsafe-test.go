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
