package main
import (
	"fmt" 
	//"strings"
	"os"
)
func main() {
	//var s, sep string
	// //sep="-"
	for idx, arg := range os.Args[1:] {
		//s += sep + arg
		//sep = "-"
		fmt.Println(idx + 1)
		fmt.Println(arg)
	}
	//fmt.Println(s)
	//fmt.Println(os.Args[0])
}