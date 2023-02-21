package main

 import (
	f "fmt"
	"time"
)

 func main(){

	start := time.Now()
	end := time.Now()
	delta := end.Sub(start)
	f.Printf("longCalculation took this amount of time: %s\n", delta)

 }