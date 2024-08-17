package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	//"server/structs"
)

var Products = []Loan{
	// предложение 1
	{
		LoanName: "займ краткосрочный",
		Sum:      20000,
		Rate:     15,
		AgeMin:   20,
		AgeMax:   70,
		Salary:   30000,
	},
	// предложение 2
	{
		LoanName: "займ среднесрочный",
		Sum:      100000,
		Rate:     14,
		AgeMin:   20,
		AgeMax:   60,
		Salary:   35000,
	},
	// предложение 3
	{
		LoanName: "займ долгосрочный",
		Sum:      1000000,
		Rate:     12,
		AgeMin:   20,
		AgeMax:   55,
		Salary:   50000,
	},
}

type Client struct {
	Name string
	Age  int
	Wage int
}

type Loan struct {
	LoanName string
	Sum      int
	Rate     int
	AgeMin   int
	AgeMax   int
	Salary   int
}

func LoanCheck(c Client) {

	var IsAnyProduct bool = false

	for i := 0; i < 3; i++ {

		if c.Wage >= (Products[i]).Salary && c.Age >= (Products[i]).AgeMin && c.Age <= (Products[i].AgeMax) {
			IsAnyProduct = true
			fmt.Println("Вам доступен ")
			fmt.Println(Products[i].LoanName)
			fmt.Println("На сумму до ", Products[i].Sum, "рублей")
			fmt.Println("со ставкой ", Products[i].Rate, "% годовых")
		}
	}
	if IsAnyProduct == false {

		fmt.Println("К сожалению у нас сегодня нет для вас предложения по займу ")
	}

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.
		Request) {
		w.Write([]byte("Hello, World!"))
	}).Methods(http.MethodGet)
	fmt.Println("Сервер запущен")

	http.ListenAndServe(":8080", r)

	//var dataClient = Client{}

	/*fmt.Println("введите ваше имя")
	fmt.Scanln(&dataClient.Name)
	fmt.Println("введите ваш возраст")
	fmt.Scanln(&dataClient.Age)
	fmt.Println("введите вашу з/п в руб.")
	fmt.Scanln(&dataClient.Wage)

	fmt.Println("Уважаемый ", dataClient.Name)
	LoanCheck(dataClient)
	*/
}
