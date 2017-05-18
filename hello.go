package main

//Go doesnt compile if you have unused imports
import (
    "fmt"

)

/* have to put var when declaring variables
* outside of function parameters
*/
var c, python bool

//but we can omit types if we initliaze
//keeping is type and initializing also fine
var numA, numB, cow = 1, 2, "moo"

func main() {
    var hi string = "hello"

    //:= only used for new variables, and inside functions
    //and the type becomes that of which its first declared to
    k:= "yo"
    k = "world"


    fmt.Println(k)
    fmt.Println(multDiv(8))
    fmt.Println(hi, k, "\nhere are numbers", numA, numB, "\nthe cow goes", cow)

}

/*type is after variable name
* can declare local variables on top of function
* as well as return more than one value
*/
func multDiv(a int) (x, y int){
  x = a/2
  y = a*2
  return

}
