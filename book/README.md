## Steps
1. export GO111MODULE=on
2. go get google.golang.org/protobuf/protoc-gen-go
3. export PATH=:$PATH:$(go env GOPATH)/bin"
4. protoc --help | less 
    To verify it works
5. Generate client
    protoc --go_out=. --go_opt=paths=source_relative book/book.proto
6. Generate server
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative book/book.proto
7. ...
8. write output to file and then:
9. hexdump -C book.protobuf
10. hexdump -C book.json



# Notes
The protoc-gen-go and protoc-gen-go-grpc are two different protoc plugins. While the former generates Go code for protobuf message definitions, the latter generates Go code for service definitions.