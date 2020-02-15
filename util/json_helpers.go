package util

import "encoding/json"

func ToJson(obj interface{}) []byte {
	json, err := json.Marshal(obj)
	handleErrors(err)

	return json
}

func handleErrors(err error) {
	if err != nil {
		panic(err)
	}
}
