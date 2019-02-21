package app

import (
	"encoding/json"
	"net/http"
)

func ResponseWriter (flag bool, msg string, data interface {}, status int, w http.ResponseWriter){
	response := Response{
		Success: flag,
		Message: msg,
		Data: data,
	}

	payload, err := json.Marshal (response)

	if err !=  nil {
		http.Error (w, err.Error (), http.StatusInternalServerError)
		return
	}

	w.Header ().Set ("Content-Type", "application/json")
	w.Write (payload)

}
