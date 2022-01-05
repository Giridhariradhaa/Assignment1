package scopes
import 
	"fmt"
//Package function

	
	var y string = "Trine University" //This is package level variable

	c:= true //This is a package level variable
	var GlobalVar int = 4 // This is a global level variable


func Packagescope() {

	var a int   //This is a local/block/function level variable
	a = 10

	fmt.Println("A's value is", a)


	
	fmt.Println("My university", y)

	

	fmt.Print("That is", c) 

 }


func Packageexample(){

	var x int = 2
	

fmt.Println("x is", x)

 }

 //