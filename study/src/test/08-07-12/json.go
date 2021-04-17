package main

import (
	"bytes"
	"encoding/json"
	"log"
)

type X struct {
	A string `json:"a,omitempty"`
}

func main() {
	data := `{
"b":"bb"
}`

	x := &X{}

	decoder := json.NewDecoder(bytes.NewBufferString(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(x)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("x: %v\n", x)
}
