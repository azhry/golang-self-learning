package main

import "fmt"
import "reflect"
import "strings"

func main() {
    var filter struct {
        CategoryID int
        Condition string `condition:"WHERE"`
    }

    reflects := reflect.ValueOf(&filter).Elem()
    condTag := reflects.Type().Field(1).Tag.Get("condition")
    fmt.Println(strings.Split(condTag, ","))
}
