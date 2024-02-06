// In this session we gonna learn about methods.

// Methods are very simular to functions but there are some differences , lets understand them.

package main

import "fmt"

// Here we are making a struct as we did before , sith the help of struct we will understand methods.

type user struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Made a struct of user type which is storing the details of the user.

func main() {

	yash := user{"Yash", "Yo@gmail.com", true, 22}

	// Here we are defining  a user with the the help of user struct and storing it inside a variable.

	fmt.Printf("The details of the user are %v and %v \n", yash.Name, yash.Email)

	// Printing the details of the user.

	yash.GetStatus()

	/*Now we are using the variable inside which we stored the details with the help of struct
	And on that we are using the method we created and now this will give us the desried result.

	Thats how we use methods.*/

	yash.NewEmail()

	/*Now we are again using a method on the variable in which we stored the details with the help of
	struct , with this method we are changing the email of the user on which we used this method,*/

	fmt.Printf("The details of the user are %v and %v ", yash.Name, yash.Email)

	/*IMPORTANT NOTE : The change we made in email will be on temporary basis as its not gonna change
	the email permanently , its just for the temporary basis.

	We can see this by again printing the details of the user , and we will see that we will get the
	actual email of the user and not the temporary one.*/

}

func (User user) GetStatus() {

	fmt.Println("Is user active: ", User.Status)

}

/*Now here we made a method , this look kinda like function but it is not.

In case of methods we provide the parenthesis first then write the name of the method and then
parenthesis(()) again , unlike functions in which we first write the name then the parenthesis.

This is the major difference between functions and methods.*/

/*Now in this method we are checking the Status of the user which we defined as a field in our struct.

Inside parenthesis we are taking a variable User which is of user type (means our struct type)
Then only when we use it on a struct , it will work properly.

Then we are printing the Status of the user means the person on whom we will use this method.*/

func (User user) NewEmail() {

	User.Email = "NewMail@gmail.com"

	fmt.Printf("The new email of the user %v is %v: \n ", User.Name, User.Email)

	/*We have made another method through which we gonna change the email of any user.

	Here also we have provided the parameter in which provided a variable and then define its type
	which is of struct type , now can enter any struct in it.

	Then we are targettng the Variable which we defined inside the method as it is also a struct
	so we are targetting its Email property and changing it with the email we want to.

	Then printing the value of the user name and email.

	IMPORTANT NOTE : The change we made in email will be on temporary basis as its not gonna change
	the email permanently , its just for the temporary basis.*/

}

//Stay Tuned For More Updates
