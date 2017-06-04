package main

//Go doesnt compile if you have unused imports
import (

)

/* have to put var when declaring variables
* outside of function parameters
*/
var c, python bool

/* but we can omit types if we initliaze
* keeping is type and initializing also fine
*/
var numA, numB, cow = 1, 2, "moo"

//or we can factor variables
var (
  tf  = true
  num = 15
  z complex128 = -5+4i

)

//constant
const e = 2.7

//factor version of const 
const (
  pi = 3.14
  tau = 6.28

)

func main() {
    var hi string = "hello"
    //:= only used for new variables, and inside functions
    //and the type becomes that of which its first declared to
    k:= "yo"
    k = "rewrite"


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

/*This is doesnt compile
* b/c y:= isnt new, its basically a local variable
* defined at the top in the return portion
func convert() (y float64) {
  x := 2
  y := float64(x)
  return
}
*/

//converting data types
func convert() (y float64) {
  x := 2
  y = float64(x)
  return
}
