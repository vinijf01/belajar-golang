package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

//type data golang -> json
func logJson(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {
	logJson("vini")
	logJson(1)
	logJson(true)
	logJson([]string{"Vini", "Jumatul", "Fitri"})
}
