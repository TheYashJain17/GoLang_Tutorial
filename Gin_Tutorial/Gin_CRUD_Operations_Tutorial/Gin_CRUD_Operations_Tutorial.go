// So in this tutorial we will gonna see that how can we perform crud operations with the help of gin framework and mongodb.

// For this tutorial install joho/godotenv and go.mongodb.org/mongo-driver/mongo.

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Making a struct with the name manager and inside maintaing the mongo client , context and cancelFunc , storing all this inside struct
so that we can use it whenever we want to.*/

type manager struct {
	connection *mongo.Client
	ctxt       context.Context
	cancel     context.CancelFunc
}

/*
Making an interface , this is similar to interface of solidity like here we are defining all the function which we are going to need just like
solidity its just like a blue print of the all the functions because there are all definition of functions and no logic is present there.

*/

type Manager interface {
	Insert(interface{}) error               //First function is Insert this will help us in creation of a new user , which will take an interface as a parameter and will return an error if there is any.
	GetAll() ([]User, error)                //Second function is GetAll this will give us all the users present inside the db , this will not take anything and it will return array of the user which is another struct and an error if there is any.
	DeleteData(id primitive.ObjectID) error //Third function is DeleteData which will help us in delete the user and this will take an id which is of type primitive.ObjectId make sure to give this type because mongodb supports this type and thsi will return an error if there is any.
	UpdateData(User) error                  //Fourth funtion is to update the exsisting user and this will take an user which is of type struct which will define below and this will return an error if any.
}

// Making a struct for the user , as you know struct work as schema in go like there is a schema in nodejs so understand struct as schema for go.

