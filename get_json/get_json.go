package get_json

import (
	"net/http"
	"encoding/json"
)
//375794

// Create http client
var Client *http.Client

func GetJson(url string, target interface{}) error {
	resp, err := Client.Get(url)
	if err != nil {
		return err
	}
	// defer closing response till we are done getting it
	defer resp.Body.Close()

	// Stream the response. This resturns an error
	return json.NewDecoder(resp.Body).Decode(target)
}