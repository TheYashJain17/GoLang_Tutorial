//In this session we gonna learn about Loops , break and continue and GoTo Keyword.

/*We have seen multiple type of loops in multiple languages like for loop , while loop and do while loop.
But in go language there is only one loop and that is for loop.
Now lets see how we use a loop in go language.*/

package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		fmt.Println("Value increasing by ", i)

	}

	/*This is how we can use loop in go language , syntax gonna be remain same , we can change the sequence
	if we want.*/

	j := 0

	for j < 10 {

		fmt.Println("Another value increasing", j)

		j++

	}

	// This is how we can change the sequence if we want to , still we will get the same result.

	// Lets understand this more with more examples.

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"} //This is slicing

	fmt.Println(days)

	for d := 0; d < len(days); d++ {

		fmt.Println(days[d])

	}

	/*First we made a slicing and inside it enter all the days , then with the help of for loop we are
	iterating every value.
	The benefit of making a slicing is that now we can use len method which can only be used on slicing
	and we can run the loop till the length of the slicing.*/

	// Now lets use loop with the help of range keyword and lets see how it gonna make our work easy.

	for i := range days {

		fmt.Println(days[i])

	}

	/*So this is kinda simple as compare to for loop , as here we just have to use range keyword
	and thats it , its gonna do the most of the work for us.*/

	// Now lets also see the foreach loop in go language.

	for j, day := range days {

		fmt.Printf("index is %v and the value is %v \n", j, day)

	}

	// So this is how we use foreach loop in go.

	// We can use any type of loop which we want.

	// Now lets see the break and continue statement

	value := 1

	for value <= 10 {

		if value == 8 {

			break

		}

		if value == 4 {

			value++

			continue

		}

		fmt.Println(value)

		value++

	}

	/*In this session we made a loop which will gonna run till value 10 and also will include 10 as we have
	used <=.

	 And we have used break and continue statement in this.

	We are saying if the value == 8 then break the loop means terminate the loop , it means
	loop will not work after reacing the value 8.

	Then we are saying if the value == 4 then continue means skip that value and move forward

	Means all the remaning values will be printed except the value in which we used continue

	And remember to increase the value in this case , otherwise the loop will gonna run
	infinite times and thats what we dont want , therefore it is important to increase the value.*/

	// One more thing which we can do with the help of loop, lets see that also

	k := 0

	for k <= 10 {

		if k == 2 {

			goto hell

		}

		k++

	}

hell:
	fmt.Println("Welcome to hell")

	/*There is a keyword names goto which helps us to go to that specific line of code.

	In this session , we made a loop inside which we are saying ,

	If value of k == 2 then goto the specific line of code and run it.*/

	// We can see above , how to define statement which we can use with goto.

}

//stay tuned for more updates.
