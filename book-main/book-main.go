package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/stuartshome/verify-document/book"
)

func main() {

	bookPlato := &book.Book{
		Id:    1,
		Title: "the road to the pier",
		Authors: []*book.Author{
			{Id: 1, Name: "Plato"},
		},
		Category: book.Category_History,
	}

	data, err := proto.Marshal(bookPlato)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	ioutil.WriteFile("book.protobuf", data, 0600)
	data, err = json.Marshal(bookPlato)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	ioutil.WriteFile("book.json", data, 0600)

	// decod the data from protobuf bytes
	data, err = ioutil.ReadFile("book.protobuf")
	if err != nil {
		log.Fatal(err)
	}

	readFromFile := book.Book{}
	if err := proto.Unmarshal(data, &readFromFile); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("book from protobuf file %+v\n", readFromFile)
}
