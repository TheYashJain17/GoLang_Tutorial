/*So in this session we gonna see that how  we can connect with mongodb with the help of gin in Go , this process will gonna remain same
for gorilla mux as well as chi router , but as we are learning Gin thats why i am covering this under gin tutorial.*/

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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

	err := godotenv.Load()
	if err != nil {

		fmt.Println("Please Create Dot ENV File")

	}

	mongo_url := os.Getenv("MONGO_URL")

	if mongo_url == "" {

		fmt.Println("Please Provide mongodb url")

	}

	clientOptions := options.Client().ApplyURI(mongo_url)

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

	fmt.Println("MONGODB CONNECTED") //we completed our last session till here.

	// Now starting this session from here.

	/*

		So basically we estblished a connection with mongo in our previous session , but the thing we werent getting anything from as we are just
		simply printing out that MONGODB CONNECTED!...

		So now lets level up the game little bit and read the data from the database to confirm that we are connected with our database.

		So lets first do few things to ensure that we will get the data.

		First obviously create a database , if you havent already then make a collection inside it like i have defined in the code and then
		inside that collection make a sample document which we can read here , so do all these things first then come back here.

		So Now Lets begin.

		so first we are getting the collection associated with the database whose link we have provided in dotenv file , if you didnt understand
		this check out my previous session then you will understand it.

		so for getting the collection we are using the client we made above and then providing the exact name of database then the exact name of
		collection we made as i told you.

		then we are creating a variable of bson.D type , we have taken this type because mongodb supports bson data and we are using bson.D
		because we want to get data in structured manner , there are mainly 2 types in bson :- bson.D which provides the structured data and
		another is bson.M which provides the unstructured data , so we have made a variable of bson.D type.

		Now we are using the collection we stored above and on that we are running the FindOne operation of mongodb and this takes 2 parameters
		first is the context we made above and second is an interface which should be of bson type we can take any type like .M or .D , i have taken
		the oppposite and we have taken an interface because data present inside the collection could be of many types like string int bool so
		in this type of case interface is best and now we will use .Decode to decode the data we are getting in the format we want to and for this
		i have defined the variable above which is of bson.D this will return the data in structured form , this will gonna return me an error
		so i am checking that if the error is not equal to nil then check for the type of error like the mongo library we installed it also gives
		the possible types of errors like here in our case we are checking if the error == mongo.ErrNoDocuments which means there is no document in
		the collection we have made then we are telling the user the same thing else there is a possibility that the user hasnt made the collection.

		and if everything is sorted we are just printing out the data we got in the decoded form to see it properly.

		Thats how we can check for a proper connection with the mongodb.


	*/

	collection := client.Database("GO-API").Collection("test-col")

	var format bson.D

	err = collection.FindOne(ctxt, bson.M{}).Decode(&format)

	if err != nil {

		if err == mongo.ErrNoDocuments {

			fmt.Println("Connected But Couldnt Read Collection")

		} else {

			fmt.Println("Collection doesnt exist")

		}

	} else {

		fmt.Println("Success", format)

	}

}

// Thats it for this session.

// Stay Tuned For More Updates.
