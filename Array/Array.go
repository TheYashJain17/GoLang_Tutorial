// In this session we gonna learn how can we make an array in go and use it.

// we have three types of array in go lanuage.

// fixed-size array - this is the array in which the size of the array is fixed.

// dynamic-size array - this is the array in which the size of array is not fixed.

//multidimensional array - this means arrays inside another array.

package main

import "fmt"

func main() {

	//fixed-size array

	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)

	/*this is the syntax of fized-size of array , we have to define type as well as the length of the array
	on both the sides and then in curly braces define the number of values we mentioned inside size and
	define according to type we mentioned int in this case.*/

	/*Now lets see what will happen when will provide a certain size to an array but will not provide
	that much arguments.*/

	var arr2 [4]int = [4]int{}

	arr2[0] = 1
	arr2[1] = 4
	arr2[3] = 8

	fmt.Println("This is our arr2", arr2)

	/*Here we can see that we havent provided 2nd index value , so the system does that for us and put a 0
	at that place , as we have provided the length of the array 4 and provided only 3 elements.*/

	var arr3 [4]string = [4]string{}

	arr3[0] = "Mumbai"
	arr3[1] = "Delhi"
	arr3[3] = "Kolkata"

	fmt.Println("This is our arr5", arr3)

	/*So now we check the same case(as above) for strings and in result we can see that now we are provided
	with extra space at the place of the element we didnt provided.*/

	//dynamic-size array

	var arr4 = [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(arr4)

	/*this is the syntax of dynamic size array in this case we dont have to metion the size of the array
	as we dont know what could be the size of array therefore using dynamic size array
	here we dont have to define the type and size on both sides , we just have to define them only on right side
	and for dynamic purpose we are using three dots which we also known as spread operator in javascript.*/

	arr5 := [...]int{1, 2, 3, 4, 4}

	fmt.Println(arr5)

	//We can make array like this also like we declare state variables , but only gonna work inside function.

	// if we want to know the length of the dynamic size array we can use len property of go.

	fmt.Println(len(arr3))

	//multidimensional array

	var matrix [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(matrix)

	/*this is the syntax of multidimensional array , with this syntax we are saying
	that make an array of size 3 and inside that array which is of size 3 , arrays made will be of size 3
	sounds confusing , print the result to understand this better.

	In simple words we ordered to make an array of size 3  , but instead of values we are saying to
	put arrays means inside of values there would be 3 arrays inside the array and of that array we are
	defining the size which is 3.
	So we are making an array of size 3 which consist 3 arrays and the arrays which are made inside
	that array is also of size 3.
	Thats how we can make a dimensional array.*/

	// if we want to see the particular value from the multidimensional array we can do it like this.

	fmt.Println(matrix[2][1])

	/*with this syntax we are saying to get the array which is on 2nd index and inside that
	get the value which is on 1st index , we can check by ourself that its gonna be 8
	and we also got the same result , this is how we get the particular value from
	a multidimensional array.*/

}

// so this was all about the array go.

// Stay tuned for more updates.
