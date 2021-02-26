// package main

// //"errors"
// import (
// 	"fmt"
// )

 func main() {
// 	// fmt.Println("Hello World!")
// 	// x := 10.2
// 	// s := "This is string text"
// 	// b := true
// 	// fmt.Printf("Num X1 : %f %T", x, x)
// 	// fmt.Printf("\nString s : %s %T", s, s)
// 	// fmt.Printf("\nBoolean b : %t %T", b, b)

// 	// x := 10
// 	// if x = 3; x%2 == 0 {
// 	// 	fmt.Println("x is even")
// 	// } else {
// 	// 	fmt.Println("x is odd")
// 	// }

// 	// r := add(10, 15)
// 	// fmt.Printf("Result is %d", r)
// 	// fmt.Printf("\nResult is %d\n", add(10, 10))
// 	// fmt.Println(add(10, 24))

// 	// x := -1
// 	// switch x {
// 	// case -1:
// 	// 	fmt.Println("x should be positive")
// 	// 	//fallthrough
// 	// case 10:
// 	// 	fmt.Println("x is even")
// 	// default:
// 	// 	fmt.Println("x")
// 	// }

// 	// for i := 0; i < 5; i++ {
// 	// 	fmt.Printf("\nx should be +ve %d", i)
// 	// }

// 	//ints:= [5]int{5,8,9,12,3}
// 	//  for i:=0;i<5;i++ {
// 	// 	 fmt.Printf("%d",ints[i])
// 	//  }
// 	//  for key,value:=range ints{
// 	//      fmt.Printf("\nArray element %d is %d", key, value)
// 	//  }
// 	// ints2:= ints[2:5]
// 	// ints2[2]=50
// 	// ints3:=[]int{23,45,46,7}

// 	// n:=10
// 	// ints3=make([]int,n)
// 	// for key,value:=range ints3{
// 	// 	     fmt.Printf("\nArray element %d is %d", key, value)
// 	// 	  }

// 	// map1 := make(map[string]int)
// 	// map1["one"] = 1
// 	// map1["two"] = 2
// 	// for key, value := range map1 {
// 	// 	fmt.Printf("\nArray element %s is %d", key, value)
// 	// }

// 	// var d int
// 	// _,err:=fmt.Scanf("%d",&d)
// 	// if err !=nil {
// 	// 	fmt.Printf("\nError in input")
// 	// 	return
// 	// }
// 	// fmt.Printf("\n input %d",d)
// 	// r,err:=divide(10,d)
// 	// if (err!=nil){
// 	// 	fmt.Printf("Results in error")
// 	// 	return
// 	// }
// 	// fmt.Printf("\nDivide %d",r)
// 	var per1 Person
// 	per1 = Person{firstName: "Anand", lastName: "babu", age: 20, address: "Bangalore"}
// 	//fmt.Printf("Person Detail %q",emp1)
// 	//Print(emp1)
// 	per1.Print()
// 	var emp1 Employee
// 	emp1.project
// 	emp1.empNo = "201"
// 	emp1.dept = "Healthcare"
// 	emp1.Person = per1
// 	emp1.Print()
// 	var empA1 empActivity
// 	empA1 = emp1
// 	empA1.applyLeave(2)
// 	empA1.assignProject("EHS")
 }

// // func add(x int, y int) int {
// // 	return x + y
// // }

// // func divide(x int, y int) (int, error) {
// // 	if y==0{
// // 		return 0, errors.New("Divisor is zero")
// // 	}
// // 	return x/y,nil
// // }
// type Person struct {
// 	firstName string
// 	lastName  string
// 	age       int
// 	address   string
// }

// func Print(p Person) {
// 	fmt.Printf("Person Detail {First Name: %s, Lastname: %s, Age: %d, Address: %s", p.firstName, p.lastName, p.age, p.address)
// }
// func (p Person) Print() {
// 	fmt.Printf("Person Detail {First Name: %s, Lastname: %s, Age: %d, Address: %s", p.firstName, p.lastName, p.age, p.address)
// }

// type Employee struct {
// 	Person
// 	empNo        string
// 	dept         string
// 	project      string
// 	leaveApplied int
// }

// func (e Employee) Print() {
// 	fmt.Printf("\nEmployee detail {Emp no: %s, dept: %s, Project: %s Leaves: %d", e.empNo, e.dept, e.project, e.leaveApplied)
// }

// type empActivity interface {
// 	assignProject(project string)
// 	applyLeave(i int)
// }

// func (e Employee) assignProject(project string) {
// 	e.project = project
// }
// func (e Employee) applyLeave(i int) {
// 	e.leaveApplied = i
// }

