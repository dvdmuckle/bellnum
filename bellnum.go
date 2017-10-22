package main

import "fmt"
import "os"
import "strconv"

func Bellnum(bellIn int) (belltotal int) {
	if bellIn <= 1 {
		return 1
	}
	return 2
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input a number")
		return
	}
	bellIn, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	bellOut := Bellnum(bellIn)
	fmt.Printf("The Bell number of %d is %d\n", bellIn, bellOut)
}
