package main

import (
	"fmt"
	"server/structs"
)

var dataClient = structs.Client{}

func main() {
	fmt.Println("введите имя")
	fmt.Scanln(&dataClient.Name)
	fmt.Println(dataClient.Name)
}
