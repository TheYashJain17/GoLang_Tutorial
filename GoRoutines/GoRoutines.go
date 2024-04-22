// In this session we gonna Learn that what Are GoRoutines And What is the benefit of using it.

/*
So first lets understand what is GoRoutines and why do we use GoRoutines.

So basically we use goroutines to send threats to get a particular job done.

For example if i want to get the statuscode of a website , so i will just create a GoRoutine
and will send it to do the job.

Now the question arises that can it be done without the goroutines the answer is yes , so why are
we using the GoRoutines , the answer is this will help us to get the job done fastly.

*/

package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
Creating A Global Variable of type sync.WaitGroup as we are storing this in a variable wg
so that we can use this at multiple places.

WaitGroup is the functionality which helps us to create a GoRoutine and do the job for us
This exist inside sync library
*/

var wg sync.WaitGroup

func main() {

	// We are creating a slice of type string which includes the websites.

	websiteList := []string{

		"https://lco.dev",
		"https://google.com",

		"https://fb.com",
		"https://github.com",
	}

	/*

		Now we are running a for loop on the the slice we created above to get all the details we are using
		range keyword for this , this will return 2 things first is index which we are ignoring in this
		case therefore put the underscore and second is the website.

		Now inside the for loop we are using wg global variable which consists of our sync.WaitGroup
		and we are using .Add method of it for creating a new GoRoutine , so that the created GoRoutine
		can do the job for us which is to get the statusCode Of given website , always remember to
		create a GoRoutine first then only assign a task to it.

		So now how can we assign a particular GoRoutine a particular means that a GoRoutine will know
		that it has to do , so for this go keyword comes into the picture , with the help of
		go keyword we are sending that work and so that the GoRoutine will get to know that it has to do
		this particular work , thats how this flow works.

		In this case we are calling the function we created below and inside that website will be provided
		one by one as it running inside a loop and go will send again and again the task and so that
		the GoRoutine will be created again and again and this will be done till the loop gets ended.

		and atlast we have provided wg.Wait() which is another property of WaitGroup() which tells the
		program to wait until all the GoRoutines complete their task , this stops the main function to get
		executed fully , therefore we usually put .Wait() at the end of main function.

	*/

	for _, web := range websiteList {

		wg.Add(1)

		go getStatusCode(web)

	}

	wg.Wait()

}

/*

we are making a funtion to get the status code of a website we will provide inside this function
as parameter.

we are using defer keyword and writing wg.Done() , so this line means that the task given has been
completed so kindly delete the GoRoutine which has been created , it is important to do so and using
defer keyword so that this will gonn execute at the end of the code as we have already read about
the defer keyword.

Now we are using http library and its Get method to get the information of the website we are providing
inside it , now this will give us 2 things one is response and the other is err

so we are checking for the error as we usually do and after that we are printing the response
we are getting from http.Get.

we are using %d which is used to print the integer value , we are using this as we are getting the
statusCode which is an integer and second thing is %s which is used to print the string which is
website name in this case.

*/

func getStatusCode(website string) {

	defer wg.Done()

	res, err := http.Get(website)

	if err != nil {

		panic(err)

	}

	fmt.Printf("%d is the StatusCode of Website %s\n", res.StatusCode, website)

}

/*

So basically the flow of this is , .Add() creates a GoRoutine then after the completion of the work
assigned to GoRoutine .Done() deletes that task and .Wait() helps all the GoRoutines to get their work done
and only after the completion of work of every GoRoutine main function gets executed/end properly.

*/

// stay tuned for more updates
