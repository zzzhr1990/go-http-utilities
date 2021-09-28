package responses

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadBody(r *http.Request, in interface{}) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, in)
	if err != nil {
		return
	}
}

func RunAndResponse(w http.ResponseWriter, r *http.Request, fn func(w http.ResponseWriter, r *http.Request) (interface{}, error)) {

	response, err := fn(w, r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		echoError(w, r, err)
		return
	}
	if response == nil {
		response = map[string]interface{}{}
	}
	// check if
	//jsonconv.JsonSnakeCase
	b, _ := json.Marshal(response)
	w.Write(b)

}

// RunAndResponse Run
func echoError(w http.ResponseWriter, r *http.Request, err error) {
	// Error
	//err.Error()

	succ, ok := err.(*ResponseEntity)
	if ok {
		if succ.Status < 1 {
			succ.Status = http.StatusOK
		}
		if !succ.Success && len(succ.Message) < 1 {
			succ.Message = "Unknown internal error"
		}
		if !succ.Success && len(succ.Reference) < 1 {
			succ.Reference = "INTERNAL_ERROR"
		}
		b, _ := json.Marshal(succ)

		w.Write(b)
		return

	}

	status := http.StatusInternalServerError

	w.WriteHeader(status)
	mp := map[string]interface{}{
		"status":    status,
		"success":   false,
		"message":   err.Error(),
		"reference": "INTERNAL_ERROR",
	}
	b, _ := json.Marshal(mp)

	w.Write(b)

	// GRPC Error

}
