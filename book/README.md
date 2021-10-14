## Steps
1. export GO111MODULE=on
2. go get google.golang.org/protobuf/protoc-gen-go
3. export PATH=:$PATH:$(go env GOPATH)/bin"
4. protoc --help | less 
    To verify it works
5. protoc --go_out=. --go_opt=paths=source_relative book/book.proto
6. ...
7. write output to file and then:
8. hexdump -C book.protobuf
9. hexdump -C book.json