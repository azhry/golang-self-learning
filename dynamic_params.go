package main

import (
	"fmt"
)

func main() {
	DynamicParams(2, "str1", "str2")
}

func DynamicParams(requiredID int, params ...string) {
	fmt.Println(fmt.Sprintf("ID. %d value: %v", requiredID, params))
}
