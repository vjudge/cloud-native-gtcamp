package main

import "fmt"

func main ()  {
    str := [...]string{"I", "am", "stupid", "and", "weak"}
    for i, val := range str {
        //fmt.Printf("%d: %s\n", i, val)
        if val == "stupid" {
            str[i] = "smart"
        }
        if val == "weak" {
            str[i] = "strong"
        }
    }
    for _, val := range str {
        fmt.Printf("%s ", val)
    }
}
