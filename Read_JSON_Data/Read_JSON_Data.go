//Continuing our API Development Journey.

//In this session we gonna learn how we can read json data in go.

package main

import (
	"encoding/json"
	"fmt"
)

//Again making a struct as exactly as we made in our previous session in which we learnt of creating json data.

//Dont just copy paste the code cmon write it again.

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"learning platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("In this session we gonna see how can we read json data in go")

	ReadJsonData() //calling the function inside main function so that we can execute it.

}

func ReadJsonData() {

	/*Now we are storing the json data in forms of bytes as we know that when we get data in respons it is
	in form of bytes so we are storing it inside slice of byte this is because if we want to store multiple
	data in it and this will also be written in back ticks like we have done below*/

	jsonData := []byte(`
	
	{
		"coursename": "Solidity",
		"price": 299,
		"learning platform": "linkedin/theyashjain17",
		"tags": ["web3","blockchain"]
	}
	
`)

	var myCourse course //now we are making a variable of type course which is the struct we made above.

	checkVaildData := json.Valid(jsonData) //now with the help of json module and its property Valid we are checking whether the json we are getting(we pasted above) is valid or not , this will give us a bool means true or false therefore we dont have to store it inside an err variable.

	if checkVaildData { //checking whether the json is valid or not

		fmt.Println("Json is valid") //if success print this

		json.Unmarshal(jsonData, &myCourse) //when we were creating json we were using MarshalIndent to create the json , so to read the json we will UnMarshal and inside it provide the json data we want to read and 2nd thing is the interface which is struct in our case by doing this we will get the data in this struct form like we did in case of creating json and we are using & for the pointer as we dont want to use the copy of the struct we want to use exact same thing therefore using pointer and passing the reference of the data we created.

		fmt.Printf("%#v \n", myCourse) //for getting this data in readable form there is a special expression which is %#v so kindly use this.

		// fmt.Printf("%v \n", myCourse) this will also give you information but not in proper way netiher in detailed manner thats why it would be better to use %#v expression and not the %v in this case.

	} else {

		fmt.Println("Json is not valid") //if failure print this

	}

	//We are trying the same thing with the help of map and why are we doing it we will get to know it below.

	var onlineData map[string]interface{} // another way of creating the map as we can make it with the help of make keyword but here we are making it directly and we are mapping the string with an interface and we are doing this because as we dont know what will be the other value it could be stirng int or any other thing therefore doing this.

	json.Unmarshal(jsonData, &onlineData) //now again doing the same thing which we did above

	fmt.Printf("%#v \n", onlineData) //for getting this data in readable form there is a special expression which is %#v so kindly use this.

	/*So the benefit of doing the same with a mapping is that now we can use a loop on this
	and from this we can sort data more than ever and it will be easy to read in comaparison
	of previous data.*/

	// So applying the loop with the help of range keyword nothing new.

	for key, value := range onlineData {

		fmt.Printf("The Key is %v And The Value is %v And The Type is %T \n", key, value, value)

	}

}

// This was all for this session.

// Stay tuned for more updates.