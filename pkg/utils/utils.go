//this file will unmarshal the json

package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

//Client sends a request → r *http.Request reads it
//Server sends a response → w http.ResponseWriter writes it

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}