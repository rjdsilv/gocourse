package intermediate

import "fmt"

func Closures() {
	sequence := adder()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()
	fmt.Println(sequence2())

	subtracterMain := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtracterMain(1))
	fmt.Println(subtracterMain(2))
	fmt.Println(subtracterMain(3))
	fmt.Println(subtracterMain(4))
	fmt.Println(subtracterMain(5))
}

func adder() func() int {
	i := 0
	fmt.Println("Previous value of i =", i)
	return func() int {
		i++
		fmt.Println("Added 1 to i")
		return i
	}
}
