//most of this is from the Tour of Go

package main

import (
  "fmt"
  "math"

  )

func main() {
  //while loop
  i := 0
  for i < 10 {
    i++
  }

//declare var as a function
  f := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)

  }

    fmt.Println(compute(f))
}

//example of short statement conditional from go docs
func cond(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v

  }

  return lim

}

//function that takes a function as parameter
func compute(fn func(float64, float64) float64) float64  {
        return fn(3,4)

}
