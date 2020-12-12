package invoice

type Invoice struct {
	country string
	city    string
	total   float64
	cliente customer.Customer
}
