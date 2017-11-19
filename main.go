package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	var data map[string]string
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data["origin"])
}
