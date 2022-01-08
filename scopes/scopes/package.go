package scopes
import  "fmt"
//Package function

	
	var y string = "Trine University" //This is package level variable

	var c bool = true //This is a package level variable
	var GlobalVar int = 4 // This is a global level variable


func Packagescope() {

	var a int   //This is a local/block/function level variable
	a = 10

	fmt.Println("A's value is", a)


	
	fmt.Println("My university", y)

	

	fmt.Print("That is ", c) 

 }


func Packageexample()  {
	var x int = 2
	fmt.Println("\nx is", x)

}

func Datatypesexample(){
//data range for int8 : -(2^(8-1)) to (2^(8-1)) = -(2^7) to (2^7 - 1) = -128 to 127
	var a int8

	a = 26


	fmt.Println("a is", a)

	// data range for int16 : -(2^(16-1)) to (2^(16-1)-1) = -(2^15) to (2^15-1) = -32768 to 32767

	var b int16

     b = 125 // This (:=) cannot be used outside the function

fmt.Println ("x is", b)


// Declaring the variable in "C" string "data type"
var c string
c = "Trine" 

fmt.Println ("Iam from", c)

// Declaring the variable "D" in boolean

var D bool

D = true

fmt.Println ("That is", D)

//data range for int32 = -(2^(32-1)) to (2^(32-1)-1)  = -(2^31) to 2^31 - 1 = -(2147483648) to 2147483647

var e int32
e = 28

fmt.Println("He is", e)

//data range for int64 = -(2^(64-1) to 2^(64-1)-1 = -(2^63) to 2^63 - 1 = -(2^63) to 2^63 -1 = -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
var f int64

f = 30

fmt.Println("He is", f)

//This is an example for Local, Package and Global level scopes..



}




