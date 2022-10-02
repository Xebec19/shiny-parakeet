package util

import (
	"encoding/json"
	"log"
)

func Stringify(payload interface{}) string {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("--error while marshalling json %s", err)
	}
	return string(data)
}
