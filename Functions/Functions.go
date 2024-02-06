// In this Session we gonna see how to make functions in go language.

/*First of all it is compulsory to define package main and after that it is also mandatory
  to define a main function inside our code otherwise we will get an error same as with
  package main ,  so use these to avoid error.*/

package main

import "fmt"

// We define a function with the help of func keyword.

func main() {

	fmt.Println(show())

	// Inside our main function we are calling our show function and printing its value to see the result.

	fmt.Println(addition(9999, 3948))

	fmt.Println(addAndSub(5, 6))

	// Now both the operations addition and subtraction will gonna be perform respectively on the given numbers.

	f := func() string {

		return "I am from f function"

	}

	fmt.Println(f())

	// Lets see how to declare anonymus functions (For more details go at the bottom of the page).

	func() {

		fmt.Println("I am from anonymus function")

	}()

	/*Here we are not providing any name to the function thats why it is an anonymus function and
	here we have put the paranthesis(()) at the end where the function is closing , so this
	another way of calling the function without using fmt.println or calling the function
	separately.*/

	proRes := proAdder(10, 203, 439384, 9392, 392284, 38973498)

	fmt.Println("The pro result is ", proRes)

	// Storing the function we made below in a variable , then printing it.

}

func show() string {

	return "GoodBye World"

	// whenever we are returning something  we have to define its type in our function like we have done above.

}

func addition(x int, y int) int {

	return x + y

	// Thats how we can make a function to perform the addition , similarly can make other functions.

}

// But if we want to perform multiple operations im a single function we can do that also like this :-

func addAndSub(x int, y int) (int, int) {

	return (x + y), (x - y)

	/*Here we are returning 2 values therefore we have to define int 2 times in the function and have to put
	them in brackets otherwise we will get an error.
	In the same way we can return more than 2 things also if we want to.*/

}

/*If we dont know how many values we gonna have as parameter and we want to perform operations on all values
then we can do it like this.*/

func proAdder(values ...int) int {

	totalValues := 0

	for _, value := range values {

		totalValues += value

	}

	return totalValues

	/*So here we are taking many values , so we are using the slices for this work , as now we can
	enter as many values as we want.

	So inside this function we have take a variable values and then (...int) means here we are making
	a slicing with the help of triple dot and then defining the type  of slicing and then have written
	int as we are returing something , so it is necessary to mention the type of thing we are returning.

	Then we are taking a temporary variable which gonna hold the total of the values which we gonna enter.

	Then we are running a loop with the range keyword to itertate every item of the slicing then
	adding it to the temporary variable we made.

	Atlast we are returning the temporary variable we made to get the value

	And then we are calling this function inside our main function so that we can run it.*/

}

/*we gonna make a function and gonna store that inside a variable that function will have to be called
in main function.

There is also a concept of anonymus function in go language lets see that also.

Anonymus function is a function in which we dont provide a name to the function , just write the function
and put the code inside it.

So Lets move above and see how to do it.*/

// So That was all for the session , stay tuned for more updates.
