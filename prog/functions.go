package prog

import "fmt"

// if arrguments have same data type then we can combine thier data type
// func sum(a int, b int)

func Sum(a, b int) int {
	//using printf to print the variable within string and in the last of string mentioning \n to move the cursor in next line
	fmt.Printf("Sum of %v and %T is = %v\n", a, b, a+b)

	return a + b
}

func Table(a int) int {
	i := 1
	// Declaring the C outside the block because it would be required outside the block to return value
	var c int
	for i <= 10 {
		c = i * a
		i++
		fmt.Println(c)
	}
	return c
}

func OddEven(a int) string {
	if a%2 == 0 {
		fmt.Printf("Number %v is Even\n", a)
		return "Even"
	}

	fmt.Printf("Number %v is Odd\n", a)
	return "Odd"
}

func Test() {
	// this function is written to understand the difference between Print functions.
	var a, b int = 19, 20
	// z := 14.0
	fmt.Print(a + b)
	fmt.Printf("sam---------\n  %v %v", a, b)
	// fmt.Println("sam---------%v %v", a, b)
	// fmt.Printf("Value of z = %d and type = %T", z, z)

}

// Golang also supports the multiple return type values
func Swap(a, b int) (int, int) {
	fmt.Println(b, a)
	return a, b
}

//different types of using for loops
// func forLoop(i int ) int{
// for i:=1; i<= 100; i++{
// 	fmt.Println(i)
// 	}
// return i
// }

func ForLoop() {
	for i := 1; i <= 9; i++ {
		j := 1
		for j <= i {
			fmt.Print(i)
			j++
		}
		fmt.Println()
	}
}

func Rangeloop() {
	num := [5]int{1, 2, 3, 4, 5}
	for i, x := range num {
		fmt.Printf("Value of x = %d and i = %d \n", x, i)
	}

}

// Check whether a number is prime or not
func CheckPrime(a int) string {
	for i := 2; i < a; i++ {
		if a >= i {
			if a%i == 0 {
				// fmt.Printf("Number %d is not prime\n",a)
				return "Not Prime"
				// break
			}
		}
	}
	fmt.Printf("Number %d is prime\n", a)
	return "Prime"
}

// we can return the named vaules
func NamedValue(a int) (sum int) {
	sum = 10
	sum = sum * a
	fmt.Println(sum)
	return
}

// To understand typecasting
func Conversion() {

	var a int = 10
	var b float32 = 12.22
	var sum int
	sum = a + int(b)
	fmt.Println(sum)
}

// Make variables constant
func Constant() {
	const (
		z, j int    = 44, 55
		x    int    = 1
		name string = "Samarpit"
	// x = z+j // const declaration error
	)

	fmt.Println(z, j, name, x)
	const i int = 12
	var b int = 1
	// i = i+b  // const declaration error
	fmt.Println(i, b)
}

// Simple infinite for loop it is easys
func ForEver() {
	for {
		fmt.Println("samarpi")

	}
}

// func main() {
// 	sum(10, 11)
// 	table(10)
// 	OddEven(15)
// 	test()
// 	swap(1, 5)
// 	forLoop()
// 	rangeloop()

// Prime function will be used by passing arrguments
// for a:=3; a<=100; a++{
// 	checkPrime(a)swap }

// namedValue(10)
// conversion()
// constant()

// Please don't run this function
// ForEver()
