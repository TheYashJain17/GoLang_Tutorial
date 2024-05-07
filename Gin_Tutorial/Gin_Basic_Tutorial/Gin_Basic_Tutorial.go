// In this session we gonna learn about the Gin framework , this is the basic working of Gin.

// We gonna see how can we make a basic API from Gin framework.

// For this we have to install gin , for that run this command :- go get -u github.com/gin-gonic/gin

// And now you would be able to use the gin framework.

package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

// So lets start the code in the main function as usual.

/*

So first of all we are intialising the router with the help of gin.Default() this will initialise it
and we are storing it inside a variable r to access it later.

Then On that variable we are using .GET method because we are making a GET  request and inside that we are providing the route
and then calling the handler which will be acitivated whenever the user will hit this route.

And after that we are running the router with the help of .Run , this will do the job of listening on the port , we are listening
this router on port :666.


You must have also seen the making of API with just gorilla mux and http inbuilt package , so in comparison of that this is much easier
We could have also used the chi router  but as we are learning Gin framework so this is fine.

And thats it now run the command :- go run Gin_Basic_tutorial.go (according to my filename) and then go to the web browser
and type localhost:666/devil then you will see the data we are sending with our handler function which is gettingData in this case.

And thats it.

*/

func main() {

	fmt.Println("Goodbye World")

	r := gin.Default()

	r.GET("/devil", gettingData)

	r.Run(":666")

}

// Now lets see what our handler function is doing.

/*

So for making a handler function for the API we have to use gin.context , as it is important to use gin.Context when we are using ,
becuase then only router will identify this as handler Function and this gin.Context is really important and have so many functionalities
in it , which we will get to know as we will build more APIs.

Now inside this we are using the variable inside which we were storing the  gin.Context and using it we are sending a JSON data
inside which we are sending the response status which we are accessing with the help of http package and sending StatusOK and then
sending the message in JSON form with the help of gin.H this makes mapping , dont worry about anything else just keep coding with me
You will understand it slowly slowly and as it accepts an object we are providing the data in the JSON form and this is the data
which will be shown when we will access the localhost:666/devil.

*/

func gettingData(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{

		"message": "Welcome To Hell",
	})

}

/*

Thats it , thats easily how you can make an API with the help of Gin framework and if you have used the gorilla mux then you will
understand that this is easy in comparison of gorilla mux.

*/

// Stay Tuned For More Updates.

//Keep Learning Keep Building.
