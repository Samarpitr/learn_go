package main

func array_basic() {
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

// function for inserting the values in array dynamically and then it will print the array values as well
func insertion() {
	var num_of_employee int
	fmt.Println("Enter the number of employees in Company")
	fmt.Scanf("%d", &num_of_employee)
	// The role of using make() and new() is indeed important ---
	var employee_name = make([]string, num_of_employee)
	fmt.Printf("It's the size of arrary %v\n", len(employee_name))
	i := 0
	for i < len(employee_name) {
		fmt.Print("Enter the name of employee>>>> ")
		fmt.Scanf("%s", &employee_name[i])
		i++
	}
	fmt.Println("------------------------Thank you for providing information--------------------------------")
	for j := 0; j < len(employee_name); j++ {
		fmt.Printf("working employee %v \n", employee_name[j])
	}
}

//insertion of an element at Nth position
// func insertion_nth()  {
// 	var phone_no = [9]{9,5,2,4,4,5,5,1,5}
// }
// func main(){
// 	// insertion()
// 	// array_basic()
// }
