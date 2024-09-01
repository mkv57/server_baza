package structs

import "fmt"

func LoanCheck(c Client) string {

	var IsAnyProduct bool = false
	k := ""

	for i := 0; i < 3; i++ {

		if c.Wage >= (Products[i]).Salary && c.Age >= (Products[i]).AgeMin && c.Age <= (Products[i].AgeMax) {
			IsAnyProduct = true
			fmt.Println("Вам доступен ")
			fmt.Println(Products[i].LoanName)
			fmt.Println("На сумму до ", Products[i].Sum, "рублей")
			fmt.Println("со ставкой ", Products[i].Rate, " % годовых")
			n1 := "rrr"
			k = k + n1
		}
	}
	if IsAnyProduct == false {

		fmt.Println("Уважаемый ", c.Name)
		fmt.Println("К сожалению у нас сегодня нет для вас предложения по займу ")
		//return ("К сожалению у нас сегодня нет для вас предложения по займу ")
		k = "Уважаемый " + c.Name + " к сожалению у нас сегодня нет для вас предложения по займу "

	}
	return k //("К сожалению у нас сегодня нет для вас предложения по займу ")
}
