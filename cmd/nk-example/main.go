package main

import (
	"C"
	"log"
	
	"demo/subpack"


)

func main() {
	d:=subpack.NewDemo()
	d.Run()
}


func onError(code int32, msg string) {
	log.Printf("[glfw ERR]: error %d: %s", code, msg)
}
