package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Cafea struct {
	Nume string `xml:"nume"`
}

// https://gobyexample.com/xml
func (c Cafea) String() string {
	return fmt.Sprintf(c.Nume)
}

func mainb() {
	cafea := Cafea{Nume: "hello"}

	out, _ := xml.MarshalIndent(cafea, " ", " ")
	fmt.Println(string(out))

	err := os.WriteFile("cafea.xml", out, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
