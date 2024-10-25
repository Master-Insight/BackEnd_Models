package validatorpartner

// * New Create and return an instance of Validator with Data
func NewWithData(data interface{}) *Validator {
	val := New().Data(data)
	return val
}
