package domain

import (
	"strconv"
)

type Clients interface {
	LoanCheck() string
}

type Client struct {
	Name string
	Age  int
	Wage int
}

func LoanCheck(c Client) string {

	var IsAnyProduct bool = false
	k := ""

	for i := 0; i < 3; i++ {

		if c.Wage >= (Products[i]).Salary && c.Age >= (Products[i]).AgeMin && c.Age <= (Products[i].AgeMax) {
			IsAnyProduct = true
			//fmt.Println("  Вам доступен  ")
			//fmt.Println(Products[i].LoanName)
			//fmt.Println("На сумму до ", Products[i].Sum, "рублей")
			//fmt.Println("со ставкой ", Products[i].Rate, " % годовых")
			n1 := "Уважаемый " + c.Name + " Вам доступен " + Products[i].LoanName
			n1 = n1 + " На сумму до " + strconv.Itoa(Products[i].Sum) + " рублей " + " с ставкой " + strconv.Itoa(Products[i].Rate) + " процентов годовых "
			k = k + n1

		}
	}

	if IsAnyProduct != true {

		//fmt.Println("Уважаемый ", c.Name)
		//fmt.Println("К сожалению у нас сегодня нет для вас предложения по займу ")
		//return ("К сожалению у нас сегодня нет для вас предложения по займу ")
		k = "Уважаемый " + c.Name + " к сожалению у нас сегодня нет для вас предложения по займу "

	}
	return k //("К сожалению у нас сегодня нет для вас предложения по займу ")
}
