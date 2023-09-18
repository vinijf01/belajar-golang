package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"001","name":"Apple Macbook Pro","image_url":"http://exmaple.com/image.png"}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}

	json.Unmarshal(jsonByte, &result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "A001",
		"name":  "Apple Macbook Pro",
		"price": 20000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
