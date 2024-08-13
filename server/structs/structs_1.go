package structs

type Client struct {
	Name string
	Age  int
	Wage int
}

/*type Wage struct {
	WageMin int
	WageMax int
}*/

/*type Age struct {
	AgeMin int
	AgeMax int
	WageMin int
}*/

type Loan struct {
	LoanName string
	Sum      int
	Rate     int
	//LoanWageRange WageRange
	//LoanAgeRange  AgeRange
	AgeMin int
	AgeMax int
	Salary int
}
