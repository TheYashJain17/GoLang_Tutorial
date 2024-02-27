/*From This session we gonna start the API Development with the help of Go.

So from now i am starting the API learning in this and this needs a backend server as 
for now we are just making the requests and not creating any server , we will learn that in later 
sessions . so if you want you can create a backend server with the help of node js , but for this 
course i am not making this complicated , i am just  keeping this simple , so just making the APIs
and not creating any server.

*/

// So Lets Begin Our API Development journey in Go.

// In this session we gonna learn about how can we make get request in go.

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("In this session we gonna learn how to make get request in go ")

	PerformGetRequest() //calling the function we made below so that we can run it , as we know that if we want to run any function it is important to call that fuction in main function

}

// Making a different function to perform go request

func PerformGetRequest() {

	const myURL = "http://localhost:8000/get" //this is the url which we gonna use in backend , you can take any url but remember to keep this and backend url.

	response, err := http.Get(myURL) //as we have already studied like how to get the url with the help of http module and its Get property so with the help of that we are fetching the information.

	if err != nil {

		panic(err)

	}

	defer response.Body.Close() //it is important to close the connection therefore we are doing it and we are using defer keyword so that it will gonna execute at the last.

	fmt.Println("The status code of the response is: ", response.StatusCode) //printing the status code as it is important.

	fmt.Println("The Length  of the response is: ", response.ContentLength) //this is another property which we can use.

	dataInBytes, err := io.ReadAll(response.Body) //reading the body of the response with io module and its property of ReadAll  and as we know we get data in bytes and there is a chance of getting error also therefore checking for both.

	if err != nil {

		panic(err)

	}

	content := string(dataInBytes) //converting the data which is in bytes into string.

	fmt.Println(content) // atlast printing it to see the response.

}
