package main

import "fmt"
func average(scores[]float64)float64{
  total:=0.0
  for _, v := range scores {
    total += v
  }
  return total / float64(len(scores))
}
func main(){
scores:=[]float64{99,78,87,81,90}
fmt.Println(average(scores))
}