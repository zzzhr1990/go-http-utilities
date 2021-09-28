package responses

import "net/http"

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

func CreateNewBadRequestResponse(refrence string, message string) *ResponseEntity {
	return &ResponseEntity{
		Status:    http.StatusBadRequest,
		Message:   message,
		Reference: refrence,
	}
}

func CreateNewInternalErrorResponse(refrence string, message string) *ResponseEntity {
	return &ResponseEntity{
		Status:    http.StatusInternalServerError,
		Message:   message,
		Reference: refrence,
	}
}

func CreateNewForbiddenResponse(refrence string, message string) *ResponseEntity {
	return &ResponseEntity{
		Status:    http.StatusForbidden,
		Message:   message,
		Reference: refrence,
	}
}

func CreateNewUnauthorizedResponse(refrence string, message string) *ResponseEntity {
	return &ResponseEntity{
		Status:    http.StatusUnauthorized,
		Message:   message,
		Reference: refrence,
	}
}
