// In this session we gonna learn that what is comma ok syntax and how can we use it.

/*We are continouing this from map file , in that file we have understand the basic of comma ok syntax.
Now in this fie we will understand more advance things of comma ok syntax.*/

// So in this file we are gonna take the input from the user.

// User gonna put the input inside terminal and there he/she will get the response.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome User" //we have already learnt to declare variables like this in our previous session.

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) //Stdin means standard input. Here we are taking the standard input from the user.

	fmt.Println("Enter the rating for our pizza")

	/*Here we are just using bufio package of go and with the help of that we are reading a file
	in standard input way.*/

	/*In programming language we have if else statment to look for the error , in go also we have that
	but for the short-form comes the use of
	comma _ok as it do the work for us , As in go language it is believed that error will come in
	true false form therefore comes the role of comma _ok syntax.*/

	input, err := reader.ReadString('\n') //here we are saying to read it as a string.The imput entered by the user will be treated as string.

	fmt.Println("Thanks for rating us , ", input)

	fmt.Println("Thanks for rating us", err)

	/*So this imput , err is like our try catch respectively , as we are getting the ans in our input variable
	and if there is an error that would be stored inside err variable.
	We can use any name instead of input and err , still it will gonna do the same job for us.*/

	// Basically we write input , _ok therefore we called it comma ok syntax.

	/*If we dont bother for one of the thing like in some cases we dont want to know the input entered by the user
	or we are not getting any error we already know that , in that case we can use _ (underscore) instead of
	that variable.*/

	// For ex if we dont care about the input we are getting from the user , we can write :-

	/*_, error := reader.ReadString('\n')

	fmt.Println(error)*/

	/*In this case we dont care about thats why we have written an underscore at the place of the name of the
	variable.*/

	/*Similarly if we dont want to know the error , just do the same thing as above.*/

	/*inputByUser , _ := reader.ReadString('\n');

	fmt.Println(inputByUser);*/

	/*Here in this case  we are just focusing on the input entered by the user , therefore we have put
	an underscore (_) at the place of error variable.*/

}

/*So this was all about the commma ok syntax and we will see multiple situatons in which we will be using
comma ok syntax , so stay tuned for that.

And this is the just the surface level things in go we gonna make a deep dive into it.

So stau tuned for all those things.*/
