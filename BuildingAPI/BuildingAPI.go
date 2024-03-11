/*In this session we gonna see how we can create a proper REST API but locally as we gonna use 
slice instead of DB.
Dont worry we are just getting a basic idea that how we can do it later on we will work with the mongodb.

This session gonna take multiple sessions but i am putting this in a single
so that you can understand it better.*/

package main

func main(){



}

								// MODEL SECTION

/*we gonna make a struct and will make a slice of that struct , as if you have follow my struct 
tutorial then you know how to create a struct and after that we are using json , you know what it means , 
it means that what will the user as a response in the form of json thats it , nothing new.*/

type Course struct{

	CourseId string `json:"id"`
	CoureName string `json:"name"`
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


func (_course *Course) checkIfEmpty(){

	return _course.CourseName == "";

}

// stay tuned for more updates.

