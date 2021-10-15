package book

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	// bk "github.com/stuartshome/verify-document/book"
	"google.golang.org/grpc"
)

type server struct{}

func (s server) FetchResponse(in *Request, srv BookStreamService_FetchResponseServer) error {
	log.Printf("fetch response for id: %d", in.Id)

	// use wait group to allow process to be concurrent
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(count int64) {
			defer wg.Done()

			time.Sleep(time.Duration(count) * time.Second)
			resp := Response{Result: fmt.Sprintf("Request #%d For Id:%d", count, in.Id)}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			log.Printf("finishing request number: %d", count)
		}(int64(i))
	}

	wg.Wait()
	return nil
}

func main() {
	// create listener
	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	// register server
	var server BookStreamServiceServer
	RegisterBookStreamServiceServer(s, server)

	log.Println("start server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve request: %v", err)
	}
}
