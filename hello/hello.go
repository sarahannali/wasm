package main

import (
    "fmt"
    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
		var a = [9][9]int{
			{3, 0, 6, 5, 0, 8, 4, 0, 0}, 
			{5, 2, 0, 0, 0, 0, 0, 0, 0}, 
			{0, 8, 7, 0, 0, 0, 0, 3, 1}, 
			{0, 0, 3, 0, 1, 0, 0, 8, 0}, 
			{9, 0, 0, 8, 6, 3, 0, 0, 5}, 
			{0, 5, 0, 0, 9, 0, 6, 0, 0}, 
			{1, 3, 0, 0, 0, 0, 2, 5, 0}, 
			{0, 0, 0, 0, 0, 0, 0, 7, 4}, 
			{0, 0, 5, 2, 0, 6, 3, 0, 0} }
    message := greetings.Solve(a)
    fmt.Println(message)
}