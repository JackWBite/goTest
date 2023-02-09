package base

import (
	"fmt"
	"runtime"
)

// first go func
func SayHi() {
	fmt.Println("hi, welcome to go world !")
}

func (*Human) SayHi() {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println("i'm human")
	}
}

func (*Student) SayHi() {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println("i'm student")
	}
}
