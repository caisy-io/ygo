// main.go

package main

/*
#cgo LDFLAGS: -L../ygo/target/release -lygo
#include "../ygo/target/ygo.h"
*/
import "C"

//export HelloWorld
func HelloWorld() {
	C.hello_world()
}

func main() {
	HelloWorld()
}
