package  main

import "fmt"

func main(){
	a:=[...]int{1,2,3,4,5,6,7,8,9}
	b:=a[2:5]
	c :=b[2:5]
	fmt.Println(c)
}
