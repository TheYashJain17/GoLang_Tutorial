// In this session we gonna learn about Defer Keyword.

// Defer is basically used to postponed things to the last.
// Lets understand this statement with the help of an example.

package main

import "fmt"

func main() {

	defer fmt.Println("I")
	fmt.Println("Am")
	fmt.Println("Lucifer")

	/*When we will run this we will see that in which line we have used defer keyword that line will be
	executed at last.
	So according to this "I" will be printed at last instead of "Lucifer".*/

	/*Basically we use this in case of database.
	Suppose we have established a connection with the database and at last we have to close it
	But it is possible that we may forget to close the connection later , so its better to close
	the connection in starting but should execute at last , this could be done with the help of
	defer keyword.
	Lets see how we gonna do it.*/

	defer fmt.Println("DataBase Connection closing...")
	fmt.Println("Establishing Database Connection...")
	fmt.Println("Connection Established Succesfully")

	/*Now with this example we can see  we are closing the connection at starting but with the help of
	defer keyword we are postponing it to the last , thats how defer keyword works for us.*/

	// Now lets see what would happen if there are more than 1 defer keywords used.

	defer fmt.Println("World")

	defer fmt.Println("One")

	defer fmt.Println("Two")

	defer fmt.Println("Hello")

	// Now according to us the result will be Hello world one two.

	// Now lets run this and see the result.

	/*After running this the result we got is Hello Two One World.
	strange , Lets understand the reason behind it.*/

	/*When there are more than 1 defer in our code those lines will be executed in reverse order
	like which happened here.

	lets take the example of here ,
	Here the world was in the first line with defer keyword so it is printed in last.
	similarly
	Hello was in the last line with defer keyword therefore it gets printed firstly.

	All this work in reverse order which means last code execution will be treated as first code
	and will be executed at first.*/

}

func myDefer() {

	for i := 0; i < 10; i++ {

		defer fmt.Println(i)

		/*When we will run this we will see that the values which gonna be print till 9 are printed in
		reverse order just because of defer keyword.*/

	}

}

//Stay tuned for more updates
