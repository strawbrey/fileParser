package main

import (
		"fmt"
		"bufio"
		"os"
		"bytes"
		"strconv"
	)

func FindLength(t string) int {	
	return len(t)
}

func FindLongestLine(s []byte) string { 

	maxLength := 0
	var lineLength int
	var longestLine string
	b := bytes.NewReader(s)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
		fmt.Println("length: " + strconv.Itoa(len(scanner.Text())))
		lineLength = len(scanner.Text())
		if lineLength > maxLength {
			maxLength = lineLength 
			longestLine = scanner.Text()
		}
		fmt.Println("max: " + strconv.Itoa(maxLength))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return longestLine
}

