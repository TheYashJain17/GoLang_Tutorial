/*
Now we are starting another important topic in go which is Concurrency as this will lead to many
other important topics.

So In this session we gonna learn about the Concurrency And Paralleism in go.

*/

package main

func main() {

	// This is a really very important topic so pay attention to it.

	// So in short if i have to explain the concurreny and paraelleism to you , it would be in this way:-

	/*

										Concurrency In Go

		Concurrency is the case where a task is dependent upon another task, it means the task has to wait
		for another task to get completed so then only the other task will be exexuted.


		Suppose there are 2 tasks :- TaskA And TaskB , so In this Case Task B Is dependent upon Task A
		For Its Completion , So first Task A Will Complete And Then Only Task B Can be completed.

		So if we execute 2 tasks in go then they will be executed one by one in a way in which they
		would be dependent upon each other.

		This is the case of Concurrency.

	*/

	/*

										Paralleism In Go

		In Case of Paralleism , No task is dependent upon another task for its Completion.

		Suppose there are 2 tasks Task A And Task B And We execute both tasks , so they will be executed
		at the same time and will not depend upon another task for the completion.

		Like in case of Concurrency Where Task B Was dependent upon Task A for its completion.

		If we execute 2 Tasks in this case , they both will be executed together irrespective of execution
		of another task.

	*/

}

/*

This was all about Concurrency and Paralleism in go , in later sessions we will discuss about them
in detail.

Stay Tuned For more Updates.

*/
