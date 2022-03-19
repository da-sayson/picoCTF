package main

import (
	"fmt"
	"os"
	"strings"
)

const aToM string = "abcdefghijklm"
const nToZ string = "nopqrstuvwxyz"
const upperAToM string = "ABCDEFGHIJKLM"
const upperNToZ string = "NOPQRSTUVWXYZ"
	
func main() {
	args := os.Args
	if (len(args) < 2) {
		fmt.Println("Missing argument <rot13>")
		fmt.Println("Syntax: mod <rot13>")
		return
	}

	var inString string = os.Args[1]
	var outStringBuilder strings.Builder
	var currLet string
	for idx := 0; idx < len(inString); idx++ {
		currLet = string(inString[idx])
		if strings.Contains(aToM, currLet) {
			inIdx := strings.Index(aToM, currLet)
			outStringBuilder.WriteString(string(nToZ[inIdx]))
		} else if strings.Contains(upperAToM, currLet) {
			inIdx := strings.Index(upperAToM, currLet)
			outStringBuilder.WriteString(string(upperNToZ[inIdx]))
		} else if strings.Contains(nToZ, currLet) {
			inIdx := strings.Index(nToZ, currLet)
			outStringBuilder.WriteString(string(aToM[inIdx]))
		} else if strings.Contains(upperNToZ, currLet) {
			inIdx := strings.Index(upperNToZ, currLet)
			outStringBuilder.WriteString(string(upperAToM[inIdx]))
		} else {
			outStringBuilder.WriteString(currLet)
		}
	}
	fmt.Println(outStringBuilder.String())
}