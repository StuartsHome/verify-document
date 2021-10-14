.PHONY: protos

protos:
	# protoc -I rpc/ rpc/currency.proto --go-grpc_out=rpc/currency
	# protoc --go_out=rpc/currency,plugins=grpc:rpc/currency rpc/currency.proto
	# works protoc -I=rpc/ --go_out=rpc/currency rpc/currency.proto
	# works protoc -I=rpc/ --go_out=rpc/currency --plugin=grpc:rpc/currency rpc/currency.proto
	# protoc -I=rpc/ --go_out=plugins=grpc:rpc/currency rpc/currency.proto
	protoc -I=rpc/ --go-grpc_out=rpc/currency rpc/currency.proto