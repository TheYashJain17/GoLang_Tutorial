// In this session we gonna learn about slices in go.

/*We have heard about the slicing , that is used in an array like in javascript we say slicing this array
means taking out a paticular part from the array.
Slicing is a part of array in javascript but here slicing is completely another topic not a part of array.

Slicing is really very similar to array even its syntax is also very same.*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	var arr1 [3]int = [3]int{1, 2} //fixed size

	fmt.Println(arr1)

	var arr2 = [...]int{1, 2, 3, 4} //dynamic size

	fmt.Println(arr2)

	// Both of the above are our arrays , One is fixed size array and another one is dynamic size array.

	// Now lets see the syntax of slicing.

	var arr3 []int = []int{1, 2, 3, 4, 5, 6}

	/*this is the difference between array and slicing as we dont have to define any size in case of slicing
	neither dont have to put three dots [...] like this which we do in array to define a dynamic size array.*/

	fmt.Println(arr3)

	/*Major difference between array and slicing except the syntax is that , when we define a fixed size array
	we have to define those number of values which we defined in size otherwise we will get a 0 at that place
	of value.
	But in slicing this is not the matter  , here we have two different properties capacity and length
	so we defined a capacity of 4 and we are defining only 2 values then also there is not any problem
	as we will get only 2 values and no 0 will be there.*/

	/* This is the make keyword of which we talked about in memory management in go
	(go check that file to know better and also go to the bottom of the page to check more details).*/

	var arr4 []int = make([]int, 4, 8)

	arr4 = []int{1, 3, 5, 4, 5, 6, 7, 7, 8, 9, 9}

	/*We can also make slicing with the help of make keyword inside it defining the array of int type
	and then defining the length of the array and then defining the capacity of the array
	Means length of this array is 4 but it can contain 8 values.

	But as we gonna put the values inside the slicing , the length and capacity will be changed according
	to the values we put inside the slicing and will gonna be similar.
	And now the slicing is kinda like the dynamic size array.
	Change the vlaue of slicing and check it by yourself.*/

	fmt.Println(arr4)

	fmt.Println(len(arr4))

	fmt.Println(cap(arr4))

	/*Now with the len method we can see the length of the slicing and cap method we can see the capacity
	of slicing.*/

	// if we want to add a value inside a slicing we can do it like this:-

	arr4 = append(arr4, 5000)

	/*just use the append method and inside that provide the array in which we want to add the new value
	and after that provide the value which we want to add , thats it new value will be added.*/

	fmt.Println(arr4)

	// we can make new slicing with the append method like this:-

	arr5 := append(arr4, 540, 900, 780)

	// This is how we make a new slicing with the help of old slicing and append method.

	fmt.Println(arr5)

	// if we want to combine two slicing in one new slicing then we can do it like this:-

	arr6 := append(arr3, arr4...)

	/*this is how we combine two arrays in one array , it is important to put three dots ... at the end
	then only we can combine them otherwise we will get an error.
	We can combine only 2 arrays not more than that.*/

	fmt.Println(arr6)

	/*if we want to slice the slicing like we slice an array means taking out a part from the array
	We can do it like this:-*/

	arr4 = append(arr4[1:])

	fmt.Println(arr4)

	/*With the above statement we are saying that take out the element which is present at 0 index and
	print the slicing from index 1 , this will result in removal of 1st element from the slicing and
	the rest of it will be printed.*/

	// If we want to print out only limited items from the slicing , we can do it like this :-

	arr4 = append(arr4[1:5])

	fmt.Println(arr4)

	/*With the above statement we are saying to make a new slicing from the arr4 which will consists
	the elements starting from 1 index to 5th index and rest of the elements will gonna be ignored.

	This means we are saying start from index 1 and go till index 5 and remember last index element
	will not gonna print only before from that index will be printed.*/

	// Another way of creating slicing with the help of make keyword

	arr7 := make([]int, 5)

	arr7[0] = 234
	arr7[1] = 945
	arr7[2] = 465
	arr7[3] = 867
	arr7[4] = 777

	fmt.Println(arr7)

	// Here we will easily get the desired slicing which we made (arr7);

	/*Now we can see that we have provided limit of 5 in arr7 means we gonna put atmost 5 values in the
	slicing , So what to do if we want to put more values in it , Lets see :-*/

	arr7 = append(arr7, 555, 666, 321)

	fmt.Println(arr7)

	/*After printing this we can see that we didnt get any error instead the values got add on
	in the arr7 . this happened even after defining the size of slicing which is 5 in this case
	because it reallocates the memory and all the values will gonna add in the array.*/

	// Lets see some more methods which array doesnt provides us but slicing do.

	sort.Ints(arr7)

	fmt.Println(arr7)

	/*sort is a package is present in go and can be used in slicing and it has many methods to use
	Here in this case we are using .Ints method which gonna sort the slicing into the increasing order
	it means all the values which will be present inside slicing will gonna be set into increasing order.*/

	// Now lets see how can we remove any value from the slicing based on its index.

	courses := []string{"solidity", "Blockchain", "Ethers.js", "Web3.js", "Hardhat", "Truffle"}

	fmt.Println(courses)

	// Now lets suppose we want to remove the Ethers.js which is on 2nd index , so lets remove it.

	index := 2 //index value of the element which we want to remove.

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

	// So lets understand how we are deleting the element we want to delete.

	/*We are using append method to delete the element , yes you heard it right instead of adding a value
	we are using append method to delete a value , Lets see how to use it in this way.

	append(courses[:index]) with this line we are saying print the slice till the index value which is 2
	in which index 2 will not be included , we have learnt about this above :index means start from 0 index
	and go till 2nd index(as we have stored 2 in index) and 2nd index value will not be included

	and then with comma we are giving the value which we want to append in the slicing for this
	we have written coursed[index+1]... which means start the slicing from the next index (3 in this case)
	as we want to delete the 2nd index value so this is how we are ignoring it.

	So we are appending the slice from index 3 by writing (index + 1) and after that we have to use ... method
	otherwise we will get an error , we will read about this(...) in detail in later sessions.

	So combining all this and making a new slicing which is ignoring 2nd index value.

	First making a slice which is going till 2nd index but not including it and after that appending
	another slicing which is starting from 3rd index as we dont want to include it and combining all this
	we are getting a new slicing which doesnt include 2index element.

	Thats how we delete an element from a slicing , i know this really is a lengthy process
	but this is how we have to move in this case.*/

}
