package utils

//this utils file is serving the marshalling/unmarshalling process of json objects
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//implementing a function to handle unmarshalling and reporting any errors using json.Unmarshall function
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
