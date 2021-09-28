package responses

import "net/http"

type ResponseEntity struct {
	// Error() string
	Reference string `json:"reference"`
	Message   string `json:"message"`
	Status    int    `json:"status"`
	Success   bool   `json:"success"`
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

func CreateNewPureInternalErrorResponse() *ResponseEntity {
	return &ResponseEntity{
		Status:    http.StatusInternalServerError,
		Message:   "The server encountered an unexpected condition that prevented it from fulfilling the request",
		Reference: "INTERNAL_ERROR",
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
