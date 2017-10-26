package main

import "fmt"
import "os"
import "strconv"
import "math"

func factorial(x float64) float64 {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// Wraper function for real function
func Bellnum(bellIn int) float64 {
	//Handle the base case to avoid any unecessary computation
	if bellIn <= 1 {
		return 1
	}
	//var bellSlice [][]int
	//bellSlice[0][0] = 1
	//bellFinal := bellcurse(bellIn, bellSlice)
	//return bellFinal[len(bellFinal)-1][len(bellFinal)-1]
	return bellmath(float64(bellIn))
}

//Double for loop version of function
func bellfor(bellIn int) int {
	bellArr := make([][]int, bellIn)
	for n, _ := range bellArr {
		bellArr[n+1] = make([]int, n+1)
	}
	bellArr[0][0] = 1
	for i, row := range bellArr {
		for j, cell := range row {
			fmt.Printf("%d ", cell)
			//If we're looking at the first element of a row
			if j == 0 {
				if i == 0 {
					//If we're looking at the top of the triangle
					goto loopback
				}
				bellArr[i][j] = bellArr[i-1][j]
			}
			fmt.Println("")
			bellArr[i][j] = bellArr[i-1][j-1] + bellArr[i][j-1]
		}
	loopback:
	}
	return bellArr[bellIn-1][bellIn-1]
}

// Actual function that recursively goes through Bell triangle
func bellcurse(bellIn int, bellArrIn [][]int) (bellArrOut [][]int) {
	return nil
}

//Mathematical implementation of function
func bellmath(bellIn float64) float64 {
	var bn float64
	for i := 1; i <= 2*int(bellIn); i++ {
		bn += math.Pow(float64(i), bellIn) / factorial(float64(i))
	}
	return math.Ceil(math.Pow(math.E, -1) * bn)
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
	fmt.Printf("The Bell number of %d is %d\n", bellIn, int(bellOut))
}
