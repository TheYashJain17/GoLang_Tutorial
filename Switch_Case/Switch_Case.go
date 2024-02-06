// In this we are seeing how to use switch case , its kinda like if else statement.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// here we are declaring a variable with value 5;

	i := 5

	// now we are using switch case.

	switch {

	case i == 5:

		fmt.Println("Its 5")

	case i == 10:

		fmt.Println("its 10")

	default:

		fmt.Println("None of the above")

	}

	/*with switch case we are saying if the case is this then run this and if it is this then run this
	and atlast default , it runs when none of the above cases run its like else condition.*/

	switch {

	case i == 5:

		fmt.Println("Its result")

		fallthrough

	case i == 8:

		fmt.Println("Its fallthrough case")

	}

	/*there is fallthrough in switch case which prints the desired result as well as run the
	next case also.*/

	/*after executing the program we can see that the desrired result gets print as well as the next case gets
	also print.*/

	// (REALLY IMPORTANT)

	// Lets explore switch case more.

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice number is ", diceNumber)

	/*So in this case suppose we are making a ludo game and we know there are  1 - 6 numbers in ludo
	and thats how the game is played by throwing the dice and getting a random number ,
	So here also we are doing the same thing getting a number randomly.

	We are getting a number randomly with the help of rand module of math which helps us to get the
	random number.

	So first we have to seed with the help of rand.seed , this is important as we can get random number
	only after doing the seed process and inside it we can use time.Now() like we did now ,
	it would be enough but for extra security use Unix.Nano() like we did above , just follow the exact steps
	we are following above.

	Now we are ready to get a random number , we want to get the int so we are using Intn and inside it
	using 6 means random number would be till 5 as range is not included so at the end we are using + 1
	We have used 6 as we are following ludo approach.

	Now after this we will get numbers randomly.*/

	switch diceNumber {

	case 1:

		fmt.Println("Dice value is 1 so you can open now")

	case 2:

		fmt.Println("Dice value is 2 , so you can move 2 spots")

	case 3:

		fmt.Println("Dice value is 3 , so you can move 3 spots")

	case 4:

		fmt.Println("Dice value is 4 , so you can move 4 spots")

	case 5:

		fmt.Println("Dice value is 5 , so you can move 5 spots")

	case 6:

		fmt.Println("Dice value is 6 , so you can move 6 spots and can have one more turn")

	}

}
