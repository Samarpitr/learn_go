package prog

import "fmt"

func sam(i int) string {

	p := &i
	fmt.Println(i, *p, p)
	*p = 21

	fmt.Printf("ths sfds %v", i)
	return "Done"
}

// commandline input
func input() {
	var a int
	fmt.Println("Enter a number")
	fmt.Scanf("%v", &a)
	fmt.Printf("Entered number is %v \n", a)
}
func main() {
	// sam(13)
	// array()
	input()
}
