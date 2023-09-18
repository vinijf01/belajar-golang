package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "001",
		Name:     "Apple Macbook Pro",
		ImageURL: "http://exmaple.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"001","name":"Apple Macbook Pro","image_URL":"http://exmaple.com/image.png"}`
	jsonByte := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonByte, product)
	fmt.Println(product)
}
