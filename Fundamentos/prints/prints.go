package main

import "fmt"

func main() {
	fmt.Print("nesta ")
	fmt.Print("linha.")
	fmt.Println(" ")
	fmt.Println("nova linha")

	x := 3.141516
	//fmt,println("O valor de x é" + x)

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)
	fmt.Printf("O valor de x é %.3f", x)
}
