package structs

import "fmt"

func LoanCheck(c Client) {
	//fmt.Println(c.Wage, (Products[0]).Salary)
	var IsAnyProduct bool = false

	for i := 0; i < 3; i++ {

		//b := (Products[i]).AgeMin
		if c.Wage >= (Products[i]).Salary && c.Age >= (Products[i]).AgeMin && c.Age <= (Products[i].AgeMax) {
			IsAnyProduct = true
			fmt.Println("Вам доступен ")
			fmt.Println(Products[i].LoanName)
			fmt.Println("На сумму до ", Products[i].Sum, "рублей")
			fmt.Println("со ставкой ", Products[i].Rate, "% годовых")
			//fmt.Println(Products[i].LoanName)
			//fmt.Println(Products[i].LoanName)
		} else if IsAnyProduct == false {

			fmt.Println("Уважаемый ", c.Name)
			fmt.Println("К сожалению у нас сегодня нет для вас предложения по займу ")
		}
	}
}

/*if age >= l.LoanAgeRange.MiniAge && age <= l.LoanAgeRange.MaxAge {
		if wage >= l.LoanWageRange.MiniWage && wage <= l.LoanWageRange.MaxWage {
			fmt.Println("вам доступен займ", l.LoanName)
			IsAnyProduct = true
		}
	}
}

func LoanAgregator(c Client) {
	for _, v := range Products {
		LoanCheck(c.Age, c.Salary, v)
	}
	if IsAnyProduct == false {
		fmt.Println("к сожалению для вас нет предложений")
	}*/
