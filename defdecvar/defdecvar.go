package defdecvar

import "fmt"

func Defdecvar(){
// here we defined and declared variable interger using hard code process
	age := 10
	name := "ravi"
    fmt.Println("hi everyone,my age is",age)
    fmt.Println("hi everyone,my name is",name)
	
    age = 15
	name = "teja"
	fmt.Println("hi all my new age is",age)
	fmt.Println("hi everyone,my name is",name)



	  // he we defining and declaring variables in second menthod

	  var agenew int
	  var named string
	  agenew = 18
	  named = "raja"
	  fmt.Println("i turned",agenew)
	  fmt.Println("hi everyone,i named as",named)

	  agenew = 21
	  named = "abcd"
	  fmt.Println("i turned",agenew)
	  fmt.Println("hi everyone,i named as",named)

	  //here we defining and declaring variable using third method

	  var age1 int = 11
	  var name1 string = "rajesh"
	  fmt.Println(age1)
	  fmt.Println(name1)

	  // here am defining and declaring multiple variable as a part of practice

	   var height,weight int = 6 , 70
	   var surname , lastname string = "ravi" , "khandavilli"
	  fmt.Println("my height n weight", height , weight )
	  fmt.Println("hi everyone,i named as", surname ,lastname)

}	
