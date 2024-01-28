// In this file we gonna see how to make a struct in go language.

// Struct is almost same as we have learnt in solidity , just difference of syntax rest is same.

// Remember we dont have classes in go language therefore we use struct.

//But like classes , struct doesnt have any super keyword nor have any parent concept.

//There is no inheritance possible in go language , so neither struct can have inheritance like classes.

package main

import "fmt"

/*Remember to make the struct outside of the function like we do in solidity
We have to use type keyword and then the name of the struct and then struct keyword and then in curly braces
we define all the fields which we are required in our struct and also their type like in solidity
Just difference is that we have to declare the type after the field name in go but in solidity we define
the type before the name.

And remember to give the struct name starting with capital letter and field names in Pascal Case otherwise
we will get an error.*/

type Employee struct {
	EmpId     int
	EmpName   string
	EmpNumber []int
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

// we have taken slicing of int for EmpNumber as there could be nore than one contact number.

func main() {

	/*Now inside function we gonna declare a variable and inside that we gonna store the values of struct like
	  there would be a particular struct for every person , so we have to follow that approach.*/

	emp1 := Employee{
		1,
		"Rashi",
		[]int{9393949393, 939929292},
	}

	/*Now inside the struct we defined the desired values according to the need and remember to put a comma
	after the last entry otherwise we will get an error .*/

	fmt.Println(emp1)

	// print the variable inside which we stored the struct and we will get the desired result.

	emp2 := Employee{
		EmpId:     2,
		EmpName:   "Aman",
		EmpNumber: []int{74747383, 99394492},
	}

	/* this is the another way of providing values , like in key value pair , these are the keys which
	we define inside our struct , we can use any way which we want.*/

	fmt.Println(emp2)

	// If we want to change the value of a  particular field then we can do it like this:-

	emp2.EmpName = "Ahana"

	fmt.Println(emp2)

	// After printing we can see that our EmpName has been changed , similarly we can change all the fields.

	user1 := User{

		Name:   "yash",
		Age:    22,
		Email:  "yo@gmail.com",
		Status: true,
	}

	fmt.Println(user1)

	// Just practicing the struct one more time nothing more than that

	fmt.Printf("The Details of the user1 are %+v\n", user1)

	// IMPORTANT NOTE

	/*When we are using struct and want to get the details of any thing
	which is made through struct(user1 in this case) remember to use fmt.Printf and %+v.

	Dont try to use %v only , use + sign also like we did above otherwise it will not gonna work.

	this will give us perfect detailed view of any struct so always remember to use this.

	For checking run this file and at the end see the difference between 2 printed lines
	one will give us really less details and another one will give us full details
	regarding the thing we were looking to get the details.*/

	/*If we want to print only some particular fields from the thing made with the struct(user1 in this case)
	Then we can do it like this:-*/

	fmt.Printf("The Name And Email Of The User1 is %v and %v respectively", user1.Name, user1.Email)

	// This will give us the desired result.

}

//stay tuned for more updates
