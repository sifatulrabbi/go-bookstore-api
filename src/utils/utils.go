package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}

}

func SendHTTPResp(wr http.ResponseWriter, status int, data interface{}) {
	var jsonData []byte
	if data != nil {
		jsonData, _ = json.Marshal(data)
	}
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(status)
	wr.Write(jsonData)
}
