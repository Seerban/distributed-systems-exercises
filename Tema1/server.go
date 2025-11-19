package main

import (
	"bufio"
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
	"tema1/config"
)

func main() {
	fmt.Println(config.Hello)

	ln, err := net.Listen("tcp", ":9000")
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

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Client ", name, " Conectat")
	conn.Write([]byte("Conectat\n"))

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(name, " a pierdut conexiunea.")
			return
		}

		// afirmatie primit req
		fmt.Println("Client ", name, " a facut request cu datele: ", msg)
		conn.Write([]byte("primit request.\n"))

		// ---- procesare request ----
		args := strings.Split(msg, " ")

		raspuns := rezolva(args)

		// raspuns catre client
		conn.Write([]byte(raspuns + "\n"))
	}
}

func rezolva(args []string) string {
	if len(args) == 0 {
		return "Exemplu call: ex1 args..."
	}

	args[len(args)-1] = strings.TrimSpace(args[len(args)-1])

	switch args[0] {
	case "ex1":
		return ex1(args)
	case "ex2":
		return ex2(args)
	case "ex3":
		return ex3(args)
	case "ex5":
		return ex5(args)
	case "ex11":
		return ex11(args)
	}

	return "ex necunoscut"
}

func ex1(args []string) string {
	length := len(args) - 1

	for i := 1; i < len(args); i++ {
		if len(args[i]) != length {
			fmt.Println(args[i], length)
			return "Parametru incorect."
		}
	}

	res := ""

	for i := 0; i < length; i++ {
		for j := 1; j < len(args); j++ {
			res += string(args[j][i])
		}
		res += " "
	}

	return res
}

func ex2(args []string) string {
	count := 0
	for i := 1; i < len(args); i++ {
		num := ""

		for _, c := range args[i] {
			if c > '0' && c <= '9' {
				num += string(c)
			}
		}

		if len(num) == 0 {
			continue
		}

		n, _ := strconv.Atoi(num)

		sqrt := int(math.Sqrt(float64(n)))
		if sqrt*sqrt == n {
			count += 1
		}
	}
	return strconv.Itoa(count) + " patrate perfecte"
}

func mirror(n string) string {
	res := ""
	for i := len(n) - 1; i >= 0; i-- {
		res += string(n[i])
	}
	return res
}

func ex3(args []string) string {
	sum := 0
	for i := 1; i < len(args); i++ {
		n, _ := strconv.Atoi(mirror(args[i]))
		sum += n
	}
	return strconv.Itoa(sum) + " suma totala"
}

func ex5(args []string) string {
	res := ""
	for i := 1; i < len(args); i++ {

		// verifica binary
		bin := 1
		for _, c := range args[i] {
			if c != '0' && c != '1' {
				bin = 0
			}
		}
		if bin == 0 {
			continue
		}

		num, _ := strconv.ParseInt(args[i], 2, 64)
		fmt.Println(num)
		res += strconv.Itoa(int(num)) + ", "
	}
	return res
}

func shift(s string) string {
	return s[1:len(s)] + string(s[0])
}

func ex11(args []string) string {
	sum := 0
	for i := 1; i < len(args); i++ {
		temp, _ := strconv.Atoi(shift(args[i]))
		sum += temp
	}
	return strconv.Itoa(sum)
}
