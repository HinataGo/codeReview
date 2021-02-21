package reflect

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}
