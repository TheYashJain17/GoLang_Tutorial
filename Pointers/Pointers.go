//In today's session we gonna learn about the Pointers in go which is very important topic.

// Lets understand the concept of pointers and its use.

// Lets understand the use of pointers.

/*Whenever we store some data into any variable and then pass it into any function or any other thing
Sometimes instead of that original data , a copy of that data gets passed on and that causes some issues
in our code.

So here comes the role of pointers as pointers saves us from this problem ,
Instead of passing the data , pointers pass the memory address of that particular data we stored inside
variable so this gives a surety that only the original data will be passed on and no copy will be passed.*/

// In short a pointer is nothing but a direct location of the memory address of the thing we stored.

package main

import "fmt"

func main() {

	fmt.Println("Lets understand pointers")

	// we have 2 ways of telling that a particular thing is pointer or not.

	// The first way is with an asterix or can say multiply sign (*).

	var ptr *int

	fmt.Println("Value of pointer is ", ptr) //we will get nil as a result , as we have not stored anything inside it yet.

	/*Now lets see 2nd way is with mpercent (&). This is the way of referring ,
	Whenever we want to refer to any variable use mpercent(&).*/

	myPtr := 23

	var anotherPtr = &myPtr

	fmt.Println("Memory Address of the pointer is ", anotherPtr)

	/*When we will run the above line , we will not get the value of the pointer
	instead we will get the memory address of that pointer.*/

	// if we want to get the value of pointer we made , write this:-

	fmt.Println("Value of the pointer is ", *anotherPtr)

	/*So with the above line we can see that whenever we want to get the value of any pointer
	use asterix/multiply sign with the pointer name , then only we can get the value of
	pointer otherwise we will get the memory address of that pointer.*/

	// Now if we want to update/change the value of the pointer , we can do it like this :-

	*&myPtr = *&myPtr * 2

	*anotherPtr = *anotherPtr * 2

	/*We can use both the ways to update the value of pointer , little tricky , but we will
	understand this better soon as we will work on it.*/

	fmt.Println("Updated Value of the pointer is ", myPtr)

	fmt.Println("Updated Value of the pointer is ", *anotherPtr)

	/*So inshort we can say that with with asterix we can target any pointer and change its value
	like we did above .

	In both the ways we will get the updated value of the pointer.*/

}

//So this was all about pointers , Hope you wont have any issues.

//Stay Tuned for more updates.
