//Continuing Our API Development Journey

// In this session we gonna learn how can we create json data in go  , Really Important Topic

package main

import (
	"encoding/json"
	"fmt"
)

//Creating a struct so that we can enter data with the help of this struct and then will convert it into json.

/*Here we have taken first letter as small letter(course) and not the capital as here is no need of exporting
this struct , therefore we can take the small letter , if we want to export this then always remember to use
the first letter as capital letter.*/

// Now inside this struct we are defining the desired fields and also their data types.

type course struct {
	Name     string   `json:"coursename"`        //Here we are mentioning json means when we create this data in json replace this field name with the one we are mentioning like in this case we have given Name as field but when we see the data in json we want this field name to be converted in coursename , thats why we have written this line.
	Price    int      `json:"price"`             //we are mentioning json means when we create this data in json means when we create this data in json replace this field name with the one we are mentioning like in this case we have provided with the field name in Capital letter but when we are providing this data as json we want this field name with small letter , now you must be thinking that if we want to show the field name with the small letter then just do it normal field name then there would be no need of converting the name in json , but it is not easy as it seems as we will give name to our field with small letter then our system wont be able to reach it , you can try by giving any field name small letter and you will get to know the same.
	Platform string   `json:"learning platform"` //here we are doing the same which we did in first line
	Password string   `json:"-"`                 //Now here comes something different till now we have seen that writing json means we want to show this(anything we are writing after json) in json response , here we have provided -(hyphen) which means dont show this field means when we are getting this data as json in response we will not be able to see this field , as we know password is a crucial thing and it should not be seen so we should hide this and that is what we are doing here.
	Tags     []string `json:"tags,omitempty"`    //Now lets come to this field here the field name is Tags means there are more one things in this field therefore we have taken a slice of string so that can store multiple things in it, now here we are saying first show the field name with small letter tags and not Tags so we have written the code for it and secondly we are using omitempty which means if the field is empty then dont show it simple , we are doing it with this field as it optional to put tags therefore we are doing so.
}

/*Note :- kindly use the same format as i have used no extra spaces or nothing otherwise it will cause an error
And i have used back ticks not string , inside backticks i have used string.*/

func main() {

	fmt.Println("In this session we gonna see how can we create json data")

	creatingJSON() //calling the function inside main function so that we can execute it.

}

func creatingJSON() { //now we have created a function and in this function we gonna execute everyting.

	courses := []course{ //we are making a slice of type course which is a struct which we created above and storing it inside a variable courses.We are making a slice of struct as want to store multiple course details in it and not a single one.

		{"Solidity", 299, "linkedin/theyashjain17", "sol123", []string{"web3", "blockchain"}},
		{"GoLang", 499, "linkedin/theyashjain17", "gol456", []string{"backend"}},
		{"DEFI", 699, "linkedin/theyashjain17", "DEF789", nil}, //here we have provided nil in the place of tags to check whether the omitempty property we have given is working or not.
	}

	jsonData, err := json.MarshalIndent(courses, "", "\t") //now we are using json module and its property MarshalIndent to convert our data into json , this property takes 3 parameters , first the data we want to convert into json which we stored in coures variable , second any prefix if we want to provide in our case we dont want to thats why providing empty string and third the indentation in simple words the spacing technique means how much space we want give to every line here we have given /t which means space of tab. And again storing this in a varible and could get an error therefore also storing it.

	if err != nil { //checking the error.

		panic(err)

	}

	fmt.Printf("%s \n", jsonData) //here are have provided %s means print the string format at every new line , thats how we will get to see all the json data in string format , if we dont use it then we will see this data in bytes format , therefore we are specifying the format for this json data which is string in this case.

}


// this was all for this session.

// Stay tuned for more updates.