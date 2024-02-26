/*In this session wew gonna learn that how can we handle URLs and Construct URLs and many more things
related to it*/

package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://pro.learner:3000/learn?coursename=solidity&paymentid=orejf4589" //this is not a real url this is just a sample url which is made by us , can say this is a fictitious url which we are creating as if you have seen any website thats how a basic url seen with domain , directories , queries and many more things.

func main() {

	fmt.Println("In this session we gonna learn about handling urls in go")

	fmt.Println(myURL)

	result, err := url.Parse(myURL) // with the help of url module and its Parse Property we can get full information regarding any URL , like here we are passing the url which we made above.

	if err != nil { //checking if there is any error

		panic(err)

	}

	// fmt.Println(result)

	//These are the properties which we can use to get the information about the URL.

	/*Lets consider the url we made above and apply those properties on it , as we have stored all the
	inforation of our URL inside result variable.*/

	// fmt.Println(result.Scheme) //https
	// fmt.Println(result.Host) //pro.learner:3000
	// fmt.Println(result.Path) ///learn
	// fmt.Println(result.Port()) //3000 , everything else is a property except port which is a method , so kindly remember this.
	// fmt.Println(result.RawQuery) //coursename=solidity&paymentid=orejf4589

	//So all these properties and method will give us the desired results.

	queryParams := result.Query() //we are storing this property in a variable so that can we can use it and apply more methods on it. like performing loop.

	fmt.Printf("The type of the query params is %T \n", queryParams)

	fmt.Println("result.query is giving us ", queryParams) //this will give us data in formatted way as comapre to .RawQuery property.

	fmt.Println("result.RawQuery is giving us ", result.RawQuery) //this will give us query data in raw form as its name says.

	// fmt.Println(queryParams.Values) //this will give us all the values which we are getting in queries also known as params.

	fmt.Println(queryParams["coursename"]) //in this way we are providing the key to get the value , as the values we are getting are in key value pairs.

	/*Here we are applying for loop with the help of range to get all the key and value we are getting
	as params.*/

	for key, value := range queryParams {

		fmt.Println(key, value)

	}

	/*we can construct a url in this way if we want to , this really helps in constructing url
	when we are handling our backend.

	Adn remember to use these keys Names similarly as i have given and as value you can give anything

	Here we are using url module and its URL property to construct the url
	And most important thing here we have to use pointers as we have send exact data and not the copy of it.*/

	partsOfURL := &url.URL{

		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=yash",
	}

	mainURL := partsOfURL.String() //converting the url into string.

	fmt.Println(mainURL)

}


// So This was all about this session.

// Stay Tuned For More Updates.