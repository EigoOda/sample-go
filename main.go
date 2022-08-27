package main

import (
    "errors"
    "fmt"
    "os"
    "strconv"
)

func isOdd(num int) (bool, error) {
    if num == 0 {
        return false, errors.New("can't specify zero")
    }
    return num%2 == 1, nil
}

func main(){
    fmt.Println("START!!!")

    var value string
    fmt.Scan(&value)

    num, err := strconv.Atoi(value)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    result, err := isOdd(num)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Printf(fmt.Sprintf("IsOdd? %t", result))
}
