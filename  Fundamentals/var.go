package main

import "fmt"

func main()  {

	tk1,tk2,tk3 := "majid",214,32.5
	tk4,tk5,tk6 := 324.7,"hi",22

	fmt.Printf("The value of tk1 is : %s\n",tk1)
	fmt.Printf("The type of tk1 is : %T\n", tk1)

	fmt.Printf("The value of tk2 is : %d\n",tk2)
	fmt.Printf("The type of tk2 is : %T\n", tk2)

	fmt.Printf("The value of tk3 is : %f\n",tk3)
	fmt.Printf("The type of tk3 is : %T\n", tk3)

	fmt.Printf("The value of tk4 is : %f\n",tk4)
	fmt.Printf("The type of tk4 is : %T\n", tk4)

	fmt.Printf("The value of tk5 is : %s\n",tk5)
	fmt.Printf("The type of tk5 is : %T\n", tk5)

	fmt.Printf("The value of tk6 is : %d\n",tk6)
	fmt.Printf("The type of tk6 is : %T\n", tk6)
}
