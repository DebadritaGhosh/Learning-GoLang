package main

import "fmt"

// --------[Public]-------------------------------------------------
const LoginToken string = "asdahkdb" // 1st letter of the variable should be capital

func main() {

	// -----------[String]----------------------------------------------
	var username string = "Debadrita Ghosh"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)
	// -------------[Boolean]--------------------------------------------
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)
	// -------------[]--------------------------------------------
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)
	// ---------------------------------------------------------
	var val int = 255
	fmt.Println(val)
	fmt.Printf("Variable is of type : %T \n", val)
	// ---------------------------------------------------------
	var smallFloat float32 = 255.45545511125353
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)
	// --------[Default Values And Some Aliases]-------------------------------------------------
	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)
	// --------[Implicit Type]-------------------------------------------------
	var website = "debadritaghosh.tech"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)
	// --------[No Var Style]-------------------------------------------------
	numberOfUser := 30000 //Inside the metod we can use := operator
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type : %T \n", numberOfUser)
	// --------[Using public Variable]-------------------------------------------------
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}
