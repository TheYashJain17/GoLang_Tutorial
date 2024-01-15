// In this session we gonna learn that how can we convert the type of a datatype in go.

/*Here we are seeing the conversion of the type of the value , means if there is a value of int32 and we want to
convert it to the int64 then we will see how we will do it.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// var a int32;

	// var b int64 = a;

	/*when we will do the thing we are doing above we will get an error saying
	cannot use a (variable of type int32) as int64 value in variable , it means here we have to convert the
	type to use it otherwise we will get an error.*/

	// so for converting the type we will do this :-

	var a int32

	var b int64 = int64(a)

	fmt.Println(a, b)

	/*this is how we do the conversion , and after doing this we will not get any error.const

	we can do similar things for other data  types also like float etc.*/

	// c := 2.5

	// var d float32 = c

	/*here we can see that we have not defined the type value but still we are getting the error ,
	so this means it is compulsory to convert the type in these type of cases.*/

	c := 2.5

	var d float64 = float64(c)

	// see now we are not getting any error.

	fmt.Println(c, d)

	// now lets see if we want to convert a int value into string then how we gonna do it.

	var strVal int = 500

	var convertedVal string = strconv.Itoa(strVal)

	/*now we have the converted value (int to string) , so for doing this we have to import a inbuilt module
	name strconv and then have to use it Itoa property to convert the int into string.
	This is how we convert an int value into a string value.
	if we will use normal string() property to convert the int into string then its not gonna happen,
	instead of converting it will give the eskai code of the desired thing.*/

	fmt.Printf("%v , %T", convertedVal, convertedVal)

	/*%v gives us the value of the thing we want and %T gives us the type of the thing we are looking for
	we have to keep both of these inside the same string and after this we have to give the desired thing
	means the thing of which we are looking the value and the type , we have written the same thing
	2 times because we are looking for 2 things(value and type) therefore we have to provide the value
	2 times.
	And for using this properly we have to use fmt.Printf command , we use fmt.println or normal print
	it will not gonna work , it will gonna work only with fmt.printf , so keep this in mind.*/

	// Now lets convert the type of the input  we are getting from the user.

	/*For this we have to use the bufio package of the go which we used inside Comma_Ok file,
	For taking the user input with the help of bufio and newreader.*/

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n In next line you can enter the input")

	input, _ := reader.ReadString('\n') //This means read till next line as we are running the ReadString method.

	fmt.Println(input)

	// In this way we will get the input entered by the user.

	// fmt.Printf("The type of the input is %T", input)

	// The answer would be string.

	/*But suppose we were asking for the ratings from the user , so obvioulsy we dont want the ratings in
	string form , we want them in int form , so for that we will convert the type of the input
	entered by the user.*/

	ratings, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if error != nil { //nil means null in go language.

		fmt.Println(error)

	} else {

		fmt.Println("Added 1 to your rating: ", ratings+1)

	}

	/*As we were getting the ratings in the string form , Now we have to convert those ratings/user input
	into the int to get it better ,
	So for that we are using strconv package and using .ParseFloat package of it , this will convert the
	string into number us , giving the input variable inside which we are getting the user input/ratings
	and 2nd parameter is the base means whether the required thing should be converted into int32 or int64
	and we have provided 64 means convert it to the int64.

	And we have used strings package , we have to do this otherwise we will get an error inside our
	terminal saying (whichever the input you will enter\n) this will be the error , so for removing this
	error we have to use strigs package and have to use its method TrimSpace which will remove \n and
	everything extra which giving us error.

	And atlast we are puting this in if-else statement saying if there is any error print the error
	else print the input.*/

	/*So we have to do this much work to get only this much result , yeah strange but we have to do it
	And we gonna use this same as on many places.*/

}

// So this was all about conversions/typeConversions.

// Stay Tuned For More Updates.
