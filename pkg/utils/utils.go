package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			log.Fatal("Unable to parse body\n", err)
		}
	} else {
		log.Fatal(err)
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
