package structs

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
