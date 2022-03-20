package main

import (
	"fmt"
	"os"
	"io"
	"strings"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Incorrect Usage")
		fmt.Println("Usage: %s <numListFile>", os.Args[0])
		return
	}
	
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	num := make([]byte, 4) // 3 characters plus newline
	var sb strings.Builder
	for {
		_, err = io.ReadFull(f, num)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		_, err = sb.WriteString(string(num))
		if err != nil {
			fmt.Println(err)
		}
	}
	str := strings.ReplaceAll(sb.String(), "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	parts := strings.Split(str, " ")
	sb.Reset()
	for i := 0; i < len(parts); i++ {
		if parts[i] == "" { continue }

		num, err := strconv.Atoi(parts[i])
		character := string(num)

		_, err = sb.WriteString(character)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Printf("%q\n", sb.String())
}