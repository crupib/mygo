package main

import "fmt"
import karatsuba "karatsuba"
import big "math/big"

func main() {
  a :=new(big.Int)  
  b :=new(big.Int)
  a.SetString("20000 000000000000000",10)
  b.SetString("200000000000000000000",10)
  pa := a
  pb := b
  fmt.Println(karatsuba.K_multiply(pa,pb))
}
