package main

import "fmt"
import "os"
import "strconv"

func Bellnum(bellstart int) (belltotal int) {
	return 0
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input a number")
		return
	}
	bellstart, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	bellend := Bellnum(bellstart)
	fmt.Printf("The Bell number of %d is %d\n", bellstart, bellend)
}
