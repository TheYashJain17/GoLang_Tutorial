// In this session we gonna understand what is mapping in go. 

/*Map is similar to mapping which we have read in solidity , here also we are assigning a value to a key,
Like we do in solidity , like assigning a value to an address and like that.
Lets understand map of go language.*/

package main

import "fmt"

func main() {

	employeeSalary := make(map[string]int) //declaration

	/*This is how we create a mapping in go language , with this mapping we are saying that
	we are mapping an int value with the string value.*/

	fmt.Println(employeeSalary)

	// when we will print this we will get an empty map.

	// Now lets see how we map values.

	employeeSalary["Sohan"] = 20000
	employeeSalary["Rahul"] = 25000
	employeeSalary["Mouni"] = 15000
	employeeSalary["Rohan"] = 10000

	fmt.Println(employeeSalary)

	// In this case we are initialising values separetely.

	// If we want to declare as well as intialise values together , then we can do it like this:-

	employeeSalary = map[string]int{ //initialisation

		"Neha":   20000,
		"Ayush":  25000,
		"Raj":    22500,
		"Ayushi": 26000,
	}

	/*This is how we insert values inside a map , using the map we declared above ,
	And remember to put a coma (,) at last entry like we did above otherwise we will get an error.*/

	fmt.Println(employeeSalary)

	// When we will print this we will the desired output.

	// If we dont want to separately do these things(declaration and initialisation) then we cam do this:-

	allExpenses := map[string]int{ //declaration with initialisation

		"Stationary":  200,
		"House":       1000,
		"Electricity": 5000,
	}

	// this is how we declare as well as initialise a map in go.

	fmt.Println(allExpenses)

	// printing to see the desired result.

	// If we want to know the length of map then use len method like this:-

	fmt.Println(len(allExpenses))

	// This will get us the length of the map.

	// if we want to get a particular thing or person value's from the map then we can do this by writing:-

	fmt.Println(allExpenses["House"])

	/*just enter the map from which we want to get the value then enter the key whose value we want to
	access and thats it we will get the desired value.*/

	// If we want to change the value of any key of any map then we can do it like this:-

	allExpenses["House"] = 2000

	fmt.Println(allExpenses["House"])

	// Now after printing the result we can see that value of house has been changed.

	// If we want to add a new value inside a map then we cam do it like this:-

	allExpenses["Water"] = 3000

	fmt.Println(allExpenses)

	// After printing this we will be able to see that a new key value pair has been added in our map.

	// If we want to delete a key value pair from the map then we can do it like this:-

	delete(allExpenses, "Water")

	fmt.Println(allExpenses)

	// After printing this we can see that the particular key value pair we mentioned above has been deleted.

	/*If there is a key which isnt present inside the map then after printing that particular key in result ,
	we will get the default value of the value which is 0.*/

	// So if we want to know whether a key is present inside a map or not , we can do it like this:-

	Water, ok := allExpenses["Water"]

	/*So for finding out whether the key we are looking for is available or not , we gonna store
	the key of that particular map in a variable and with coma we gonna use ok , this ok will gonna
	store the boolean value means true or false , Thats all now we can find out whether the particular
	key exist or not.

	It is not important to keep the name ok we can use any other name which we want like flag , other etc*/

	fmt.Println(Water, ok)

	/*We have to print both the variables , first one wlll return a value and another one means ok will
	return a boolean value means true or false.*/

	/*So after printing the result if we get false(the value of ok) , then it means that particular key
	isnt available in the map and if we get true the that particular is available in the map.*/

	// Lets see how to use loops in map.

	for key, value := range employeeSalary {

		fmt.Printf("Salary of %v is %v\n", key, value)

	}

	/*This is how we use loops in map , there is completely a different syntax of loops for map
	As according to map we have key value pairs so suppose we want to access both values
	so we are taking 2 variables key and value and then using  range keyword and after it defining
	the map name which is employeeSalary in this case , now inside it we are using printf method of fmt
	because we want to use placeholders(%v or %T are known as placeholders) , so now we are printing the
	value with the help of fmt.printf and using %v placeholder to get the value of the variables
	we define above  and using \n to get the data in the new line otherwise
	all the data will gonna be print in single line , and atlast we are providing the names of variables
	inside which all the values exist this  will give us the desired result.*/

	for key, value := range allExpenses {

		fmt.Printf("EXpense of %v is %v", key, value)

	}

	/* LEARN ABOUT COMMA_OK SYNTAX IN MORE DETAIL IN COMMA-OK.GO ,
	THERE ARE MORE DETAILS AVAIABLE REGARDING THE SAME.*/
}


// so this was all about mapping in go.

// Stay tuned for more updates.
