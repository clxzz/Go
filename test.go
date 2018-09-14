package main
import "fmt"

func main() {
	var i, j = 1.3, 1.4
	point := new(int)
	*point = 2
	p := &i
	i, j = j, i                   //交换变量值
	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(increase(p))
	fmt.Println(newInt() == newInt())
	fmt.Printf("i=%f, j=%f, point=%d, %s\n", i, j, *point, medals[0])
	x, y := 12, 8
	fmt.Println(maxCommenDiviser(x, y))
	fmt.Println(fibnanci(1))
}

func increase(p *float64) float64 {
	*p++
	return *p
}

func newInt() *int {
	return new(int)
}

func maxCommenDiviser(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x 
}

func fibnanci(n int) int {
	x, y := 1, 1
	for i:=2;i<n;i++ {
		x, y = y, x+y
	}
	return y
}