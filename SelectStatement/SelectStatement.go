// In this session we gonna learn about the Select statement in Go.

// You can understand Select Statement as Switch Case Of Go.

// But there is a huge difference in working of both of them.

// Lets See

package main

import (
	"fmt"
	"time"
)

/*

So now lets understand everything about select Statement , so basically as i said that you can understand select statement
as switch case which we have learnt about in many languages like in javascript , but there is a huge difference in working
of both of them.
switch case prints the result according to the desired case we gets , but in case of select statement we gets the result according
to the case we first get.

Lets understand this with the help of an example.

*/

/*

So in the main function i have made 2 channels which are of string type , hopefully till now you may have understood that how one
can make channels in go and if you still dont know then check out my previous posts you will know about channels.

Anyways , after this i have made 2 functions which are just putting the value inside the channels i have created.

then i am calling those functions with go keyword , if you call them without go then you will get a deadlock error.


after calling them , here comes the main syntax which we all are waiting for.

for creating select statement you have to use select keyword then use curly braces and now inside it , we will declare cases
just like switch cases , now we are using case keyword to define a case and then we are storing the value of chan1 inside a
variable then using colon (:) to say if there is a value in the variable means in the chan1 as we are getting the value from chan1
so just print that value.

Similarly defining second case , doing everything same but for the chan2 to get the value from chan2 then if there is a value in chan2
which we stored inside the variable then just print it.



And also there is default case just like switch case and we are just adding random thing inside it to test it , if other cases
failed to perform the task then the default case will do the job.

*/

/*

Now lets see the working of the working of select statement , there is no rocket science in this case , its just
first come first serve scene , means whichever the function you are calling first that function value will gonna reach inside the
variable first and that case will be printed , yeah yeah this is as simple as that . this case only occurs when both the functions
are taking same time in the execution , for example we can increade the time of the execution of the function with the help of
time.Sleep which i have used in below functions , you just have to use them as it is and the execution of the function will be delayed.

but if you are using this time.Sleep in both the functions then also the result will gonna remain same , whatever the function we are
calling  accordingly that case will be printed and if you talk about the default case , the default case will be printed
when both the functions gets failed to execute for some reason then default case will be printed , you can also check this by commenting
out both the functions. And thats it that was about the selet statement in go , there is more use of this in further sessions.


NOTE :- If you are using an old system just like me then there could be changes in the result because system will slow down the
speed of the execution of the function and that will cost the change in the results.

*/

func main() {

	chan1 := make(chan string)

	chan2 := make(chan string)

	go putValueInChan1(chan1)

	go putValueInChan2(chan2)

	select {

	case chan1Value := <-chan1:

		fmt.Println(chan1Value)

	case chan2Value := <-chan2:

		fmt.Println(chan2Value)

	default:

		fmt.Println("GoodBye From Default Channel")

	}

}

/*

These are the functions i have made to put the value in both the channels , they are just simply taking a parameter which is of string
channel type and then inside the function we are entering a value inside the channel we took as parameter.

Hopefully til now you would have understood thats how do we enter a value inside the channel and if you still havent check out my previous
posts that will help you in learning that.

*/

func putValueInChan1(_chan chan string) {

	time.Sleep(1 * time.Second)

	_chan <- "GoodBye From Chan1"

}

func putValueInChan2(_chan chan string) {

	time.Sleep(1 * time.Second)

	_chan <- "GoodBye From Chan2"

}

// Stay tuned for more updates.
