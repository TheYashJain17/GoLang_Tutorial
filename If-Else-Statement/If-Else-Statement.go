//In this Session we gonna see that how can we use if-else statement in go.

// We have already use if-else statement in many languages , now lets see how to use if-else statement in go.

// Lets understand if-else statement with loops to understand it best.

package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		if i%2 == 0 {

			fmt.Println("Even Values Are", i)

		} else {

			fmt.Println("Odd Values Are", i)

		}

	}

	/*So this is how we are using if-else with for loop.
	We have declared a for loop which will run till number 10.
	Now inside that for loop we are using if-else statement
	Inside that if-else statement we are saying if i is dividing by 2 and we are getting 0 as remainder
	in that case print even values as when the remainder will be 0 it means the value
	is perfectly divisible by 2 and it means it is a even number , so that we are printing those numbers
	otherwise number would be odd and will go inside else condition and will be printed as odd number.*/

	// Remember one thing , not to do this error:-

	/*i := 4

	if i==4 {

		fmt.Println("GoodBye world");

	}
	else{

		fmt.println("Hello world")

	}*/

	/*Here you can see that we have put the else in another line , we can do this thing in other languages
	But in case of go we will get an error by doing this , so remember to declare else in same line
	where the if last bracket is.*/

	// If we want to use more than 1 if condition , we can also do that like this:-

	i := 10

	if i == 5 {

		fmt.Println("5 It is")

	} else if i == 9 {

		fmt.Println("9 it is")

	} else {

		fmt.Println("10 it is")

	}

	// this is how we can make if-else ladder if we want.

	// And remember we can use all the operands like or(||) and And(&&) like we do in other languages.

	// We can also initialise a variable and put a condition with it together in go , lets see how to do it :-

	if num := 3; num < 10 {

		fmt.Println("Number is less than 10")

	} else {

		fmt.Println("Number is greater than 10")

	}

	/*Here we can see that we are initialising a new variable name num and after that putting a semicolon(;)
	it is important to use this and after that we can define the condition like we did above.
	And Thats how we define condition with initalising a new variable.*/

}

//stay tuned for more updates
