package rest

import (
	"encoding/json"
	"net/http"
)

type Url struct{
	Url string `json:"url"`
	Method string `json:"method"`
}

func Verify(w http.ResponseWriter, r *http.Request){
	var reqUrl Url
	err := json.NewDecoder(r.Body).Decode(&reqUrl)
	if err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

}