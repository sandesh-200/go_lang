package main

import "fmt"

const LoginToken string = "sjdgfjsdhjgfjhsdvb" //public

var jwtToken = 300000

func main() {
	var username string = "Sandesh"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloatVal float64 = 255.455535325235434
	fmt.Println(smallFloatVal)
	fmt.Printf("variable is of type: %T \n", smallFloatVal)

	//default  values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

}
