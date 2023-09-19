package main

import "data-structs/list"

func main() {
    var numbers = list.NewList()
    numbers.Add(1)
    numbers.Add(3)
    numbers.Add(5)
    list.PrintNodes(*numbers)

}
