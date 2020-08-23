package main

import "fmt"
import "reflect"
import "strings"

func main() {
    type Filter struct {
        CategoryID int
        Condition string `condition:"WHERE"`
    }
    filter := Filter{
        CategoryID: 1,
        Condition: "weekend",
    }

    reflects := reflect.ValueOf(&filter).Elem()
    condTag := reflects.Type().Field(1).Tag.Get("condition")
    fmt.Println(strings.Split(condTag, ","))
    fmt.Println(reflects.Field(1).Interface())
}
