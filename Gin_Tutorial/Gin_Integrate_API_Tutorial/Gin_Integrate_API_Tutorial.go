/*

In this session we gonna see how we can integrate 3rd party APIs into go with gin framework and get
the data.

*/

/*

so Basically we are going to take a 3rd party url which we know returns a json Data(it is important to know that what type of data we are
getting) and we are going to create our API and whenever the user will hit our API he/she will see the data what we are getting through
the 3rd party url.

*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This is the url for dummy json data as it name says

var url = "https://dummyjson.com/products/1"

/*

This is our main function and we are not doing much things inside it  , first of all initialising the gin router (hopefully till now
you must have known that how do we do it).

After that we are creating a route which whenever will get run this handler named gettingData.

And atlast running our API on the port 666 with the help of .Run.

*/

func main() {

	fmt.Println("GoodBye World")

	r := gin.Default()

	r.GET("/getData", gettingData)

	r.Run(":666")

}

/*

Now here comes the function inside which we are doing everything.

So first if we want to make this a handler then we have to use gin context as we have learnt  , and now inside it we are using
http package and its Get method to get the data from the url we are using that dummyjson url now this will return 2 things
The response and the other thing  is error , we are checking for the error if it is not equal to nil send the response to the user
with the help of gin.Context and if there is no error then we will proceed further.

So first of all we are closing the resp.Body but it should be close at the very end therefore using defer so that when all the work
has been done resp.Body should be closed it is important to do so because if we dont do it we can suffer lose of data as well as
security issues with the data.

Now after this we have to read the data present inside the resp we are getting and for that we are using io package because ioutill
has been depreciated , so therefore using io and its property .ReadAll , this again return 2 things first the data which is array of bytes
([]bytes) like this and another thing is the error so we are storing both of them and checking for the error , if it is not
equal to the nil then simply send a response to the user by showing the error and then we are proceeding further.

we are creating a mapping in which we are mapping string with interface{} we are using interface and not the string or int because
there is a possibility that the data we would be getting contain multitple types of data like string , int etc so the interface will
take care of it and storing this mapping inside a variable.

after that we are using json library and its Unmarshal property to convert the data which is of array of bytes ([]bytes) into the json
data and anoher parameter we have to provide the refrence of the mapping , like in most cases we make a mapping to convert the
bytes data into human readbale form therefore mapping is needed and make sure to provide the reference otherwise you will get an error.

This only returns an error so we are checking that like we did above and doing same thing if the error is not equal to nil.

And if everything is ok then sending the success response inside which we are sending the variable inside which we stored the mapping
because now the converted data exist in the mapping we created which is format in my case.

NOTE :- Dont do the mistake of sending the data variable thinking that , data variable will do the job because then you will get the data
in the bytes , i was doing this mistake thats why telling you make sure to take care of this.

And thats it noe you can render the data of a 3rd party API through your  API.

*/

func gettingData(c *gin.Context) {

	resp, err := http.Get(url)

	if err != nil {

		c.JSON(404, gin.H{

			"Message": err,
		})

	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {

		c.JSON(404, gin.H{

			"Message": err,
		})

	}

	var format map[string]interface{}

	err = json.Unmarshal(data, &format)

	if err != nil {

		c.JSON(404, gin.H{

			"Message": err,
		})

	}

	c.JSON(200, gin.H{

		"Success": format,
	})

}

// This was all for this session.

// Stay Tuned for more updates.
