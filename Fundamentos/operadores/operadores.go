package main

import (
  "fmt"
  "math"
)
func main(){
  a := 3
  b := 2

  fmt.Println("Soma =>", a+b)
  fmt.Println("sub =>", a-b)
  fmt.Println("divisao =>", float64(a)/float64(b))
  fmt.Println("multiplicação =>", a*b)
  fmt.Println("modulo =>", a%b)

  //bitwise
  fmt.Println("E =>", a&b)

  //usando o pacote math

  c:= 3.0
  d:= 2.0

  fmt.Println("Maior=>", math.Max(float64(c), float64(d)))
  fmt.Println("Menor=>", math.Min(c,d))
  fmt.Println("Expo=>", math.Pow(c,d))
}
