// Continuing Our API Developement Series In Go.

//In this session we gonna see how can we make post request in go.

package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("In this session we gonna learn about post request in go")

	PerformPostRequest()

}

//Making a different function to perform the post request.

func PerformPostRequest() {

	const myURL = "http://localhost:8000/post" //this is the url which we gonna use in backend , you can take any url but remember to keep this and backend url.

	//fake json payload or data

	/*Here we are using strings module and its New Reader property to send the data in json form
	we gonna read about json in details in later sessions.

	Inside NewReader we are using back ticks and inside that using square brackets to send json data.*/

	requestBody := strings.NewReader(` 
	
		{

			"CourseName" : "Go Lang",

			"Price" : 0,

			"Platform" : "yash jain"

		}

	`)

	response, err := http.Post(myURL, "application/json", requestBody) //now we are using http module and its Post property , inside that providing the url which we made abov above and which is common in our backend also , then setting the header if you know little bit about backend then you must knwo about applicatiob/json meaning , it means we are sending the data in json form , therefore we have written it and atlast the json data we created above.

	if err != nil {

		panic(err)

	}

	defer response.Body.Close() //it is important to close the response at last therefore doing so with the help of reponse.Body.Close mehtod and using defer method so that this will gonna execute at the last.

	content, err := io.ReadAll(response.Body) //using io module and its property ReadAll property to read the response body.

	if err != nil {

		panic(err)

	}

	fmt.Println(string(content)) //converting the data into string which we are getting in bytes so that we can read it.

}


//Stay Tuned For More Updates
