/*In this session we gonna see how we can create a proper REST API but locally as we gonna use 
slice instead of DB.
Dont worry we are just getting a basic idea that how we can do it later on we will work with the mongodb.

This session gonna take multiple sessions but i am putting this in a single
so that you can understand it better.*/

package main

import(

	"encoding/json"
	"net/http"
	"fmt"

	"math/rand"
	"strconv"
	"time"
	"github.com/gorilla/mux"

)

func main(){



}

								// MODEL SECTION

/*we gonna make a struct and will make a slice of that struct , as if you have follow my struct 
tutorial then you know how to create a struct and after that we are using json , you know what it means , 
it means that what will the user as a response in the form of json thats it , nothing new.*/

type Course struct{

	CourseId string `json:"id"`
	CourseName string `json:"name"`
	CoursePrice int `json:"price"`
	CourseAuthor Author `json:"author"`//made this of Author type which is the struct we made below.
	Tags []string `json:"tags,omitempty"`

}

// Integrated struct into another struct just to get more information and also increasing little bit complexity.

/*
NOTE- Usually these structs goes into a separate file but here we are making it in one for the simplicity

If you have ever worked with nodejs or have a knowledge about backend so you can understand this as a model.

So you can understand it like this like we are making model in this session
*/


type Author struct{

	AuthorName string `json:"authorname"`
	AuthorWebsite string `json:"authorwebsite"`

}

var courses []Course;//Now this will work as our local database or can say fake DB. Our all work gonna revolve around this now.

// We are creating a method not a function to check whether the course exist.

/*
In this method we are a taking a parameter from the user that will be of Course type which is our struct
and using pointer to get the exact data and not a copy of it
And we are saying that it will gonna return a bool value

Now inside the function it is returning that the course we are providing as parameter ,
whether its id is equal to empty string and its coursename is equal to the empty string means it is empty
if that so it will gonna return true which means the provided course doesnt exist.

And if you have taken my methods tutorial then you must know that how do we use a method
*/


func (_course *Course) checkIfEmpty() bool {

	return _course.CourseName == "";

}


/*
In our above session we saw kind of models , so now in this session we gonna see controllers.

You can understand controllers as fucntions which gets executed whenever we make a request to the route.

You can understand this if you have seen my sample server tutorial , so now lets see a example of it.

Usually controllers are also made in different file but here we are keeping this in the same file to keep
the things simple.

*/

/*
So created a function can say controller to the control the homepage route so whenver someone
will visit to our homepage he/she will see this response , i have discussed about this in detail
in sample server tutorial so kindly check it for more details.
*/

func renderingHomePage(w http.ResponseWriter , r *http.Request){

	w.Write([]byte("<h1>This is Our Homepage</h1>"))


}

/*
So this is our another controller which gonnna handle the route whenver someone wanna see all the
courses we have.

So to do this we are making this function or can say controller.

So again taking the paramters as it is required to do so , check sample server for more details.

Therefore doing so and inside the funciton  we are first printing to tell what we are doing with
this controller.

After that we have to set headers , if you know about something backend then you must know about headers
And if you dont know that what does they do let me tell you in easy words ,

So basically till now we have understood that we deal data in json for, in backend , so by setting header
we are saying that we are setting the content-type to json , it should be written in the same way  as i have
written below else you will get error , so basically by doing this we are telling that the type of data will
gonna be of json.

after setting the header we are sending the response in the form of json with the help of json library.

So with the help of len method we are checking of the length of courses slice that whether it is greater than 0 or not

if the length is greater than 0 then we are sending the response.

so for sending response in json we have to use NewEncoder property of json and inside that we have to provide
a writer to write the response which is w in this case what we have taken in paramter this will gonna do the
job of writing the response for us and then we have to use the Encode property to send the data and inside
this we are passing the slice of Course struct we declare above.

So thats how user can get all the courses as gonna store all the courses in this slice.

else there will be a response saying there are no courses to display.

for starting there are no courses , but after sometime we will gonna add some data manually , this process
is known as data seeding.

dont worry if you are finding this little complicated , its absolutely fine to feel so , its just about the
understanding the syntax, you  will get familiar with the syntax soon dont worry just code along with me.
*/

