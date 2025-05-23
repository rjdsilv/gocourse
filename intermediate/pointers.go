package intermediate

import "fmt"

func Pointers() {
	var ptr *int
	var a int = 10
	ptr = &a

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
