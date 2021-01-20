package main

import "fmt"

var (
	a = 10
	b = 4
	c = 70
	d = 8
	e = 10
)

func main()  {

	add  := a+d
	sub  := a-d
	mul := b*d
	div  :=	c/b
	mod  :=	c%a
	fmt.Println(add)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)
	fmt.Println(mod)
}