package structs

var Products = []Loan{
	// предложение 1
	{
		LoanName: "займ краткосрочный",
		Sum:      20000,
		Rate:     15,
		//LoanWageRange WageRange
		//LoanAgeRange: AgeRange,
		AgeMin: 20,
		AgeMax: 70,
		Salary: 30000,
	},
	// предложение 2
	{
		LoanName: "займ среднесрочный",
		Sum:      100000,
		Rate:     14,
		//LoanWageRange WageRange
		//LoanAgeRange: AgeRange,
		AgeMin: 20,
		AgeMax: 60,
		Salary: 35000,
	},
	// предложение 3
	{
		LoanName: "займ долгосрочный",
		Sum:      1000000,
		Rate:     12,
		//LoanWageRange WageRange
		//LoanAgeRange: AgeRange,
		AgeMin: 20,
		AgeMax: 55,
		Salary: 50000,
	},
}
