package main

import (
	f "fmt"
)

type Sleeper interface{

	Sleep()

}


type Dog struct{
	Name string
}

func (d Dog) Sleep() {
	f.Println(d.Name)
}




type Cat struct{
	Name string
}


func (c Cat) Sleep() {
	f.Println(c.Name)
}


func AnimaSleep(s Sleeper) {
	s.Sleep()
}



func main(){
	var s string
	dog := 
}