package intermediate

import "fmt"

func Pointers() {
	var ptr *int
	var a = 10
	ptr = &a

	if ptr == nil {
		fmt.Println("ptr is nil")
	}

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	modifyValue(ptr)

	fmt.Println(*ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	*ptr++
}
