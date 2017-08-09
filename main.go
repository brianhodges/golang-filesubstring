package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

var xml = ""

func check(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}

func main() {
	fmt.Println()
	if len(os.Args) <= 1 {
		fmt.Println("You must enter parameter(s).")
		fmt.Println("Ex. go run main.go C:\\path\\to\\xml 1 1000")
		os.Exit(1)
	}
	
	dir := os.Args[1]
	xmlFile, err := os.Open(dir)
	check(err)
	defer xmlFile.Close()
	
	scanner := bufio.NewScanner(xmlFile)
    for scanner.Scan() {
		xml = xml + strings.TrimSpace(scanner.Text())
	}
	
	if len(os.Args) != 4 {
		fmt.Println(xml)
		os.Exit(1)
	}
	
	start, err := strconv.Atoi(os.Args[2])
	stop, err := strconv.Atoi(os.Args[3])
	check(err)
	
	if start >= stop {
		fmt.Println("Stop parameter must be greater than start paramter.")
		os.Exit(1)
	}
	fmt.Println(xml[start:stop])
}