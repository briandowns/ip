package ip

import (
	"encoding/json"
	"net/http"
)

// IP gets the calling IP address from httpbin.org
func IP() (string, error) {
	res, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	var data map[string]string
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return "", err
	}
	return data["origin"], nil
}
