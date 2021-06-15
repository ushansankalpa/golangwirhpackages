package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Time        string `json:"time"`
	Infromation Info   `json:"info"`
}

type Info struct {
	Description string `json:"description"`
}

func main() {

	jsonString := `{"name": "battery sensor", "capacity": 40 , "time": "2019-01-21T19:07:283" , "infor":{
		"desc":"a sensor reading"
	}}`

	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)
}
