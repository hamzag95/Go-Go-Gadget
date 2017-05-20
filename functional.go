package main

import (

    "fmt"
)

//compare function takes names and and a compare functions
func battle(x int, y int, f func(int, int) int) int  {
    
    return f(x, y)

}


func main() {

    m := func(x, y int) int {
        return(x*y)
    }

    a := func(x, y int) int {
        return(x+y)
    }

    fmt.Println(battle(2, 3, a))
    fmt.Println(battle(2, 3, m))
    fmt.Println(battle(2, 3, func(x int, y int) int {
        return(x-y)
    }))

}
