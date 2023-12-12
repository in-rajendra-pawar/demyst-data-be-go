
package helper

import (
	"net/http"
)

// ReturnJsonResponse function for returning movies data in JSON format
func ReturnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(resMessage)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	message := []byte(`{
		"success": false,
		"message": "Invalid HTTP method " ` + r.Method + `,
	   }`)
	ReturnJsonResponse(w, http.StatusMethodNotAllowed, message)
}