type User struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` //Taking id as a field , here defining json and bson both json is for us to enter the data and bson is for the mongodb to understand the data we are sending because mongo understands bson and not json therefore using bson also and also you can see that i have used omitempty this means it is not required true in language of node js , if this field is not entered by us then it will be disappread but will not give any error if we dont provide it , and it is important to mark this omitempty else mongo db wont be able to provide the id from its side , if you know that mongo provides a unique id to its every document so it is important to mark this omitempty.
	Name  string             `json:"name" bson:"name"`                   //Taking name of the user as another field and same json and bson thing.
	Email string             `json:"email" bson:"email"`                 //Taking email of the user as another field and same json and bson thing.
}

var Mgr Manager //storing the interface we created in a variable.

/*
Now lets come to our main function , i will suggest you to first check the below code then only you can understand what is happening in the
main function.
*/

func main() {

	fmt.Println("GoodBye World")

	//Here we are adding one user into our database , creating a data of user with the help of user struct

	user := User{Name: "Yash", Email: "Theyashjain17.com"}

	/*

		Now entering that data into the Insert method which is present inside Mgr therefore using it. This is returning an error so we are checking
		whethere there is an error or not.
		And thats how we can successfully add data into the database.

	*/

	err := Mgr.Insert(user)

	fmt.Println(err)

	//Here we are calling the GetAll method to get all the users this will return data and error so we are printing both of them.

	data, err := Mgr.GetAll()

	fmt.Println(data, err)

	//Here we are delelting the record of a particular user with the help of deletedata method.

	id := "" //here we will provide the id whose data we want to delete.

	objectId, err := primitive.ObjectIDFromHex(id) //then we are using the primitive package and its ObjectIDFromHex to convert the normal id into the object id because in case of mongo only the object id is supported , so this function will return object id and an error checking for error.

	if err != nil {

		fmt.Println(err)

		return

	}

	//Passing the object id we got from above into method deleteData which is present inside Mgr , this will return the data so checking it out.

	err = Mgr.DeleteData(objectId)

	fmt.Println(err)

	//Here we are Updating the  Data of the user.

	objectId, err = primitive.ObjectIDFromHex(id) //then we are using the primitive package and its ObjectIDFromHex to convert the normal id into the object id because in case of mongo only the object id is supported , so this function will return object id and an error checking for error.

	if err != nil {

		fmt.Println(err)

	}

	// providing the updated data below , this is the same user object we created at top at the time of inserting data into mongo so we are updating that else we would have to create a new.

	user.Id = objectId               //setting the user id to the same id because we dont want to change the id
	user.Name = "Lucifer"            //updating the name
	user.Email = "Lucifer@gmail.com" //updating the email

	//After calling the UpdateData method present inside the Mgr the data present at that specific id will be updated , this will give us an error so we will check it and it would be nil as always.

	err = Mgr.UpdateData(user)

	fmt.Println(err)

}

// Here we have used init function which helps us to run a particular function only a single time at the time when we will run the function.

func init() {

	connectMongo() //here we have defined our connectMongo function so that the connection with mongo should get established only single time at the time of starting run of the program.

}

// I have already describe you the connect db function in my previous notes so wont gonna repeat this thing here.

func connectMongo() {

	err := godotenv.Load()
	if err != nil {

		fmt.Println("Please Create Dot ENV File")

	}

	mongo_url := os.Getenv("MONGO_URL")

	if mongo_url == "" {

		fmt.Println("Please Provide Mongo URL")

	}

	clientOptions := options.Client().ApplyURI(mongo_url)

	ctxt, cancel := context.WithTimeout(context.TODO(), time.Second*10)

	client, err := mongo.Connect(ctxt, clientOptions)

	if err != nil {

		log.Fatalf("The Error is %+v", err)

	}

	err = client.Ping(ctxt, nil)

	if err != nil {

		log.Fatalf("The error is %+v", err)

	}

	/*

		Only thing which i will describe you in this is this below thing.

		We are storing all the information in the manager struct and that struct is getting stored inside the variable which is type Manager which
		an ineterface this will helps us keeping all things at a single place.

	*/

	Mgr = &manager{connection: client, ctxt: ctxt, cancel: cancel}

}

// This is a function to cancel the connection with mongo , this just simply takes cancel function as parameter and then call that cancel function.

func close(cancel context.CancelFunc) {

	defer cancel()

}

/*

Here we have started with our first method (hopefully you know the difference between funtion and method if not then check my previous sessisns)
which is insert taking manager as parameter first so that we can use fields which we have defined inside our struct and taking data whose type
is an interface and then returning an error.

Inside the method we are getting the collection with the help of struct and then using its .connection and then from that connection getting
the database and then from the database getting the collection.

Now the collection we have got , on that collection we are calling mongo function insertOne to insert data inside that collection inside this
providing context and data which is our interface , this will return a result or an error if there is any.

Then at last printing the result InsertedID this is provided by the mongo db which shows as success as this returns the id of the data
which has been added in to mongodb successfully.

*/

func (mgr *manager) Insert(data interface{}) error {

	orgCollection := mgr.connection.Database("GO-API").Collection("test-col")

	result, err := orgCollection.InsertOne(context.TODO(), data)

	fmt.Println(result.InsertedID)

	return err

}

/*

Here we have started with our second method which is GetAll taking manager as parameter first so that we can use fields which we have
defined inside our struct and then returning  array of the user which is the struct we defined above and an error.

Inside the method we are getting the collection with the help of struct and then using its .connection and then from that connection getting
the database and then from the database getting the collection.

then we created a variable which consist of options this is a package which comes with mongo package and using .Find() method which will work as
findAll method of mongodb.

then using the collection we got and on that using .Find() method again and inside it passing the todo then bson.D{} to get the data in a
structured manner you could also use bson.M{} but then you will get the data in non structured manner and then providing findOptions to get
all the data , this will return 2 things one is cur and another is err.

you can understand cur as encoded data so we have to decode it , so for that we are running for loop with .Next method and inside that providing
context and inside the loop we are taking a variable of type User and now we are decoding the data in the form of user struct so that we can
understand it fully.

then at last storing this decoded data into the variable data which is of type array of user then closing the cur as we close the resp body.

*/

func (mgr *manager) GetAll() (data []User, err error) {

	orgCollection := mgr.connection.Database("GO-API").Collection("test-col")

	findOptions := options.Find()

	cur, err := orgCollection.Find(context.TODO(), bson.D{}, findOptions)

	for cur.Next(context.TODO()) {

		var d User

		err := cur.Decode(&d)

		if err != nil {

			log.Fatal(err)

		}

		data = append(data, d)
	}

	if err := cur.Err(); err != nil {
		return nil, err

	}

	cur.Close(context.TODO())

	return data, nil

}

/*

Here we are with our third method which is Delete data taking manager as parameter first so that we can use fields which we have defined
inside our struct and taking id which is of type primitive.ObjectID and then returning an error.

Inside the method we are getting the collection with the help of struct and then using its .connection and then from that connection getting
the database and then from the database getting the collection.

then creating a filter which will work for us in the way that what data we want to delete therefore inside this providing the key and value
make sure to provide the key as same as mongo db and value should be same as you are taking as parameter.

Then on the collection we got we are calling the DeleteOne method of mongo and then providing the context and filter we made above to tell
what data we want to delete.


we dont care what this will return just checking for the error.

*/

func (mgr *manager) DeleteData(id primitive.ObjectID) error {

	orgCollection := mgr.connection.Database("GO-API").Collection("test-col")

	filter := bson.D{{"_id", id}}

	_, err := orgCollection.DeleteOne(context.TODO(), filter)

	return err

}

/*

Here we are makinf our last method which is UpdateData taking manager as parameter first so that we can use fields which we have defined
inside our struct and taking data which is of type user which is a struct and then returning an error.

Inside the method we are getting the collection with the help of struct and then using its .connection and then from that connection getting
the database and then from the database getting the collection.

then creating a filter which will work for us in the way that what data we want to update therefore inside this providing the key and value
make sure to provide the key as same as mongo db and value would be Id from the user we are providing as parameter.

then creating a update variable which is again using bson.D for structured thing and inside that using $set operator of mongo db to set the
updated data into the document.

calling UpdateOne method on the collection we got and inside it providing the context then the filter to find out that specific data
and then update to update the data , we dont care what we are getting just checking for the error and then returning it.

*/

func (mgr *manager) UpdateData(data User) error {

	orgCollection := mgr.connection.Database("GO-API").Collection("test-col")

	filter := bson.D{{"_id", data.Id}}

	update := bson.D{{"$set", data}}

	_, err := orgCollection.UpdateOne(context.TODO(), filter, update)

	return err

}

// So this was all in this session of Performing CRUD with Gin framework.

// Stay Tuned For More Updates.
