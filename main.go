package main

import (
	"fmt"
	"time"
)

func sayHello(str string){
	for i:=0;i<10;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(str)
	}
}
func main() {
		go sayHello("World!")
		sayHello("Hello,")

}
