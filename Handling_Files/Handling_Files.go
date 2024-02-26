/*In this file we gonna see how to handle files with go , like how we can read or write files with the help of
go like with the help of os module , we gonna create a new file and then how we can read it or any other file.

Also I have introduced a new file go.mod , dont worry about that i will make you undertand this 
later in detail.*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// Now we gona create txt files with the help of go.

	content := "Just go to Hell"

	location, err := os.Create("./hellfile.txt")

	if err != nil {

		panic(err)

	}

	length, _ := io.WriteString(location, content)

	fmt.Println("Length of the string is :", length)

	defer location.Close()

	/*First of all we are storing the string which we want to write in a new file in the content variable,

	After that we are using os module and its Create method to create the location of the file
	means where we want to create the file , now there can be 2 outputs first the successfull location
	or second we will get an error ,

	So we are checking that if the error is not equal to nil then panic the error means print the error.

	Now we gonna create the file with the help of io module and its WriteString method , inside it
	we first gonna provide the location of the file means where should the file be created and second
	the data which should be written in the file which is inside our content variable ,

	Now again there can be 2 outputs first the file has been successfully made and second there could
	be any error , but now we know that there would be no error ,

	So we are just directly printing the length of the string which get written in the newly made file

	This will give us the length as it is written in the documentation of go , so we are printing the
	length to know whether the operation is successfull or not.

	And atlast we have to use .close on the variable inside which we store the location of the file
	And we are using defer keyword here so that this operation gets executed at last.*/

	/*After running this code we will see that a file has been created with name hellfile.txt
	and inside that we could see the text whixh we want to write inside it.
	And in the console we will have the length of the string which we wrote inside the new file.*/

	readFile("./hellfile.txt")

}

func readFile(filename string) {

	databyte, err := ioutil.ReadFile(filename)

	/*if err != nil {

		panic(err)

	}*/

	checkNilError(err)

	fmt.Println("The data inside the file is :\n", string(databyte))

	/*So now for reading the file we created or any other file we are using ioutil module and and its ReadFile
	property and inside it we are providing the filename which we taken as a parameter in string form
	as we have to provide the exact location of the file we want to read and its gonna be in string format
	therefore we have do so.

	Now there could be 2 things which can happen first we will get the data successfully and second
	we will get an error so we are treating the error with checkNilError function which we will discuss about
	below.

	We have taken the name databyte of the variable as if we will get the data it would be in databyte form

	So when we are printing the result we are typecasting the resulted data into string

	And after running this we will get the data which is present inside that file.*/

}

func checkNilError(err error) {

	if err != nil {

		panic(err)

	}

	/*We have made a function to check for the error , inside this we are writing the same code which
	we are writting again and again to check for the error and now we can use this function in any function
	and pass the error variable inside it to check for the error.*/

}

// Stay tuned for more updates.
