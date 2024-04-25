// In this session we gonna learn about Mutex which is another important topic in Go.

// So basically mutex plays a major role when it comes to sending goroutines to do a same job.

// You will understand this with below exmaple.

package main

import (
	"fmt"
	"sync"
)

var sharedResources int //created a global variable of type int so that can access it from anywhere.

var mutex sync.Mutex //importing mutex with sync package.

/*

Inside main function we are declaring the value of the global variable we made above.

And then we are running a for loop hopefully till now you have understood about the for loop
and if not then check my previous sessions.

So we are running this for loop 500 times and why we have taken this much big number you will understand
in sometime.

And inside this loop we are sending the go routines to another function to perform a specific task
Whihc is nothing but to increment the value of the global variable by one , this means that 500 goroutines
will go to that function and increment the global variable by 500 and after the loop we are printing the value
of the global variable so it should be 500 , try to run it and see the ans , you will see thats not the case.


Lets understand whats happening , here we are sending 500 goroutines and telling them to increment the value
according to our function we are calling after go keyword , so when they will reach there , they will all try to
increment the value of that variable in a single go so suppose there are 2 goroutines when they reach there they
see that value for example at that time the value of that variable is 5 and now they both will see and will decide
that they have to make it to 6 in that case one goroutine will make it and another one will see as job done and thats
why that go routine will not increment the value to 7 and hence it will happen again and again and thats why  we are not getting
500 , always we are getting the less value , now you must be thinking that why i took such a big number so thats because
if you will try this with a short value say 50 then you will get the exact ans and you will not understand the value of mutex.

so what mutex does in simple language is it lets to get in a single variable at a time , means those 500 goroutines will enter
turn by turn instead of entering in a single time , obviously this will take much time as compare to before but this will give
us the exact result every single time.

I know this will lead to the situation of concurrency but for getting the exact same result it is important for us because parallelism
will lead to wrong result.

So thats why the role of mutex is crucial.

*/

func main() {

	fmt.Println("GoodBye World")

	sharedResources = 0

	for i := 1; i <= 500; i++ {

		go incSharedResources()

	}

	fmt.Println(sharedResources)

}

/*

So this the function we are creating to increment the value of that global variable.

Now inside it we are using mutex.lock and mutex.Unlock , let understand there use cases.

So mutex.Lock means whenever a goroutine will enter into the function mutex will lock the global variable we are incrementing here

And will keep it lock until the gorotuine has completed it job which is to increment the value in this case and therefore we have to
write this is in the very begining of the function.

So now lets understand the mutex.Unlock , so when the job of goroutine is done it is important for us to unlock the value which we
locking in starting global variable in this case so that another goroutine can come and perform the job on it , as until the value
is locked no other goroutine can enter into the function , so it is important to unlock and we do it at the very end , therefore
we are using defer keyword hopefully you know this much else go back and clear your basics first.


So thats how the whole procedure is going on.

*/

func incSharedResources() {

	mutex.Lock()

	defer mutex.Unlock()

	sharedResources++

}

//Stay tuned For More upates
