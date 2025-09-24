package main

import (
	"bufio" //to read the data
	"fmt"   //to print out
	"os"    //to open the file
)

func main() {
	fmt.Println("yo yo yo")
	file, f_err := os.Open("example.txt")
	if f_err != nil {
		panic(f_err)
	}
	scanner := bufio.NewScanner(file) //buffer where we can read the file into
	for scanner.Scan() {
		//this for loop is going to scan the file line by line
		line := scanner.Text() //going to populate the current line
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
