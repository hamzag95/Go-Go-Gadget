package main

//Go doesnt compile if you have unused imports
import (
    "fmt"

)


func main() {
    fmt.Println(multDiv(8))

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
