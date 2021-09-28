package responses

type ResponseEntity struct {
	// Error() string
	Reference string
	Message   string
	Status    int
	Success   bool
}

func (e *ResponseEntity) Error() string {
	return e.Message
}
