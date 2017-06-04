package main

import "fmt"


var a, b, c int = 7, 5, 3


func printA(num int) {

    a-= num
    for i := 0; i < a; i++ {
        fmt.Print("o")
    }
    fmt.Println()
}

func printB(num int) {

    b-= num
    for i := 0; i < num; i++ {
        fmt.Print("o")
    }
     fmt.Println()
}

func printC(num int) {

    c -= num
    for i := 0; i < num; i++ {
        fmt.Print("o")
    }
     fmt.Println()

}

func initialize () {
    fmt.Println("Row a: ooooooo")
    fmt.Println("Row b: ooooo")
    fmt.Println("Row c: ooo")


}

func main() {

    printA(2)
    printB(3)
    printC(1)



}
