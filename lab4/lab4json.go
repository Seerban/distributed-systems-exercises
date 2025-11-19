package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Fructe struct {
	Mere     float32 `json:"mere"`
	Pere     float32 `json:"pere"`
	Piersici float32 `json:"piersici"`
	Capsuni  float32 `json:"capsuni"`
}

func mainc() {
	fmt.Println("hello world")

	jsonb, err := os.ReadFile("cantitati_fructe.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(jsonb)

	var fructe Fructe

	json.Unmarshal(jsonb, &fructe)

	fructe.Mere *= 2
	fructe.Pere -= 10

	fmt.Println(fructe)

	fructe2 := Fructe{
		Mere:     10.5,
		Pere:     9,
		Piersici: 5,
		Capsuni:  15,
	}
	fructejson, _ := json.Marshal(fructe2)

	err = os.WriteFile("fructe.json", fructejson, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
