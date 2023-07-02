package main

import "hex_practicce/application"

func main() {
	a := application.Run()
	a.SignOn(nil, nil)
	print("test")
}
