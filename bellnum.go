package main

import "fmt"
import "os"
import "strconv"

// Wraper function for real function
func Bellnum(bellIn int) (belltotal int) {
	//Handle the base case to avoid any unecessary computation
	if bellIn <= 1 {
		return 1
	}
	var bellSlice [][]int
	bellSlice[0][0] = 1
	bellFinal := bellcurse(bellIn, bellSlice)
	return bellFinal[len(bellFinal)-1][len(bellFinal)-1]
}

// Actual function that recursively goes through Bell triangle
func bellcurse(bellIn int, bellArrIn [][]int) (bellArrOut [][]int) {
	return nil
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
