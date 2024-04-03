// In this session we gonna learn about channels , we gonna see the basics of channels in this session.

// So first lets understand that what are channels in go.

// If i have to tell this in single statement then it would be that a channel is a medium of communication.

// We needs channels Whenever we are dealing with multiple go routines , we will discuss about go routines in later sessions.

/*
As we read that channel is a medium of communication or can say that medium of medium of getting information , so basically
channels helps us to get the required information from a specific go routine.
*/

package main

import "fmt"

// Now lets understand that how can we create channel in go.
func main() {

	// So there are 2 ways of creating a channel:-

	// var numberChannel chan int - This is the first of declaring a channel. we just have to use var then the channel name then chan keyword which helps us to create the channel and then the type of channel like integer in this case thats why we are using int.

	// numberChannel := make(chan int) - This is the second way of declaring a channel. we are using shorthand method in this case and then using make keyword to create a channel inside that we are using chan keyword which helps us to create a channel and then the type of the channel whichh is integer in this case also.

	/*

		You can use whatever you want to use , as if you have read my previous notes you must know that we can use the first approach on
		global level as well as on local level , whereas we can use the 2nd approach only on local level.


	*/

	numberChannel := make(chan int)

	fmt.Println("The value of the channel is ", numberChannel)

	fmt.Printf("the type of the channel is %T \n", numberChannel)

	/*

		So in the above case we are creating a channel with the help of short hand method and it is of type int.

		this is an empty channel still we are checking its value , when we will print this out we will get an address value.

		And when will print the type of channel we will get the type of the channel which is of int , we are using %T to get the type of the
		channel.

	*/

	/*

		 Now we will see that how can we send a value inside a channel.

		 we are going to use the channel we created above.


		And make sure you are sending the value of same type of whose you have made the channel , in this case it would be of int type ,

		So we will send an integer type of value in our channel.

	*/

	// numberChannel <- 278

	/*

		so this is the syntax for sending a value into the channel.

		So basically we use a left side slim arrow to send the value inside the channel.

	*/

	/*

		Now we will see that how can we recieve this value inside another variable to print this and if we want we can also directly
		print it also.

	*/

	// numberChannelValue := <- numberChannel;

	// fmt.Println(numberChannelValue);

	// <- numberChannel

	/*

		so these are the ways of printing the values of a channe.

		In first case we are storing the every value present inside the channel into a variable with the help of short hand method
		And then the left side slim arrow and then the channel name.
		Like this is the difference between sending the value inside a channel and sending the whole channel values inside a variable.

		And then we are printing that.

		In second case we are not storing this insie a variable we are just directy using the left side slim arrow and then the channel name.

	*/

	// so now lets see another working of this by creating a new channel and a separate function.

	testingChannel := make(chan int)

	fmt.Println("GoodBye World")

	go multiplychan(testingChannel)

	testingChannel <- 89
	fmt.Println("Another GoodBye")

	/*

		So lets see what we are doing here:-

		So first we are creating another channel for this purpose.

		And then we are printing a statement to understand the flow of it.

		then we are using the go keyword and then caliing the function and inside it passing the channel we created above ,
		why we are doing this you will get to know when we will learn about go routines thats why not digging this much.

		After that we are sending a value inside the channel so that channel coulnt be empty , you might be thinking that why we
		are sending the value inside the channel after calling the function and not before that , this is because if we will do so
		we will get a deadlock error , therefore we are sending the value inside the channel later because we are using go keyword
		it will send the function but will not wait for its return and will execute the next line something like javascript we use await
		in that case , so the value will be set and then the channel will contain the value and not when we send the value before calling
		the function , i know this sounds crazy but this is what it is.

		And then atlast again printing the statement to check the flow.

		And thats it

	*/

}

/*
so this is the function we are using above , this is accepting a channel of type of int as a parameter and inside it
we are just simply printing the result , we are multiplying the value we are geting from the entered channel we are multiplying it
with 100 , for getting the value we are using left side slim arrow , we learnt this above and directly multiplying it with 100.
*/

func multiplychan(_chan chan int) {

	fmt.Println(100 * <-_chan)

}

// So this was all for the channels , we will learn more about channels after studying the go routines.

// so thats it for this session.

// Stay tuned for more updates.
