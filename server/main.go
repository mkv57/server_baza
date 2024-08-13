package main

import (
	"fmt"
	"server/structs"
)

func main() {

	var dataClient = structs.Client{}

	fmt.Println("введите ваше имя")
	fmt.Scanln(&dataClient.Name)
	fmt.Println("введите ваш возраст")
	fmt.Scanln(&dataClient.Age)
	fmt.Println("введите вашу з/п в тыс.руб.")
	fmt.Scanln(&dataClient.Wage)
	//fmt.Println(dataClient)
	fmt.Println("Уважаемый ", dataClient.Name)
	structs.LoanCheck(dataClient)
}
