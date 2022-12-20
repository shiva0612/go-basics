package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	marshalling()
	encoding_decoding()
}

func encoding_decoding() {
	f, _ := os.Create("encode.json")
	a := map[int]int{1: 1, 2: 2}
	err := json.NewEncoder(f).Encode(a)
	if err != nil {
		log.Println(err.Error())
	}

	ff, _ := os.Open("encode.json")
	aa := map[string]int{}
	json.NewDecoder(ff).Decode(&aa)
	fmt.Println(aa)
}

/*
type myresponse struct{

	extresponse json.RawMessage
	requestedby string
	timetaken time.Time
	cost int

}
extapi <- json response
directly put it in myreponse
when we add some stats to it and marshall it will be just fine

similar to marshall u can also unmarshall to rawjson or map[string]interface{} if u need to query

	{
		api:{weird big json}
	}
*/

func marshalling() {
	simple_bytes_gone_wrong()
	map_string_interface()
	json_raw_msg()
}

func simple_bytes_gone_wrong() {
	h := []byte(`{"precomputed": true}`)

	c := struct {
		Header *[]byte `json:"header"`
		Body   string  `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}

func map_string_interface() {
	hb := []byte(`{"precomputed": true}`)

	mpi := map[string]interface{}{}
	json.Unmarshal(hb, &mpi)

	c := struct {
		Header *map[string]interface{} `json:"header"`
		Body   string                  `json:"body"`
	}{Header: &mpi, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}

func json_raw_msg() {
	h := json.RawMessage(`{"precomputed": true}`)

	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}
