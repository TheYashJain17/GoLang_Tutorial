//Continuing Our API Development Series.

// In this session we gonna see how can we send form data with the help of go.

// It is important to know that how we can send form data thats why we are learning it.

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	fmt.Println("In this session we gonna see how we can send form data in go")

	PerformSendFormData() //calling the function in main so that we can run it.

}

func PerformSendFormData() {

	const myURL = "http://localhost:8000/postForm" //this is the url which we gonna use in backend , you can take any url but remember to keep this and backend url.

	data := url.Values{} //we gonna use url module and its Values property to make a form therefore using square brackets.Storing this in a variable so that can use add property on it to add fields and data in it to send.

	//This is the sample data we are sending , Adding Fields an asnfand its data , you can take any field and any sample data and we are adding all this in the form we made with the help of .Add property.

	data.Add("FirstName", "Yash")
	data.Add("LastName", "Jain")
	data.Add("Email", "yash@gmail.com")

	response, err := http.PostForm(myURL, data) //using http module and its PostForm property , we could have use Post property but PostForm is easy to use as compare to Post in sending form data. Inside it providing the url and the form we created above. Storing in 2 variables as it can give us resposne as well as an err.

	if err != nil { //checking for the error.

		panic(err)

	}

	content, err := io.ReadAll(response.Body) //using io modue and it property ReadAll to read the body of the response we are getting and again storing this in 2 variables as we can get data as well as the error.

	defer response.Body.Close() //its important to close the connection with the response therefore doing so with the help of .Close() function and using defer keyword so that it will gonna execute at the last.

	if err != nil { //checking for the error

		panic(err)

	}

	fmt.Println(string(content)) //converting the content into string so that we can read it as we know we would be it getting in bytes form.

}

//This was all for this session 

// Stay Tuned for more updates
