package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/stuartshome/verify-document/book"
	"google.golang.org/grpc"
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

	// decode the data from protobuf bytes
	data, err = ioutil.ReadFile("book.protobuf")
	if err != nil {
		log.Fatal(err)
	}

	readFromFile := book.Book{}
	if err := proto.Unmarshal(data, &readFromFile); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("book from protobuf file %+v\n", readFromFile)

	// 2.
	// connect to ther server
	conn, err := grpc.Dial(":9092")
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create service client
	client := book.NewBookStreamServiceClient(conn)
	in := &book.Request{Id: 1}
	stream, err := client.FetchResponse(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true // means stream if finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("response received: %s", resp.Result)
		}
	}()

	<-done // wait until all responses are received
	log.Printf("finished")
}
