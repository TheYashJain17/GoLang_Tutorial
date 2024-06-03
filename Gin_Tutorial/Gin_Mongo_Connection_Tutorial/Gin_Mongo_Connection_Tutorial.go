/*So in this session we gonna see that how  we can connect with mongodb with the help of gin in Go , this process will gonna remain same
for gorilla mux as well as chi router , but as we are learning Gin thats why i am covering this under gin tutorial.*/

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"         //install this package as we need this package to make a connection with mongo , just type go get -u go.mongodb.org/mongo-driver/mongo
	"go.mongodb.org/mongo-driver/mongo/options" //this package will gonna install with the above package.
)

func main() {

	fmt.Println("Goodbye World")

	connectionWithMongo() //calling this function inside main function so that it can get executed.

}

// So we made a function and inside this function we gonna make the connectio with our mongodb.

/*

See this is the syntax you have to remember as we are not writing any logic here , this is the exact syntax you have to follow to make
the connection , there are other ways also to do the same , but this seems little easy to me thats why telling you.

So you have to get familiar with the syntax and here i will explain you this syntax , so you can understand the working better.

So we have made a variable clientOptions and inside that we are using options library which comes with mongo library and with that we are creating
a client and on that we are applying our mongodb url (make sure to get one from the mongodb , i have also used this is blank here because i have
removed mine but you have to use it else you will get an error).

after that in the next line we have to create a context , so we will use context library which returns 2 things one is the context and another
one is cancel function this will be invoked in case when the system wont be able to establish the connection with the mongodb.
So Now we are using WithTimeout functionality from context and inside that providing context.Background which will make that the connection is
running behind the scene and other thing is time which are 10 seconds in this case using time library for this.

So basically we are creating a context with timeout functionality which will make sure that the context is running in background and if the
connection time exceeds 10 seconds(because we have mentioned 10 seconds) then the cancel function will be activated else we will get
the context.
and we are using defer with cancel() so that when it is checked properly then only the cancel function should get called otherwise it may
get called midway which is not good.

Now we have to make a connection with the mongodb and we are going to need the ctxt and clientOptions we have created in it and this will return
2 things a client and an error and for this we are using mongo.Connect function this will give us the client if there is not any error.
Then we are checking for the error just like normally we do.

then at last we are remained with one thing and that is to ping the client like to check whether we are connected wth the mongodb or not.

this returns an error so we are just checking it and if there is no error we will simply get MONGODB CONNECTED which we are printing in next line.


And thats how we can get to know that whether we are connected with the mongodb successfully or not.

Congrats you have learnt to connect with the mongodb with the help of go.

*/

func connectionWithMongo() {

	clientOptions := options.Client().ApplyURI("mongodb+srv://Yashjain:thisismypasswordyouasshole666@cluster0.r9okdwn.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	ctxt, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	client, err := mongo.Connect(ctxt, clientOptions)

	if err != nil {

		log.Fatalf("The error is %+v", err)

	}

	err = client.Ping(ctxt, nil)

	if err != nil {

		log.Fatalf("The error is %+v", err)

	}

	fmt.Println("MONGODB CONNECTED")

}

// Thats it for this session.

// Stay Tuned For More Updates.
