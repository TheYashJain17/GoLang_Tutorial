// In this session we gonna learn about a really important thing which is variables.

// We have three types of variables available in Go language.

// local variables
// global variables
// package variables

// And There are 4 ways in which we can made/declare and initialize the variables.

// 1st way of declaration is :-

// var a int //here we are declaring the variable

// a = 500 //and here we are initiliasing the variable means providing the value of the variable declared above.

// 2nd way of declaration is :-

// var b int = 500 //here are declaring the variable as well as also initialising its value.

// 3rd way of declaration is :-

// var c = 500 //here are declaring the variable as well as also initialising its value.And also here we are not using int as it is not necessary to use.

// 4th way of declaration is :-

// d := 500 //this is another and shortest way of declaring and initialising a variable. But the important thing is that this kind of declaration only works in case of local variable and not in global as well as package level variable. so if we want to declare variable with this method remember to use this only in case of local variable.

package main

import "fmt"

///local variables

// if we are declaring any variable inside a function then it is a local variable same as javascript.

func main() {

	var myvalue1 int //this is the local variable as it is defined inside a function and as we havent provided any value inside it , so in result we will get 0 as answer.

	myvalue1 = 5000

	fmt.Println(myvalue1)

}

///global variables

// global variables are the variables which we are defined outside the function same as javascript.

/*global variables are always declared in Pascal Case means the first letter of the name of the variable
should be capital.
it is important to keep the global variable name in Pascal Case , we will understand this better while we
will read about the Package Level variables*/

var MyValue2 = 500

///package level variables

/*Package level variables are always declared in camel Case. It is important to keep the package level
variables in camel Case and not in Pascal Case otherwise this will become global variable.
Same case is with global variables if we declared global variables  in camel Case they will become
package level variables.
So always keep this in mind.*/

var myValue3 = 50

/*In case of variables sometimes  we faced a problem name shadowing , this problem arises when we give the
same name to the local variable as well as to the global variable or package level variable.*/

/*So in the above case when we will print the result we will get the value of the local variable and
not the global variable or package level variable.
As the machine will gonna look for the nearest thing means the local variable is near than the global
or package level variable , therefore it gonna print the value of local variable in case of same names.*/

// So this was all about the session hope you have understood all the ways of declaring variables.

// stay tuned for more updates.
