package main

import (
	"bufio" //to read the data
	"fmt"   //to print out
	"os"    //to open the file
)

func main() {
	file, f_err := os.Open("example.txt")
	if f_err != nil {
		panic(f_err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// if len(os.Args)<2{
	// 	fmt.Print("Error: no input text")
	// 	os.Exit(1)
	// }
}

//plan:
//1. flag reading (flags: -w, -p, -l, -help)
//2. program takes a text file -> create readfile function
//3. program can also take input from stdin
//4.
