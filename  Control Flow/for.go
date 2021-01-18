package main

import "fmt"

func main()  {

	for i:=0;i<10;i++{
		if i == 4{
			continue
		}
		if i == 9{
			break
		}
		fmt.Println(i)

	}

	x : for j:=0;j<3;j++{
		if j==3{
			goto x
		}
		fmt.Println(j)
	}
}
