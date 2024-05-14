// In this session we are going to see how can we make basic REST API with the help of Gin framework.

/*

In this session we only gonna see GET , POST And DELETE request not including PUT because it is little difficult for now , so we will
see it later.

*/

// This is on a very basic level , we will see it on an advanced level in later sessions.

// For this first install Gin Framework for that write go  get -u github.com/gin-gonic/gin this will install the Gin framework.

package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

/*

First of all we are creating a struct of type User and inside it taking 2 informations Name and Email and if you are taking my sessions
you must know that why does i have written the json thing and  if not then first take those sessions then come to this.

*/

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

/*

Now we are creating global variable userData which is of type User struct we created above , so it means the data will stored inside
this userData varaiable in the above struct format.

*/

var userData User

/*

Now inside the main function we are initialising the gin framework with  the help of gin.Default which we learnt about in previous
session.


After that creating mulitple requests with the help of router initiliasation we stored inside the variable r and now using that r to make
multiple requests , like GET for gettingt the user data , POST for adding the data user data and then DELETE for deleting the user data.

You must have known about this like providing the url and then the handler means that handler function will be executed whenever the
user will visit this route.

And atlast r.Run will help us to listen this application on a particular port.

If you dont know what i am talking about just pay a visit to my previous session whether on linkedin or on github.

Then you can understand this better.

*/

func main() {

	fmt.Println("GoodBye World")

	r := gin.Default()

	r.GET("/userData", gettingData)

	r.POST("/userData", addingData)

	r.DELETE("/userData", deletingData)

	r.Run(":666")

}

/*

This  is our first handler which is helping us getting the data , so for this to make a handler function for gin api we have to use
c *gin.Context this will tell the system that this is an handler function.

Now inside this function we are using the variable c inside which we stored the *gin.Context , now we are using .JSON property inside
it providing the status with the help of http package and then using gin.H which will generate the mapping through which we are sending
the object inside which we are returning the userData Which is the variable inside which we are storing all the user data ,so thats how
the user would be able to see the user data.

*/

func gettingData(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{

		"userdata": userData,
	})

}

/*

This our second handler which is helping us adding the data , so for this to make a handler function for gin api we have to use
c *gin.Context this will tell the system that this is an handler function.


so we are using a variable c inside which we are storing *gin.Context and now on it we are using .BindJSON property which lets us
bind the data as its name says in easy language adding the data which we are receiving from the body like in nodejs form it is req.body
into a variable or anything , so we have used  here & which means we are borrowing the variable this means take the same variable and
dont make another variable , so thats how we can add all the data at a single place and not at multiple places therefore it is important
to use & sign and there is a possibility of getting the error therefore we are taking care of that.

We are checking for the error , if it is not equal to nil then send the reponse as we have send it above , you must have learnt it till yet
if dont then see it again and try to understand it.

if there is no error we are just sending the response with the success message and the data user entered in the json form which will be
stored in the variable.

*/

func addingData(c *gin.Context) {

	err := c.BindJSON(&userData)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{

			"message": "Please Enter Valid Data",
		})

	}

	c.JSON(http.StatusOK, gin.H{

		"Success": userData,
	})

}

/*

This is our third handler which is helping us getting the data , so for this to make a handler function for gin api we have to use
c *gin.Context this will tell the system that this is an handler function.

So for deleting the data we are simply making the userData variable which is of type user equal to empty User object means User{} which
is our struct in simple language by doing this we are setting the value of userData to empty and then sending the response.

*/

func deletingData(c *gin.Context) {

	userData = User{}

	c.JSON(http.StatusOK, gin.H{

		"message": "User has been deleted Successfully",
	})

}

/*

So thats how we can make basic REST API with the help of Gin framework.

You can test this on postman or on thunderclient first run this file with go run then filename, then  just type localhost:666/userData
this is according to my port and routes i have provided you can change it accordingly.

And thats it

Stay tuned for more updates.

*/
