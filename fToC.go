package main
import "fmt"

func main() {
	var boilingF float64 = 212.0
	var freezingF float64 = 32.0
	var boilingC float64 = fToC(boilingF)
	var freezingC float64 = fToC(freezingF)
	fmt.Printf("%.2f %.2f\n", boilingC, freezingC)
}

func fToC (f float64) float64 {
	return (f - 32) * 5 / 9
}