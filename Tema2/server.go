package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"tema1/config"
)

type KeyValue struct {
	Key   string
	Value int
}

func main() {
	ln, err := net.Listen("tcp", config.Port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go connect(conn)
	}
}

func connect(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// citeste nume si raspunde
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // scoate endline si spatii
	fmt.Println(name, " s-a conectat.")
	conn.Write([]byte("Conectat.\n"))

	for {
		var lines []string

		// citeste pana ""
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(name, "a pierdut conexiunea.")
				return
			}

			line = strings.TrimSpace(line)

			// opreste citirea dupa o linie blank
			if line == "" {
				break
			}

			lines = append(lines, line)
		}

		fmt.Println(name, " request: ", lines)

		rsp := rezolva(lines)

		fmt.Println("trimite raspuns " + rsp)

		conn.Write([]byte(rsp + "\n"))
	}
}

func rezolva(args []string) string {
	if len(args) == 0 {
		return "Exemplu call: \n ex1 \n args0 \n args1..."
	}

	cmd := args[0]

	switch cmd {
	case "ex1":
		return ex1(args)
	case "ex6":
		return ex6(args)
	case "ex9":
		return ex9(args)
	}

	return args[0] + " exercitiu necunoscut"
}

func Map(document string, valuefunc func(string) int) []KeyValue {
	words := strings.Fields(document)
	acc := make(map[string]int) // acumulat total

	// total la fiecare cheie
	for _, word := range words {
		acc[word] += valuefunc(word)
	}

	keyValues := make([]KeyValue, 0, len(acc))
	for k, v := range acc {
		keyValues = append(keyValues, KeyValue{Key: k, Value: v})
	}

	return keyValues
}

func Reduce(keyValues []KeyValue) int {
	sum := 0
	for _, kv := range keyValues {
		sum += kv.Value
	}
	return sum
}

// 2|n Vocale si 3|n Consoane
func ex1Func(s string) int {
	vocs := 0
	cons := 0
	vocale := "aeiou"

	for _, char := range s {
		if strings.ContainsRune(vocale, char) {
			vocs++
		} else {
			cons++
		}
	}

	if vocs%2 == 0 && cons%3 == 0 {
		return 1
	}
	return 0
}

func ex1(args []string) string {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += Reduce(Map(args[i], ex1Func))
	}
	avg := float64(sum) / float64(len(args)-1)
	return fmt.Sprintf("%.2f", avg)
}

func ex6Func(s string) int {
	if len(s) == 0 {
		return 0
	}

	first := s[0]
	last := s[len(s)-1]

	if first >= 'A' && first <= 'Z' && last >= 'A' && last <= 'Z' {
		return 1
	}
	return 0
}

func ex6(args []string) string {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += Reduce(Map(args[i], ex6Func))
	}
	avg := float64(sum) / float64(len(args)-1)
	return fmt.Sprintf("%.2f", avg)
}

// n cuvinte rimeaza cu alt cuvant din arr
func ex9Map(document string) []KeyValue {
	words := strings.Fields(strings.ToLower(document))
	acc := make(map[string]int) // acumulat total

	// total la fiecare cheie
	for _, word := range words {
		key := word
		if len(word) >= 2 {
			key = word[len(word)-3:] // doar ultimele 3 litere sunt pt rima
		}
		acc[key] += 1
	}
	for k, v := range acc {
		acc[k] = v - 1 // ca sa inceapa de la 0 rime nu 1
	}

	keyValues := make([]KeyValue, 0, len(acc))
	for k, v := range acc {
		keyValues = append(keyValues, KeyValue{Key: k, Value: v})
	}

	return keyValues
}

func ex9(args []string) string {
	sum := 0
	for i := 1; i < len(args); i++ {
		fmt.Println(ex9Map(args[i]))
		sum += Reduce(ex9Map(args[i]))
	}
	avg := float64(sum) / float64(len(args)-1)
	return fmt.Sprintf("%.2f", avg)
}
