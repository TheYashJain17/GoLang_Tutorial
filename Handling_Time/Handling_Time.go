// In this session we gonna see that how can we handle time in go language.

package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("This is the current date and time we are getting")

	presentTime := time.Now()

	//With the help of above statement we can get the current date and time.

	fmt.Println(presentTime)

	// We are printing the current date and time with the help of fmt package.

	/*After printing the date and time in this way we were getting it in little ugly way so therefore
	it is important to format the date and time little bit.*/

	fmt.Println(presentTime.Format("02-01-2006  15:04:05 Monday"))

	/*So for formatting the date and time we have to use .format on the variabe in which we store the
	date and time we got and inside that we have to provide parameters to format the date and time.

	So if we want to format the date and time we have to pass the exact same parameteres as we have passed
	above it should be exactly same word to word , (we can read more about this in the documentation of
	go in time section) then only we can get the formatted date and time.*/

	createdDate := time.Date(2000, time.November, 17, 00, 00, 0, 0, time.Local)

	/*If we want to create any particular date we can do it like this(see above line) , use the time.date
	and inside provide the same number of parameters as we have provided inside this
	first is the year then if we want to choose the month we have to do it in this way like we did here
	and after that the day then hour then then minutes then seconds then nano seconds and at last time factor

	just follow the same process as we have followed here.*/

	fmt.Println("This is the date and time which we are creating")

	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("02-01-2006 15:04:05 Monday"))

	/*5Now atlast formatting the date and time we created to show it in a better way , have to provide the
	parameters in exact same otherwise we wont get the formatted date and time.*/

}

// So this was all about the time in go , you wont gonna use much of it and also not gonna use more than this.

// Stay tuned for more updates.
