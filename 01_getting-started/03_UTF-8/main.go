package main

import "fmt"

func main() {
	/* This line is incorrect and will only print up to lower-case y
	   for i := 60; i < 122; i++
	   122 needs to be 123 or operator could be <= 
	*/
	for i := 60; i <= 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
