package scopes
import 
	"fmt"

func Funcscope() {
//first method of declare and define
	var a int
	a = 10
	fmt.Println("A's value is", a)

//second method of declare and define

	var y string = "Trine University"
	fmt.Println("My university", y)
      //Third method of declare and define
	c:= true

	fmt.Print("That is", c) //spaces issue.

 }


func Funcexample(){

var x int = 2
var y int = 4
fmt.Println("x is", x)
fmt.Println("y is", y)
 }