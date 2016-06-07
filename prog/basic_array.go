package prog

import "fmt"

// ArrayBasic will help to understand some basics about array
func ArrayBasic() {
	// this array has default value as 0,0,0 // we can reassign value and the same is going on in the below function
	var testArray [3]int
	var marks = [5]int{12, 5, 8, 9, 9}
	fmt.Println("Range of testArray =", len(testArray))
	fmt.Println("Range of marks =", len(marks))
	for j := 0; j < len(testArray); j++ {
		fmt.Printf("Enter the %v elements >>>>>>>", j)
		fmt.Scanf("%v", &testArray[j])
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("		Elements in array 'testArray'>>> %v \n", testArray[i])
	}
}

// Insertion for inserting the values in array dynamically and then it will print the array values as well
func Insertion() {
	var NumOfEmployee int
	fmt.Println("Enter the number of employees in Company")
	fmt.Scanf("%d", &NumOfEmployee)
	// The role of using make() and new() is indeed important ---
	var employeeName = make([]string, NumOfEmployee)
	fmt.Printf("It's the size of arrary %v\n", len(employeeName))
	i := 0
	for i < len(employeeName) {
		fmt.Print("Enter the name of employee>>>> ")
		fmt.Scanf("%s", &employeeName[i])
		i++
	}
	fmt.Println("------------------------Thank you for providing information--------------------------------")
	for j := 0; j < len(employeeName); j++ {
		fmt.Printf("working employee %v \n", employeeName[j])
	}
}

//NthInsertion is used to insert an int at nth position
func NthInsertion() {
	var a int
	fmt.Printf("Enter the size of Array>>>>>>")
	fmt.Scanf("%d", &a)

	var phoneNum = make([]int, a)
	for i := 0; i <= a; i++ {
		fmt.Print("Enter the number")
		fmt.Scanf("%d", &phoneNum[i])
		for j := 0; j <= len(phoneNum); j++ {
			fmt.Printf("")
		}
	}

	// fmt.Printf("Enter the")
	// fmt.Printf("%T  %v\n", make([]int, 10), make([]int, 10))
	// fmt.Printf("%T  %v\n", new([10]int), new([10]int))

}