func getAllCourses(w http.ResponseWriter , r *http.Request){

	fmt.Println("Getting All Courses")

	w.Header().Set("Content-Type" , "application/json");

	if len(courses) > 0 {

		json.NewEncoder(w).Encode(courses);


	}else{

		json.NewEncoder(w).Encode("There are no courses to display")

	}

}	


/*

Now this is our controller for getting a single course based on the id the user is providing.

In case of these controllers most of the process will gonna remain same like taking parameters
like http.ResponseWriter and Request then inside the controller/function we have to set the content-type
which is application/json so that we can send in the form of json and user can get the response.

Now in this controlller the main thing is getting the id from the user because on the basis of that id
we gonna find the course , the user is looking for.

so for getting the id which the user is providing we are using mux which is gorilla mux which we installed
and using .Vars property of it and inside providing the r which is handling the user request.

Now we got all the things in the form of key value pairs inside params variable.

Now we are looping through all the courses with the help of for loop using range keyword to match
the id with the user is providing with the actual course id to see whether there is such id or not.

if there is one then we are sending that in the form of json with the help of json library and its Property
of NewEncoder and writing with the w variable which is our ResponseWriter then using .Encode to send the course
we found in the form of json, then atlast using return keyword to say that task is completed now return
its upto use whether you want to use it or not as it gonna work same ways.

and atlast after exiting the loop we are sending response in json saying that there is no course with this id.

and atlast again using return to say that process is completed and now return again its upto you whether you
want to use it or not as it gonna work same ways.

*/

func getOneCourse(w http.ResponseWriter , r *http.Request){

	fmt.Println("Getting Single Course");

	w.Header().Set("Content-Type" , "application/json");

	params := mux.Vars(r);

	for _ , course := range courses{

		if course.CourseId == params["id"]{

			json.NewEncoder(w).Encode(course);

		}else{

			json.NewEncoder(w).Encode("Please Enter A Valid Id");

		}

	}


}	

/*

So Now Making another controller/function for adding a course into our fake db for now , dont worry we
will see/use a real database.

So in the beginning  we are doing the same things like we were doing till now , it would be better if you
would type this again and again then only you will get familiar with the syntax instead of copy paste
you can do the copy paste when you get familiar with this.

so first of all we are checking that if the user is providing something in the body or not.

if there is nothing in the body then we will send the response saying ody of the course cannot be empty.

Now after checking the body we are creating a variable of type struct

After this we are decoding the data we are getting from the user which is stored inside r as it is our request
and with r.Body we are getting request.Body so getting the data provided by the user and with the help of
Decode we gonna decode the data into our struct format  as we are providing the refrence of the variable
which is of our struct type.

After this we are using isEmpty method on the course we are getting to check whether the course we are getting
is valid or not , if the course is not valid then we are sending a response
saying Data is empty , exiting from the process and after that using return.

Now we have to enter the course which the user is providing , so for that we have to generate an id
so we can add that course with respect to that id , so therefore we will generate an id now.

So for generating a random number i have used rand(this should be math/random and not crypto/random , i have
seen many people making this mistake then complain about not getting the random number or some error).

rand.Seed has been depreciated therefore i am using different method that will gonna generate the random number
for us

Dont worry we are doing this just for our fake db as we will use mongo or any other database that will
automatically generate an id for us.

After getting that random number we are setting it in the place of courseId and after that adding new course
into our fake db with the help of append method , if you have taken my previous sessions then you know
how to use append method and how we can create a new slice with the help of append method.

Sending the new course as response to show that the new course has been added successfully into our fake db.

Then at last using return.

*/


func addOneCourse(w http.ResponseWriter , r *http.Request){

	fmt.Println("Adding One Course");

	w.Header().Set("Content-Type" , "application/json");


	if r.Body == nil{

		json.NewEncoder(w).Encode("Please Provide Data To Add Course")

	}

	var course Course

	json.NewDecoder(r.Body).Decode(&course);

	if course.checkIfEmpty(){

		json.NewEncoder(w).Encode("Body of the course cannot be empty");

	}

	source := rand.NewSource(time.Now().UnixNano());

	randomNumber := rand.New(source);

	course.CourseId = strconv.Itoa(randomNumber.Intn(20));

	courses = append(courses , course);

	json.NewEncoder(w).Encode(course);

}	

//stay tuned for more updates

