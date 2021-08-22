package main

type (
	Customer struct {
		FirstName string
		LastName  string
		FullName  string
	}
)

func GetCustomers() (customers []Customer) {
	customers = append(customers,
		Customer{FirstName: "Marc", LastName: "Chagal"},
		Customer{FirstName: "Dan", LastName: "Henderson"},
	)
	return customers
}
