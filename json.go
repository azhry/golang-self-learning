package main

import (
	"fmt"
	"encoding/json"
)

func main() {
        type CameraAngle struct {
		Angle string `json:"angle"`
	}
	var angle CameraAngle
	if err := json.Unmarshal([]byte(`{"angle": "https://google.com"}`), &angle); err != nil {
		fmt.Println(err)
	}
	fmt.Println(angle.Angle)
	
	jsonStr, _ := json.Marshal(angle)
	fmt.Println(string(jsonStr))
}
