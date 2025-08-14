package main

import "fmt"

type alamType func(num int) //signature
func printValue(p alamType) {
	p(12)
}
func main() {

	fn := func(num int) {
		fmt.Println(num)
	}
	printValue(fn)

}
